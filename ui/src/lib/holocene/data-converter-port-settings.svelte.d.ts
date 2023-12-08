import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        port?: string;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type DataConverterPortSettingsProps = typeof __propDef.props;
export declare type DataConverterPortSettingsEvents = typeof __propDef.events;
export declare type DataConverterPortSettingsSlots = typeof __propDef.slots;
export default class DataConverterPortSettings extends SvelteComponentTyped<DataConverterPortSettingsProps, DataConverterPortSettingsEvents, DataConverterPortSettingsSlots> {
}
export {};
