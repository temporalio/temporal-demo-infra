import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        selectable?: boolean;
        selected?: boolean;
        checkboxLabel?: string;
        indeterminate?: boolean;
    };
    events: {
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
export declare type TableHeaderRowProps = typeof __propDef.props;
export declare type TableHeaderRowEvents = typeof __propDef.events;
export declare type TableHeaderRowSlots = typeof __propDef.slots;
export default class TableHeaderRow extends SvelteComponentTyped<TableHeaderRowProps, TableHeaderRowEvents, TableHeaderRowSlots> {
}
export {};
