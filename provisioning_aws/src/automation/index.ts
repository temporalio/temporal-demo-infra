import { LocalProgramArgs, LocalWorkspace } from "@pulumi/pulumi/automation";
import * as upath from "upath";
import {Context} from "@temporalio/activity";
import {provider} from "@pulumi/pulumi";
// import { setInterval } from 'timers/promises';

const args = process.argv.slice(2);
let destroy = false;
if (args.length > 0 && args[0]) {
    destroy = args[0] === "destroy";
}

interface ProvisionFoundationResourcesRequest {
    applicationId: string
    region: string
    profile: string
    bucketName: string
    teamId: string
    applicationName: string
}
export interface ProvisionFoundationResourcesResponse {
    applicationId: string
    summary: string[]
    completionDateTime: string
    region: string
}

export const provisionFoundationResources = async (cmd: ProvisionFoundationResourcesRequest) : Promise<ProvisionFoundationResourcesResponse> => {
    const { log, info, sleep } = Context.current();
    log.info('provisioning resources', { 'cmd': cmd})
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
    if(!cmd.region || !cmd.profile) {
        throw new Error('region and profile is required')
    }

    // Create our stack using a local program
    // in the ../foundation directory
    const args: LocalProgramArgs = {
        stackName: "dev",
        workDir: upath.joinSafe(__dirname, "..", "foundation"),
    };

    // create (or select if one already exists) a stack that uses our local program
    const stack = await LocalWorkspace.createOrSelectStack(args);

    log.info("successfully initialized stack");
    log.info("setting up config");
    await stack.setConfig("aws:region", { value: cmd.region });
    await stack.setConfig("aws:profile", { value: cmd.profile });
    await stack.setConfig('foundation:applicationId', { value: cmd.applicationId})
    await stack.setConfig('foundation:applicationName', { value: cmd.applicationName})
    await stack.setConfig('foundation:bucketName', { value: cmd.bucketName})
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
    const outs:string[] = []
    const upRes = await stack.up({ onOutput: function(str) {
            log.info(str)
            outs.push(str)
        }});
    summary = JSON.stringify(upRes.summary.resourceChanges, null, 4)
    log.info(`update summary: \n${summary}`);

    return {
        summary: outs,
        applicationId: cmd.applicationId,
        completionDateTime: new Date().toUTCString(),
        region: cmd.region,
    }
}
