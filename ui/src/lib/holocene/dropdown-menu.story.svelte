<script>import DropdownMenu from './dropdown-menu.svelte';
import Icon from '$holocene/icon/icon.svelte';
import MenuItem from './primitives/menu/menu-item.svelte';
export let Hst;
</script>

<Hst.Story>
  <DropdownMenu id="dropdown-1">
    <Icon name="vertical-ellipsis" slot="trigger" />
    <div class="w-56" slot="items">
      <MenuItem href="https://temporal.io">Link</MenuItem>
      <MenuItem disabled href="https://temporal.io">Disabled Link</MenuItem>
      <MenuItem selected>Selected</MenuItem>
      <MenuItem>Standard</MenuItem>
      <MenuItem destructive>Destructive</MenuItem>
    </div>
  </DropdownMenu>
</Hst.Story>
