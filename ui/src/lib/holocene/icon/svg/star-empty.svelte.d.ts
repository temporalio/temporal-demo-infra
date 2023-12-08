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
export declare type StarEmptyProps = typeof __propDef.props;
export declare type StarEmptyEvents = typeof __propDef.events;
export declare type StarEmptySlots = typeof __propDef.slots;
export default class StarEmpty extends SvelteComponentTyped<StarEmptyProps, StarEmptyEvents, StarEmptySlots> {
}
export {};
