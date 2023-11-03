<script context="module"></script>

<script>import { setContext } from 'svelte';
import TableHeaderRow from './table-header-row.svelte';
import Table from './table.svelte';
export let items;
export let checkboxLabel = null;
export let allSelected = false;
export let selectedItems = [];
export let id = null;
const handleSelectAll = (event) => {
    allSelected = !allSelected;
    selectedItems = event.detail.checked ? items : [];
};
const handleSelectRow = (event, item) => {
    const { checked } = event.detail;
    if (checked) {
        selectedItems.push(item);
        selectedItems = selectedItems;
    }
    else {
        selectedItems = selectedItems.filter((i) => i.id !== item.id);
    }
    allSelected = selectedItems.length === items.length;
};
$: indeterminate =
    selectedItems.length !== 0 && selectedItems.length !== items.length;
setContext('selectable-table-context', {
    handleSelectRow,
});
</script>

<Table variant="fancy" {id} class={$$props.class} {...$$props}>
  <TableHeaderRow
    slot="headers"
    selectable
    checkboxLabel={selectedItems.length > 0 ? null : checkboxLabel}
    {indeterminate}
    on:change={handleSelectAll}
    bind:selected={allSelected}
  >
    {#if selectedItems.length > 0}
      <slot name="bulk-action-headers" />
    {:else}
      <slot name="default-headers" />
    {/if}
  </TableHeaderRow>
  <slot />
</Table>
