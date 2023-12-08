import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        href?: string;
        label: string;
        amount?: number;
        active?: boolean;
        disabled?: boolean;
    };
    events: {
        click: MouseEvent;
        keypress: KeyboardEvent;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type TabProps = typeof __propDef.props;
export declare type TabEvents = typeof __propDef.events;
export declare type TabSlots = typeof __propDef.slots;
export default class Tab extends SvelteComponentTyped<TabProps, TabEvents, TabSlots> {
}
export {};
