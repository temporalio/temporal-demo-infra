import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        id: string;
        value: string;
        label?: string;
        units?: string;
        placeholder?: string;
        name?: string;
        disabled?: boolean;
        theme?: 'dark' | 'light';
        autocomplete?: boolean;
        hintText?: string;
        max?: null | number;
        spellcheck?: boolean;
    };
    events: {
        input: Event;
        change: Event;
        focus: FocusEvent;
        blur: FocusEvent;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type NumberInputProps = typeof __propDef.props;
export declare type NumberInputEvents = typeof __propDef.events;
export declare type NumberInputSlots = typeof __propDef.slots;
export default class NumberInput extends SvelteComponentTyped<NumberInputProps, NumberInputEvents, NumberInputSlots> {
}
export {};
