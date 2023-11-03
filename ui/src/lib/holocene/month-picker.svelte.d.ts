import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        months: string[];
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type MonthPickerProps = typeof __propDef.props;
export declare type MonthPickerEvents = typeof __propDef.events;
export declare type MonthPickerSlots = typeof __propDef.slots;
export default class MonthPicker extends SvelteComponentTyped<MonthPickerProps, MonthPickerEvents, MonthPickerSlots> {
}
export {};
