<script>import { fly } from 'svelte/transition';
import { createEventDispatcher } from 'svelte';
import IconButton from './icon-button.svelte';
const dispatch = createEventDispatcher();
export let id;
export let variant;
const handleDismiss = () => {
    dispatch('dismiss', { id });
};
</script>

<article
  {id}
  class="flex grow-0 items-center justify-between gap-4 rounded py-3 px-4 shadow {variant}"
  transition:fly={{ x: 250 }}
>
  <p class="font-secondary text-sm">
    <slot />
  </p>
  <IconButton icon="close" on:click={handleDismiss} />
</article>

<style>
  .primary {

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .success {

    --tw-bg-opacity: 1;

    background-color: rgb(240 253 244 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(20 83 45 / var(--tw-text-opacity))
}

  .error {

    --tw-bg-opacity: 1;

    background-color: rgb(254 242 242 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(127 29 29 / var(--tw-text-opacity))
}

  .info {

    --tw-bg-opacity: 1;

    background-color: rgb(239 246 255 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(30 58 138 / var(--tw-text-opacity))
}

  .warning {

    --tw-bg-opacity: 1;

    background-color: rgb(254 252 232 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(113 63 18 / var(--tw-text-opacity))
}</style>
