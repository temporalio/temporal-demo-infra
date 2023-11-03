<script>import Checkbox from '../checkbox.svelte';
export let href = '';
export let selectable = false;
export let selected = false;
</script>

{#if href}
  <a class="table-row align-middle {$$props.class}" {href}>
    <td class:selectable on:click|stopPropagation>
      {#if selectable}
        <div class="absolute">
          <Checkbox bind:checked={selected} on:change />
        </div>
      {/if}
    </td>
    <slot />
    <td />
  </a>
{:else}
  <tr on:click class={$$props.class}>
    <td class:selectable>
      {#if selectable}
        <div class="absolute">
          <Checkbox bind:checked={selected} on:change />
        </div>
      {/if}
    </td>
    <slot />
    <td />
  </tr>
{/if}
