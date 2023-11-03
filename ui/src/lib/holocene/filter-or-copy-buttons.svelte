<script>import Icon from '$holocene/icon/icon.svelte';
import { copyToClipboard } from '../utilities/copy-to-clipboard';
import { noop } from 'svelte/internal';
export let show = false;
export let filterable = true;
export let copyable = true;
export let content;
export let onFilter = noop;
export let filtered = false;
let className = '';
export { className as class };
const { copy, copied } = copyToClipboard(700);
</script>

{#if show}
  <div
    class="copy-or-filter {className}"
    on:click|preventDefault|stopPropagation={noop}
  >
    {#if filterable}
      <button on:click|preventDefault|stopPropagation={onFilter}>
        {#key filtered}
          <Icon
            name="filter"
            class="h-4 w-4 rounded-sm {filtered
              ? 'bg-gray-900 text-white'
              : ''}"
          />
        {/key}
      </button>
    {/if}
    {#if copyable}
      <button on:click|preventDefault|stopPropagation={(e) => copy(e, content)}>
        <Icon name={$copied ? 'checkmark' : 'copy'} stroke="#000" />
      </button>
    {/if}
  </div>
{/if}

<style>
  .copy-or-filter {
    position: absolute;
    right: 0px;
    top: 0px;
    bottom: 0px;
    display: inline-flex;
    gap: 0.5rem;
    padding-left: 0.5rem;
    padding-right: 0.5rem
}</style>
