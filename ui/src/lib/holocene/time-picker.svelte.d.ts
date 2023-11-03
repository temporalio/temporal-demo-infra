import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        hour?: string;
        minute?: string;
        second?: string;
        half?: 'AM' | 'PM';
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type TimePickerProps = typeof __propDef.props;
export declare type TimePickerEvents = typeof __propDef.events;
export declare type TimePickerSlots = typeof __propDef.slots;
export default class TimePicker extends SvelteComponentTyped<TimePickerProps, TimePickerEvents, TimePickerSlots> {
}
export {};
