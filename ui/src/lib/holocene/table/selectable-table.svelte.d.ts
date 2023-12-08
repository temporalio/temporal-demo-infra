import { SvelteComponentTyped } from "svelte";
export declare type SelectableTableContext<T> = {
    handleSelectRow: (event: CustomEvent<{
        checked: boolean;
    }>, item: T) => void;
};
declare class __sveltets_Render<T> {
    props(): {
        variant?: "simple" | "fancy";
        updating?: boolean;
        id?: string;
        items: (T & {
            id: string;
        })[];
        selectedItems: (T & {
            id: string;
        })[];
        allSelected?: boolean;
        checkboxLabel?: string;
        class?: string;
    };
    events(): {} & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        'bulk-action-headers': {};
        'default-headers': {};
        default: {};
    };
}
export declare type SelectableTableProps<T> = ReturnType<__sveltets_Render<T>['props']>;
export declare type SelectableTableEvents<T> = ReturnType<__sveltets_Render<T>['events']>;
export declare type SelectableTableSlots<T> = ReturnType<__sveltets_Render<T>['slots']>;
export default class SelectableTable<T> extends SvelteComponentTyped<SelectableTableProps<T>, SelectableTableEvents<T>, SelectableTableSlots<T>> {
}
export {};
