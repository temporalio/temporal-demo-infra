import { SvelteComponentTyped } from "svelte";
declare const __propDef: {
    props: {
        [x: string]: any;
        dataCy?: string;
        disabled?: boolean;
        variant?: 'primary' | 'destructive';
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
export declare type BulkActionButtonProps = typeof __propDef.props;
export declare type BulkActionButtonEvents = typeof __propDef.events;
export declare type BulkActionButtonSlots = typeof __propDef.slots;
export default class BulkActionButton extends SvelteComponentTyped<BulkActionButtonProps, BulkActionButtonEvents, BulkActionButtonSlots> {
}
export {};
