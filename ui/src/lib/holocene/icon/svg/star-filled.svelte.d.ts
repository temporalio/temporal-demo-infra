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
export declare type StarFilledProps = typeof __propDef.props;
export declare type StarFilledEvents = typeof __propDef.events;
export declare type StarFilledSlots = typeof __propDef.slots;
export default class StarFilled extends SvelteComponentTyped<StarFilledProps, StarFilledEvents, StarFilledSlots> {
}
export {};
