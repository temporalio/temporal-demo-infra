import { SvelteComponentTyped } from "svelte";
declare class __sveltets_Render<T> {
    props(): {
        items: T[];
        floatId?: string;
        startingIndex?: string | number;
        perPageKey?: string;
        currentPageKey?: string;
        itemsPerPage?: number;
    };
    events(): {} & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        'action-top-left': {};
        'action-top-center': {};
        'action-top-right': {};
        default: {
            visibleItems: T[];
            initialItem: T;
        };
        'action-bottom-left': {};
        'action-bottom-right': {};
    };
}
export declare type PaginationProps<T> = ReturnType<__sveltets_Render<T>['props']>;
export declare type PaginationEvents<T> = ReturnType<__sveltets_Render<T>['events']>;
export declare type PaginationSlots<T> = ReturnType<__sveltets_Render<T>['slots']>;
export default class Pagination<T> extends SvelteComponentTyped<PaginationProps<T>, PaginationEvents<T>, PaginationSlots<T>> {
}
export {};
