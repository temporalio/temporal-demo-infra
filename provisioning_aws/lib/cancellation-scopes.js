"use strict";
// TODO: put this into the activities-cancellation-heartbeating repo
Object.defineProperty(exports, "__esModule", { value: true });
exports.shieldAwaitedInRootScope = exports.sharedScopes = exports.cancellationScopesWithCallbacks = exports.resumeAfterCancellation = exports.multipleActivitiesSingleTimeout = exports.nonCancellable = exports.handleExternalWorkflowCancellationWhileActivityRunning = exports.cancelTimerAltImpl = exports.cancelTimer = void 0;
/**
 * Demonstrates the basics of cancellation scopes.
 */
// @@@SNIPSTART typescript-cancel-a-timer-from-workflow
const workflow_1 = require("@temporalio/workflow");
async function cancelTimer() {
    // Timers and Activities are automatically cancelled when their containing scope is cancelled.
    try {
        await workflow_1.CancellationScope.cancellable(async () => {
            const promise = (0, workflow_1.sleep)(1); // <-- Will be cancelled because it is attached to this closure's scope
            workflow_1.CancellationScope.current().cancel();
            await promise; // <-- Promise must be awaited in order for `cancellable` to throw
        });
    }
    catch (e) {
        if (e instanceof workflow_1.CancelledFailure) {
            console.log('Timer cancelled 👍');
        }
        else {
            throw e; // <-- Fail the workflow
        }
    }
}
exports.cancelTimer = cancelTimer;
// @@@SNIPEND
/**
 * Alternative implementation with cancellation from an outer scope.
 */
// @@@SNIPSTART typescript-cancel-a-timer-from-workflow-alternative-impl
async function cancelTimerAltImpl() {
    try {
        const scope = new workflow_1.CancellationScope();
        const promise = scope.run(() => (0, workflow_1.sleep)(1));
        scope.cancel(); // <-- Cancel the timer created in scope
        await promise; // <-- Throws CancelledFailure
    }
    catch (e) {
        if (e instanceof workflow_1.CancelledFailure) {
            console.log('Timer cancelled 👍');
        }
        else {
            throw e; // <-- Fail the workflow
        }
    }
}
exports.cancelTimerAltImpl = cancelTimerAltImpl;
// @@@SNIPEND
/**
 * Demonstrates how to clean up after cancellation.
 */
// @@@SNIPSTART typescript-handle-external-workflow-cancellation-while-activity-running
const workflow_2 = require("@temporalio/workflow");
const { httpPostJSON, httpGetJSON, cleanup } = (0, workflow_2.proxyActivities)({
    startToCloseTimeout: '10m',
});
async function handleExternalWorkflowCancellationWhileActivityRunning(url, data) {
    try {
        await httpPostJSON(url, data);
    }
    catch (err) {
        if ((0, workflow_2.isCancellation)(err)) {
            console.log('Workflow cancelled');
            // Cleanup logic must be in a nonCancellable scope
            // If we'd run cleanup outside of a nonCancellable scope it would've been cancelled
            // before being started because the Workflow's root scope is cancelled.
            await workflow_1.CancellationScope.nonCancellable(() => cleanup(url));
        }
        throw err; // <-- Fail the Workflow
    }
}
exports.handleExternalWorkflowCancellationWhileActivityRunning = handleExternalWorkflowCancellationWhileActivityRunning;
// @@@SNIPEND
// @@@SNIPSTART typescript-non-cancellable-shields-children
async function nonCancellable(url) {
    // Prevent Activity from being cancelled and await completion.
    // Note that the Workflow is completely oblivious and impervious to cancellation in this example.
    return workflow_1.CancellationScope.nonCancellable(() => httpGetJSON(url));
}
exports.nonCancellable = nonCancellable;
// @@@SNIPEND
// @@@SNIPSTART typescript-multiple-activities-single-timeout-workflow
function multipleActivitiesSingleTimeout(urls, timeoutMs) {
    // If timeout triggers before all activities complete
    // the Workflow will fail with a CancelledFailure.
    return workflow_1.CancellationScope.withTimeout(timeoutMs, () => Promise.all(urls.map((url) => httpGetJSON(url))));
}
exports.multipleActivitiesSingleTimeout = multipleActivitiesSingleTimeout;
// @@@SNIPEND
/**
 * Demonstrates how to make Workflow aware of cancellation while waiting on nonCancellable scope.
 */
// @@@SNIPSTART typescript-cancel-requested-with-non-cancellable
async function resumeAfterCancellation(url) {
    let result = undefined;
    const scope = new workflow_1.CancellationScope({ cancellable: false });
    const promise = scope.run(() => httpGetJSON(url));
    try {
        result = await Promise.race([scope.cancelRequested, promise]);
    }
    catch (err) {
        if (!(err instanceof workflow_1.CancelledFailure)) {
            throw err;
        }
        // Prevent Workflow from completing so Activity can complete
        result = await promise;
    }
    return result;
}
exports.resumeAfterCancellation = resumeAfterCancellation;
// @@@SNIPEND
/**
 * Demonstrates how to use cancellation scopes with callbacks.
 */
// @@@SNIPSTART typescript-cancellation-scopes-with-callbacks
function doSomething(callback) {
    setTimeout(callback, 10);
}
async function cancellationScopesWithCallbacks() {
    await new Promise((resolve, reject) => {
        doSomething(resolve);
        workflow_1.CancellationScope.current().cancelRequested.catch(reject);
    });
}
exports.cancellationScopesWithCallbacks = cancellationScopesWithCallbacks;
// @@@SNIPEND
// @@@SNIPSTART typescript-shared-promise-scopes
async function sharedScopes() {
    // Start activities in the root scope
    const p1 = httpGetJSON('http://url1.ninja');
    const p2 = httpGetJSON('http://url2.ninja');
    const scopePromise = workflow_1.CancellationScope.cancellable(async () => {
        const first = await Promise.race([p1, p2]);
        // Does not cancel activity1 or activity2 as they're linked to the root scope
        workflow_1.CancellationScope.current().cancel();
        return first;
    });
    return await scopePromise;
    // The Activity that did not complete will effectively be cancelled when
    // Workflow completes unless the Activity is awaited:
    // await Promise.all([p1, p2]);
}
exports.sharedScopes = sharedScopes;
// @@@SNIPEND
// @@@SNIPSTART typescript-shield-awaited-in-root-scope
async function shieldAwaitedInRootScope() {
    let p = undefined;
    await workflow_1.CancellationScope.nonCancellable(async () => {
        p = httpGetJSON('http://example.com'); // <-- Start activity in nonCancellable scope without awaiting completion
    });
    // Activity is shielded from cancellation even though it is awaited in the cancellable root scope
    return p;
}
exports.shieldAwaitedInRootScope = shieldAwaitedInRootScope;
// @@@SNIPEND
//# sourceMappingURL=cancellation-scopes.js.map