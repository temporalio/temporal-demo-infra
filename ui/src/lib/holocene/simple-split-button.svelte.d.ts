import { SvelteComponentTyped } from "svelte";
import type { IconName } from './icon/paths';
declare const __propDef: {
    props: {
        [x: string]: any;
        label?: string;
        icon?: IconName | undefined;
        id: string;
        disabled?: boolean;
        position?: 'left' | 'right';
    };
    events: {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        'middle-button': {};
        default: {};
    };
};
export declare type SimpleSplitButtonProps = typeof __propDef.props;
export declare type SimpleSplitButtonEvents = typeof __propDef.events;
export declare type SimpleSplitButtonSlots = typeof __propDef.slots;
export default class SimpleSplitButton extends SvelteComponentTyped<SimpleSplitButtonProps, SimpleSplitButtonEvents, SimpleSplitButtonSlots> {
}
export {};
