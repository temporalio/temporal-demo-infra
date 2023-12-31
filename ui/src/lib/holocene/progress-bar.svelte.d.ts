import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        width?: string;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type ProgressBarProps = typeof __propDef.props;
export declare type ProgressBarEvents = typeof __propDef.events;
export declare type ProgressBarSlots = typeof __propDef.slots;
export default class ProgressBar extends SvelteComponentTyped<ProgressBarProps, ProgressBarEvents, ProgressBarSlots> {
}
export {};
