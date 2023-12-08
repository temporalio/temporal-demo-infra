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
export declare type BookmarkProps = typeof __propDef.props;
export declare type BookmarkEvents = typeof __propDef.events;
export declare type BookmarkSlots = typeof __propDef.slots;
export default class Bookmark extends SvelteComponentTyped<BookmarkProps, BookmarkEvents, BookmarkSlots> {
}
export {};
