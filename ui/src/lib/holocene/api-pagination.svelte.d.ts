import { SvelteComponentTyped } from "svelte";
declare class __sveltets_Render<T> {
    props(): {
        onError: (error: any) => void;
        onFetch: () => Promise<(size: number, token: string) => Promise<{
            items: any[];
            nextPageToken: string;
        }>>;
    };
    events(): {} & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        error: {};
        empty: {};
        'action-top-left': {};
        'action-top-center': {};
        'action-top-right': {};
        default: {
            visibleItems: any[];
            initialItem: any[];
        };
        'action-bottom-left': {};
        'action-bottom-right': {};
    };
}
export declare type ApiPaginationProps<T> = ReturnType<__sveltets_Render<T>['props']>;
export declare type ApiPaginationEvents<T> = ReturnType<__sveltets_Render<T>['events']>;
export declare type ApiPaginationSlots<T> = ReturnType<__sveltets_Render<T>['slots']>;
export default class ApiPagination<T> extends SvelteComponentTyped<ApiPaginationProps<T>, ApiPaginationEvents<T>, ApiPaginationSlots<T>> {
}
export {};
