import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        showOn?: number;
        scrollToContainer?: boolean;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        default: {};
    };
};
export declare type ScrollToTopProps = typeof __propDef.props;
export declare type ScrollToTopEvents = typeof __propDef.events;
export declare type ScrollToTopSlots = typeof __propDef.slots;
export default class ScrollToTop extends SvelteComponentTyped<ScrollToTopProps, ScrollToTopEvents, ScrollToTopSlots> {
}
export {};
