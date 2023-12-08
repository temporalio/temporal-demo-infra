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
export declare type EyeHideProps = typeof __propDef.props;
export declare type EyeHideEvents = typeof __propDef.events;
export declare type EyeHideSlots = typeof __propDef.slots;
export default class EyeHide extends SvelteComponentTyped<EyeHideProps, EyeHideEvents, EyeHideSlots> {
}
export {};
