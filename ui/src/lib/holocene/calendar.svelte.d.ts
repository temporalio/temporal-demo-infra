import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        date: Date | undefined;
        month: number | undefined;
        year: number | undefined;
        isAllowed?: (_date: Date) => boolean;
    };
    events: {
        datechange: CustomEvent<any>;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type CalendarProps = typeof __propDef.props;
export declare type CalendarEvents = typeof __propDef.events;
export declare type CalendarSlots = typeof __propDef.slots;
export default class Calendar extends SvelteComponentTyped<CalendarProps, CalendarEvents, CalendarSlots> {
}
export {};
