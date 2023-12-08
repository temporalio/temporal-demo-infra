import { SvelteComponentTyped } from "svelte";
declare class __sveltets_Render<T> {
    props(): {
        checked?: boolean;
        value: T;
    };
    events(): {
        select: CustomEvent<{
            value: T;
        }>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots(): {
        default: {};
    };
}
export declare type RadioMenuItemProps<T> = ReturnType<__sveltets_Render<T>['props']>;
export declare type RadioMenuItemEvents<T> = ReturnType<__sveltets_Render<T>['events']>;
export declare type RadioMenuItemSlots<T> = ReturnType<__sveltets_Render<T>['slots']>;
export default class RadioMenuItem<T> extends SvelteComponentTyped<RadioMenuItemProps<T>, RadioMenuItemEvents<T>, RadioMenuItemSlots<T>> {
}
export {};
