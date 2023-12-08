import { SvelteComponentTyped } from "svelte";
import type { IconName } from '../../../holocene/icon/paths';
declare const __propDef: {
    props: {
        [x: string]: any;
        id: string;
        options?: {
            label: string;
            value: string;
        }[];
        placeholder?: string;
        icon?: IconName;
        autoFocus?: boolean;
        unroundRight?: boolean;
        unroundLeft?: boolean;
        onChange?: (value: string) => void;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type TypeaheadInputProps = typeof __propDef.props;
export declare type TypeaheadInputEvents = typeof __propDef.events;
export declare type TypeaheadInputSlots = typeof __propDef.slots;
export default class TypeaheadInput extends SvelteComponentTyped<TypeaheadInputProps, TypeaheadInputEvents, TypeaheadInputSlots> {
}
export {};
