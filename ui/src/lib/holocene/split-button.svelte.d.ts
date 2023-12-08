import { SvelteComponentTyped } from "svelte";
import type { IconName } from './icon/paths';
declare const __propDef: {
    props: {
        [x: string]: any;
        label?: string;
        icon?: IconName | undefined;
        id: string;
        variant?: 'primary' | 'secondary' | 'destructive' | 'login' | 'link';
        thin?: boolean;
        disabled?: boolean;
        position?: 'left' | 'right';
        href?: string;
        primaryActionDisabled?: boolean;
    };
    events: {
        click: MouseEvent;
    } & {
        [evt: string]: CustomEvent<any>;
    };
    slots: {
        default: {};
    };
};
export declare type SplitButtonProps = typeof __propDef.props;
export declare type SplitButtonEvents = typeof __propDef.events;
export declare type SplitButtonSlots = typeof __propDef.slots;
export default class SplitButton extends SvelteComponentTyped<SplitButtonProps, SplitButtonEvents, SplitButtonSlots> {
}
export {};
