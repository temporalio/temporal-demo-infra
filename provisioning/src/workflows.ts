import { proxyActivities, isCancellation, ActivityCancellationType, log } from '@temporalio/workflow';
import type * as activities from './activities';
import {ProvisionResourcesResponse} from "./activities";

const { fakeProgress, provisionResources } = proxyActivities<typeof activities>({
  startToCloseTimeout: '120s',
  // heartbeatTimeout: '3s',
  // Don't send rejection to our Workflow until the Activity has confirmed cancellation
  cancellationType: ActivityCancellationType.WAIT_CANCELLATION_COMPLETED,
});

export async function runCancellableActivity(): Promise<ProvisionResourcesResponse> {
  try {
    let resp = await provisionResources({
      applicationId: 'foo',
      region: 'us-east-1',
      profile: 'iac',
      bucketName: `foo-${new Date().getUTCDate()}`,
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
