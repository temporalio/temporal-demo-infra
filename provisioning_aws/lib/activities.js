"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.fakeProgress = void 0;
// @@@SNIPSTART typescript-activity-fake-progress
const activity_1 = require("@temporalio/activity");
const common_1 = require("@temporalio/common");
async function fakeProgress(sleepIntervalMs = 1000) {
    try {
        // allow for resuming from heartbeat
        const startingPoint = activity_1.Context.current().info.heartbeatDetails || 1;
        console.log('Starting activity at progress:', startingPoint);
        for (let progress = startingPoint; progress <= 100; ++progress) {
            // simple utility to sleep in activity for given interval or throw if Activity is cancelled
            // don't confuse with Workflow.sleep which is only used in Workflow functions!
            console.log('Progress:', progress);
            await activity_1.Context.current().sleep(sleepIntervalMs);
            activity_1.Context.current().heartbeat(progress);
        }
    }
    catch (err) {
        if (err instanceof common_1.CancelledFailure) {
            console.log('Fake progress activity cancelled');
            // Cleanup
        }
        throw err;
    }
}
exports.fakeProgress = fakeProgress;
// @@@SNIPEND
//# sourceMappingURL=activities.js.map