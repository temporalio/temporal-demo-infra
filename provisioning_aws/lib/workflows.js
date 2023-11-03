"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.runCancellableActivity = void 0;
const workflow_1 = require("@temporalio/workflow");
const { fakeProgress } = (0, workflow_1.proxyActivities)({
    startToCloseTimeout: '60s',
    heartbeatTimeout: '3s',
    // Don't send rejection to our Workflow until the Activity has confirmed cancellation
    cancellationType: workflow_1.ActivityCancellationType.WAIT_CANCELLATION_COMPLETED,
});
async function runCancellableActivity() {
    try {
        await fakeProgress();
    }
    catch (err) {
        if ((0, workflow_1.isCancellation)(err)) {
            console.log('Workflow cancelled along with its activity');
            // To clean up use CancellationScope.nonCancellable
        }
        throw err;
    }
}
exports.runCancellableActivity = runCancellableActivity;
//# sourceMappingURL=workflows.js.map