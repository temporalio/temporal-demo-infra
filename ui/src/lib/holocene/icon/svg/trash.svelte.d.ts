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
export declare type TrashProps = typeof __propDef.props;
export declare type TrashEvents = typeof __propDef.events;
export declare type TrashSlots = typeof __propDef.slots;
export default class Trash extends SvelteComponentTyped<TrashProps, TrashEvents, TrashSlots> {
}
export {};
