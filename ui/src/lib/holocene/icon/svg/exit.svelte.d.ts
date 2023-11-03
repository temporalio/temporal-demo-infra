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
export declare type ExitProps = typeof __propDef.props;
export declare type ExitEvents = typeof __propDef.events;
export declare type ExitSlots = typeof __propDef.slots;
export default class Exit extends SvelteComponentTyped<ExitProps, ExitEvents, ExitSlots> {
}
export {};
