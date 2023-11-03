import { SvelteComponentTyped } from "svelte";
declare class __sveltets_Render<T> {
    props(): {
        href?: string;
        selectable?: boolean;
        selected: boolean;
        item: T & {
            id: string;
        };
        class?: string;
    };
    events(): {
        click: MouseEvent;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        default: {};
    };
}
export declare type SelectableTableRowProps<T> = ReturnType<__sveltets_Render<T>['props']>;
export declare type SelectableTableRowEvents<T> = ReturnType<__sveltets_Render<T>['events']>;
export declare type SelectableTableRowSlots<T> = ReturnType<__sveltets_Render<T>['slots']>;
export default class SelectableTableRow<T> extends SvelteComponentTyped<SelectableTableRowProps<T>, SelectableTableRowEvents<T>, SelectableTableRowSlots<T>> {
}
export {};
