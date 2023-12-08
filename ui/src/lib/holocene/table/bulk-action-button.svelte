<script>export let dataCy = null;
export let disabled = false;
export let variant = 'primary';
</script>

<button
  data-cy={dataCy}
  {disabled}
  on:click
  class="bulk-action-button {variant} {$$props.class}"
>
  <slot />
</button>

<style>
  .bulk-action-button {

    border-radius: 0.25rem;

    border-width: 1px;

    padding-left: 0.5rem;

    padding-right: 0.5rem;

    padding-top: 0.25rem;

    padding-bottom: 0.25rem;

    font-size: 0.75rem;

    line-height: .75rem;

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .bulk-action-button.primary {

    --tw-border-opacity: 1;

    border-color: rgb(255 255 255 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity))
}

  .bulk-action-button.primary:hover {

    --tw-bg-opacity: 1;

    background-color: rgb(29 78 216 / var(--tw-bg-opacity))
}

  .bulk-action-button.destructive {

    --tw-border-opacity: 1;

    border-color: rgb(185 28 28 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(185 28 28 / var(--tw-bg-opacity))
}

  .bulk-action-button.destructive:hover {

    --tw-border-opacity: 1;

    border-color: rgb(255 255 255 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(29 78 216 / var(--tw-bg-opacity))
}

  .bulk-action-button:disabled {

    pointer-events: none;

    cursor: not-allowed;

    opacity: 0.75
}</style>
