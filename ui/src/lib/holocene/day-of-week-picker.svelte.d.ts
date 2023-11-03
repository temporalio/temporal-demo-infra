import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        daysOfWeek: string[];
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type DayOfWeekPickerProps = typeof __propDef.props;
export declare type DayOfWeekPickerEvents = typeof __propDef.events;
export declare type DayOfWeekPickerSlots = typeof __propDef.slots;
export default class DayOfWeekPicker extends SvelteComponentTyped<DayOfWeekPickerProps, DayOfWeekPickerEvents, DayOfWeekPickerSlots> {
}
export {};
