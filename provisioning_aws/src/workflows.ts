import {proxyActivities, isCancellation, ActivityCancellationType, log, setHandler} from '@temporalio/workflow';
import type * as activities from './activities';
import {ProvisionFoundationResourcesResponse} from "./activities";

const { fakeProgress, provisionFoundationResources } = proxyActivities<typeof activities>({
  startToCloseTimeout: '120s',
  // heartbeatTimeout: '3s',
  // Don't send rejection to our Workflow until the Activity has confirmed cancellation
  cancellationType: ActivityCancellationType.WAIT_CANCELLATION_COMPLETED,
});

export async function runCancellableActivity(): Promise<ProvisionFoundationResourcesResponse> {
  try {
    let resp = await provisionFoundationResources({
      applicationId: 'foo',
      region: 'us-east-1',
      profile: 'iac',
      bucketName: `foo-${new Date().getUTCDate()}`,
      teamId: '123',
      applicationName: 'youknow',
    });
    // fakeProgress(100)
    // log.info('completed', {'summary': resp.summary})
    return resp
  } catch (err) {
    if (isCancellation(err)) {
      log.info('Workflow cancelled along with its activity');
      // To clean up use CancellationScope.nonCancellable
    }
    throw err;
  }
}

async function charge(obj:any, amount: any) { return undefined}
async function sleep(obj: any) { return undefined}
async function sendEmail(obj: any, val: string) { return undefined}

export async function subscription(user, amount) {
  let canceled = false
  // @ts-ignore
  setHandler('cancel', async () => {
    canceled = true
    await sendEmail(user, 'your subscription has been canceled')
  })
  while (!canceled) {
    await charge(user, amount)
    await sleep('30 days')
  }
}
