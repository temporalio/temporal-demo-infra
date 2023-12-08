import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        isAllowed?: (date: Date) => boolean;
        selected?: Date;
    };
    events: {
        datechange: CustomEvent<any>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type DatePickerProps = typeof __propDef.props;
export declare type DatePickerEvents = typeof __propDef.events;
export declare type DatePickerSlots = typeof __propDef.slots;
export default class DatePicker extends SvelteComponentTyped<DatePickerProps, DatePickerEvents, DatePickerSlots> {
}
export {};
