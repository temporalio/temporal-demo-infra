import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        href?: string;
        selectable?: boolean;
        selected?: boolean;
    };
    events: {
        click: MouseEvent;
        change: CustomEvent<{
            checked: boolean;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        default: {};
    };
};
export declare type TableRowProps = typeof __propDef.props;
export declare type TableRowEvents = typeof __propDef.events;
export declare type TableRowSlots = typeof __propDef.slots;
export default class TableRow extends SvelteComponentTyped<TableRowProps, TableRowEvents, TableRowSlots> {
}
export {};
