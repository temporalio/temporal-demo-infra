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
export declare type DescendingProps = typeof __propDef.props;
export declare type DescendingEvents = typeof __propDef.events;
export declare type DescendingSlots = typeof __propDef.slots;
export default class Descending extends SvelteComponentTyped<DescendingProps, DescendingEvents, DescendingSlots> {
}
export {};
