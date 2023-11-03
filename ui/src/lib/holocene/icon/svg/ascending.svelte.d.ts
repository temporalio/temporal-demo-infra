import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type AscendingProps = typeof __propDef.props;
export declare type AscendingEvents = typeof __propDef.events;
export declare type AscendingSlots = typeof __propDef.slots;
export default class Ascending extends SvelteComponentTyped<AscendingProps, AscendingEvents, AscendingSlots> {
}
export {};
