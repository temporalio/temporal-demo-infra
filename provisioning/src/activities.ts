// @@@SNIPSTART typescript-activity-fake-progress
import { CancelledFailure, Context } from '@temporalio/activity';
import {LocalProgramArgs, LocalWorkspace} from "@pulumi/pulumi/automation";
import upath from "upath";
import {ProvisionResourcesResponse} from "./automation";

export * from  './automation'
export async function fakeProgress(sleepIntervalMs = 1000): Promise<void> {
  const { log, info, sleep, heartbeat } = Context.current();
  try {
    // allow for resuming from heartbeat
    const startingPoint = info.heartbeatDetails || 1;
    log.info('Starting activity at progress', { startingPoint });
    for (let progress = startingPoint; progress <= 100; ++progress) {
      // simple utility to sleep in activity for given interval or throw if Activity is cancelled
      // don't confuse with Workflow.sleep which is only used in Workflow functions!
      log.info('Progress', { progress });
      await sleep(sleepIntervalMs);
      heartbeat(progress);
    }
  } catch (err) {
    if (err instanceof CancelledFailure) {
      log.warn('activity cancelled', { message: err.message });
      // Cleanup
    }
    throw err;
  }
}
// interface ProvisionResourceRequest {
//   applicationId: string
//   region: string
//   profile: string
//   bucketName: string
// }
// export interface  ProvisionResourcesResponse {
//   applicationId: string
//   summary: string
// }
// export const provisionResources = async (cmd: ProvisionResourceRequest) : Promise<ProvisionResourcesResponse> => {
//   const { log, info, sleep, heartbeat } = Context.current();
//   let summary = ''
//   // heartbeat('foo')
//   // const timeout = setInterval(heartbeat.bind(this),1000)
//
//
//   // Create our stack using a local program
//   // in the ../website directory
//   const args: LocalProgramArgs = {
//     stackName: "dev",
//     workDir: upath.joinSafe(__dirname, "website"),
//   };
//
//   // create (or select if one already exists) a stack that uses our local program
//   const stack = await LocalWorkspace.createOrSelectStack(args);
//
//   log.info("successfully initialized stack");
//   log.info("setting up config");
//   await stack.setConfig("aws:region", { value: cmd.region });
//   await stack.setConfig("aws:profile", { value: cmd.profile });
//   await stack.setConfig('website:applicationId', { value: cmd.applicationId})
//   await stack.setConfig('website:bucketName', { value: cmd.bucketName})
//   log.info("config set");
//   log.info("refreshing stack...");
//   await stack.refresh({ onOutput: log.info });
//   log.info("refresh complete");
//
//   // if (destroy) {
//   //   console.info("destroying stack...");
//   //   await stack.destroy({onOutput: log.info});
//   //   console.info("stack destroy complete");
//   //   process.exit(0);
//   // }
//
//   log.info("updating stack...");
//   const upRes = await stack.up({ onOutput: log.info });
//   summary = JSON.stringify(upRes.summary.resourceChanges, null, 4)
//   log.info(`update summary: \n${summary}`);
//   // log.info(`website url: ${upRes.outputs.websiteUrl.value}`);
//   // clearTimeout(timeout)
//   return {
//     summary,
//     applicationId: cmd.applicationId,
//   }
// }
// @@@SNIPEND
