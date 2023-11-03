<script>import Icon from './icon.svelte';
import { icons } from './paths';
export let Hst;
export let size = 24;
const iconNames = Object.keys(icons);
</script>

<Hst.Story>
  {#each iconNames as icon}
    <Hst.Variant title={icon}>
      <Icon name={icon} height={size} width={size} />
    </Hst.Variant>
  {/each}

  <svelte:fragment slot="controls">
    <Hst.Number title="Icon Size:" bind:value={size} />
  </svelte:fragment>
</Hst.Story>
