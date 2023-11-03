export declare function cancelTimer(): Promise<void>;
/**
 * Alternative implementation with cancellation from an outer scope.
 */
export declare function cancelTimerAltImpl(): Promise<void>;
export declare function handleExternalWorkflowCancellationWhileActivityRunning(url: string, data: unknown): Promise<void>;
export declare function nonCancellable(url: string): Promise<any>;
export declare function multipleActivitiesSingleTimeout(urls: string[], timeoutMs: number): Promise<any>;
/**
 * Demonstrates how to make Workflow aware of cancellation while waiting on nonCancellable scope.
 */
export declare function resumeAfterCancellation(url: string): Promise<any>;
export declare function cancellationScopesWithCallbacks(): Promise<void>;
export declare function sharedScopes(): Promise<any>;
export declare function shieldAwaitedInRootScope(): Promise<any>;
//# sourceMappingURL=cancellation-scopes.d.ts.map