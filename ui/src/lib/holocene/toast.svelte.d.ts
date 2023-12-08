import { SvelteComponentTyped } from "svelte";
import type { ToastVariant } from './toaster.svelte';
declare const __propDef: {
    props: {
        id: string;
        variant: ToastVariant;
    };
    events: {
        dismiss: CustomEvent<{
            id: string;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        default: {};
    };
};
export declare type ToastProps = typeof __propDef.props;
export declare type ToastEvents = typeof __propDef.events;
export declare type ToastSlots = typeof __propDef.slots;
export default class Toast extends SvelteComponentTyped<ToastProps, ToastEvents, ToastSlots> {
}
export {};
