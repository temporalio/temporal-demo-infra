import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        daysOfMonth: number[];
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type DayOfMonthPickerProps = typeof __propDef.props;
export declare type DayOfMonthPickerEvents = typeof __propDef.events;
export declare type DayOfMonthPickerSlots = typeof __propDef.slots;
export default class DayOfMonthPicker extends SvelteComponentTyped<DayOfMonthPickerProps, DayOfMonthPickerEvents, DayOfMonthPickerSlots> {
}
export {};
