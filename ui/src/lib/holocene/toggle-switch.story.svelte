<script>import ToggleSwitch from './toggle-switch.svelte';
export let Hst;
let checked = false;
</script>

<Hst.Story>
  <ToggleSwitch id="toggle-1" bind:checked />
</Hst.Story>
