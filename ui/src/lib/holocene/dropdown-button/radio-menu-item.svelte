<script>import { createEventDispatcher } from 'svelte';
import MenuItem from '$holocene/primitives/menu/menu-item.svelte';
export let checked = false;
export let value;
const dispatch = createEventDispatcher();
function handleItemClick() {
    dispatch('select', { value });
}
</script>

<MenuItem active={checked} on:click={handleItemClick}>
  <label class="flex w-full cursor-pointer items-center">
    <input {checked} class="mr-4" type="radio" {value} />
    <slot />
  </label>
</MenuItem>
