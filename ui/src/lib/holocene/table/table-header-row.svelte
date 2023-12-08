<script>import Checkbox from '../checkbox.svelte';
export let selectable = false;
export let selected = false;
export let checkboxLabel = null;
export let indeterminate = false;
</script>

<tr class={$$props.class}>
  <th class:selectable>
    {#if selectable}
      <Checkbox
        id="selectable-table-check-all"
        onDark
        bind:checked={selected}
        {indeterminate}
        label={checkboxLabel}
        on:change
      />
    {/if}
  </th>
  <slot />
  <th class:selectable />
</tr>
