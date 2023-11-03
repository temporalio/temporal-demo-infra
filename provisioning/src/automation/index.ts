import { LocalProgramArgs, LocalWorkspace } from "@pulumi/pulumi/automation";
import * as upath from "upath";
import {Context} from "@temporalio/activity";
// import { setInterval } from 'timers/promises';

const args = process.argv.slice(2);
let destroy = false;
if (args.length > 0 && args[0]) {
    destroy = args[0] === "destroy";
}

interface ProvisionResourceRequest {
    applicationId: string
    region: string
    profile: string
    bucketName: string
}
export interface  ProvisionResourcesResponse {
    applicationId: string
    summary: string
}
async function autoheartbeat(fn:any) {
    async function heartbeat() {
        const cx =  Context.current();
        for (;;) {
            await cx.sleep((cx.info.heartbeatTimeoutMs || 100) / 2);
            cx.heartbeat();
        }
    }
    return (...args:any[]) => Promise.race([heartbeat(), fn(...args)]);
}
export const provisionResources = async (cmd: ProvisionResourceRequest) : Promise<ProvisionResourcesResponse> => {
    const { log, info, sleep } = Context.current();
    let summary = ''
    async function heartbeat() {
        const cx =  Context.current();
        for (;;) {
            await cx.sleep((cx.info.heartbeatTimeoutMs || 100) / 2);
            cx.heartbeat();
            if(summary.length > 0) {
                break
            }
        }
    }
    heartbeat()

    // heartbeat('foo')
    // const timeout = setInterval(heartbeat.bind(this),1000)


    // Create our stack using a local program
    // in the ../website directory
    const args: LocalProgramArgs = {
        stackName: "dev",
        workDir: upath.joinSafe(__dirname, "..", "website"),
    };

    // create (or select if one already exists) a stack that uses our local program
    const stack = await LocalWorkspace.createOrSelectStack(args);

    log.info("successfully initialized stack");
    log.info("setting up config");
    await stack.setConfig("aws:region", { value: cmd.region });
    await stack.setConfig("aws:profile", { value: cmd.profile });
    await stack.setConfig('website:applicationId', { value: cmd.applicationId})
    await stack.setConfig('website:bucketName', { value: cmd.bucketName})
    log.info("config set");
    log.info("refreshing stack...");
    // idempotency!
    await stack.refresh({ onOutput: log.info });
    log.info("refresh complete");

    if (destroy) {
        console.info("destroying stack...");
        await stack.destroy({onOutput: log.info});
        console.info("stack destroy complete");
        process.exit(0);
    }

    log.info("updating stack...");
    const upRes = await stack.up({ onOutput: log.info });
    summary = JSON.stringify(upRes.summary.resourceChanges, null, 4)
    log.info(`update summary: \n${summary}`);
    // log.info(`website url: ${upRes.outputs.websiteUrl.value}`);
    // clearTimeout(timeout)
    return {
        summary,
        applicationId: cmd.applicationId,
    }
}
