import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        endpoint?: string;
        passToken?: boolean;
        error?: string;
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {};
};
export declare type CodecEndpointSettingsProps = typeof __propDef.props;
export declare type CodecEndpointSettingsEvents = typeof __propDef.events;
export declare type CodecEndpointSettingsSlots = typeof __propDef.slots;
export default class CodecEndpointSettings extends SvelteComponentTyped<CodecEndpointSettingsProps, CodecEndpointSettingsEvents, CodecEndpointSettingsSlots> {
}
export {};
