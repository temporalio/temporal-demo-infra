"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const client_1 = require("@temporalio/client");
const workflows_1 = require("./workflows");
async function run() {
    const connection = new client_1.Connection();
    const client = new client_1.WorkflowClient(connection.service);
    const handle = await client.start(workflows_1.runCancellableActivity, {
        taskQueue: 'cancellation-heartbeating',
        workflowId: 'cancellation-heartbeating-0',
    });
    // Simulate waiting for some time
    // Cancel may be immediately called, waiting is not needed
    await new Promise((resolve) => setTimeout(resolve, 40 * 1000));
    await handle.cancel();
    console.log('Cancelled workflow successfully');
    try {
        await handle.result();
    }
    catch (err) {
        if (err instanceof client_1.WorkflowFailedError && err.cause instanceof client_1.CancelledFailure) {
            console.log('handle.result() threw because Workflow was cancelled');
        }
        else {
            throw err;
        }
    }
}
run().catch((err) => {
    console.error(err);
    process.exit(1);
});
//# sourceMappingURL=client.js.map