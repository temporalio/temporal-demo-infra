<script>import Icon from '$holocene/icon/icon.svelte';
import MenuContainer from '$holocene/primitives/menu/menu-container.svelte';
import MenuButton from '$holocene/primitives/menu/menu-button.svelte';
import Menu from '$holocene/primitives/menu/menu.svelte';
export let label = '';
export let icon = undefined;
export let id;
export let disabled = false;
export let position = 'left';
let show = false;
</script>

<MenuContainer class={$$props.class}>
  <div class="split-button" class:disabled>
    <slot name="middle-button" />
    <MenuButton
      class="segment right {$$props.buttonClass}"
      bind:show
      controls={id}
      {disabled}
    >
      {#if icon}
        <Icon name={icon} class="mr-2" />
      {/if}
      {label}
      <Icon name="chevron-down" class="ml-2" />
    </MenuButton>
  </div>
  <Menu class="min-w-max" {id} {show} {position}>
    <slot />
  </Menu>
</MenuContainer>

<style>
  .split-button {

    display: flex;

    height: 2.5rem;

    flex-grow: 1;

    cursor: pointer;

    flex-direction: row;

    gap: 0px;

    font-family: Poppins, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"
}

    .split-button :global(.segment) {

    position: relative;

    display: flex;

    width: -webkit-fit-content;

    width: -moz-fit-content;

    width: fit-content;

    align-items: center;

    justify-content: center;

    padding-top: 0.5rem;

    padding-bottom: 0.5rem;

    font-family: Poppins, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";

    font-size: 0.875rem;

    line-height: 1.25rem
}

    .split-button :global(.segment.left) {

    border-radius: 0px;

    border-top-left-radius: 0.25rem;

    border-bottom-left-radius: 0.25rem;

    padding-left: 0.5rem;

    padding-right: 0.5rem
}

    .split-button :global(.segment.right) {

    border-top-right-radius: 0.25rem;

    border-bottom-right-radius: 0.25rem;

    padding-left: 0.5rem;

    padding-right: 0.5rem
}

    .split-button :global(.segment.right:disabled) {

    cursor: not-allowed;

    opacity: 0.5
}</style>
