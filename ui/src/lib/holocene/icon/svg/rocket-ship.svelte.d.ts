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
export declare type RocketShipProps = typeof __propDef.props;
export declare type RocketShipEvents = typeof __propDef.events;
export declare type RocketShipSlots = typeof __propDef.slots;
export default class RocketShip extends SvelteComponentTyped<RocketShipProps, RocketShipEvents, RocketShipSlots> {
}
export {};
