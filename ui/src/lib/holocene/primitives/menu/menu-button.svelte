<script>import Icon from '../../icon/icon.svelte';
import { triggerMenu } from './trigger-menu';
export let show;
export let controls;
export let dark = false;
export let disabled = false;
export let hasIndicator = false;
export let keepOpen = false;
export let id = null;
const close = () => {
    !disabled && (show = false);
};
const toggle = () => {
    !disabled && (show = !show);
};
</script>

<button
  type="button"
  {id}
  aria-haspopup={!disabled}
  aria-controls={controls}
  aria-expanded={show}
  use:triggerMenu={keepOpen}
  on:close-menu={close}
  on:toggle-menu={toggle}
  on:click|preventDefault
  class={$$props.class}
  class:dark
  class:show
  {disabled}
>
  <slot />
  {#if hasIndicator}
    <Icon
      class="pointer-events-none"
      name={show ? 'chevron-up' : 'chevron-down'}
    />
  {/if}
</button>

<style>
  button.dark, 
  button.dark > * {

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  button.disabled {

    --tw-bg-opacity: 1;

    background-color: rgb(250 250 250 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(82 82 91 / var(--tw-text-opacity))
}</style>
