<script>import { getContext } from 'svelte';
import TableRow from './table-row.svelte';
export let item;
export let selected;
const { handleSelectRow } = getContext('selectable-table-context');
const handleChange = (event) => {
    handleSelectRow(event, item);
};
</script>

<TableRow
  on:change={handleChange}
  on:click
  selectable
  bind:selected
  {...$$props}
>
  <slot />
</TableRow>
