import { SvelteComponentTyped } from "svelte";
import type { IconName } from '../../../holocene/icon/paths';
declare const __propDef: {
    props: {
        [x: string]: any;
        id: string;
        value: string;
        label?: string;
        icon?: IconName;
        placeholder?: string;
        suffix?: string;
        name?: string;
        copyable?: boolean;
        disabled?: boolean;
        clearable?: boolean;
        theme?: 'dark' | 'light';
        autocomplete?: boolean;
        valid?: boolean;
        hintText?: string;
        maxLength?: number;
        spellcheck?: boolean;
        unroundRight?: boolean;
        unroundLeft?: boolean;
        autoFocus?: boolean;
        error?: boolean;
    };
    events: {
        input: Event;
        change: Event;
        focus: FocusEvent;
        blur: FocusEvent;
        clear: CustomEvent<any>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type InputProps = typeof __propDef.props;
export declare type InputEvents = typeof __propDef.events;
export declare type InputSlots = typeof __propDef.slots;
export default class Input extends SvelteComponentTyped<InputProps, InputEvents, InputSlots> {
}
export {};
