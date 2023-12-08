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
export declare type CartProps = typeof __propDef.props;
export declare type CartEvents = typeof __propDef.events;
export declare type CartSlots = typeof __propDef.slots;
export default class Cart extends SvelteComponentTyped<CartProps, CartEvents, CartSlots> {
}
export {};
