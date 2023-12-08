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
export declare type EyeShowProps = typeof __propDef.props;
export declare type EyeShowEvents = typeof __propDef.events;
export declare type EyeShowSlots = typeof __propDef.slots;
export default class EyeShow extends SvelteComponentTyped<EyeShowProps, EyeShowEvents, EyeShowSlots> {
}
export {};
