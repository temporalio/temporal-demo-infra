<script>import { writable } from 'svelte/store';
import { onDestroy, setContext } from 'svelte';
import { clickOutside } from '../outside-click';
import Input from './input.svelte';
import Menu from '$holocene/primitives/menu/menu.svelte';
import MenuContainer from '$holocene/primitives/menu/menu-container.svelte';
import { noop } from 'svelte/internal';
import Option from '../select/option.svelte';
export let id;
export let options = [];
export let placeholder = '';
export let icon = null;
export let autoFocus = false;
export let unroundRight = false;
export let unroundLeft = false;
export let onChange = noop;
let value = '';
let showMenu = false;
let filterOptions = options;
$: {
    if (value) {
        filterOptions = options.filter((o) => o.label.toLowerCase().includes(value.toLowerCase()));
    }
    else {
        filterOptions = options;
    }
}
const context = writable({
    selectValue: value,
    onChange: () => {
        onChange(value);
        showMenu = false;
    },
});
const unsubscribe = context.subscribe((ctx) => {
    value = ctx.selectValue;
});
onDestroy(() => {
    unsubscribe();
});
$: {
    if (value) {
        context.update((previous) => ({ ...previous, selectValue: value }));
    }
    setContext('select-value', context);
}
</script>

<div
  class="relative"
  use:clickOutside
  on:click-outside={() => (showMenu = false)}
>
  <MenuContainer class={$$props.class}>
    <Input
      {id}
      {icon}
      class={$$props.class}
      bind:value
      {placeholder}
      {autoFocus}
      {unroundRight}
      {unroundLeft}
      on:focus={() => (showMenu = true)}
    />
    <Menu show={showMenu} id={`menu-${id}`} class="h-auto max-h-80 w-64">
      {#each filterOptions as { label, value }}
        <Option {value}>{label}</Option>
      {:else}
        <Option>No Results</Option>
      {/each}
    </Menu>
  </MenuContainer>
</div>
