<script>import { createEventDispatcher } from 'svelte';
import Icon from '$holocene/icon/icon.svelte';
export let id = '';
export let checked = false;
export let label = null;
export let onDark = false;
export let indeterminate = false;
export let disabled = false;
const dispatch = createEventDispatcher();
const handleChange = (event) => {
    dispatch('change', { checked: event.target.checked });
};
</script>

<label
  on:click
  on:keypress
  class="checkbox {$$props.class}"
  class:disabled
  class:on-dark={onDark}
>
  <span class="label">
    {#if label}
      {@html label}
    {:else}
      &nbsp;
    {/if}
  </span>
  <input
    on:click|stopPropagation
    on:change={handleChange}
    {id}
    type="checkbox"
    bind:checked
    {indeterminate}
    {disabled}
    class:indeterminate
  />
  <span class="checkmark" class:on-dark={onDark}>
    {#if indeterminate}
      <Icon class="absolute top-0 left-0 h-4 w-4" name="hyphen" />
    {:else if checked}
      <Icon
        class="absolute top-0 left-0 h-4 w-4"
        name="checkmark"
        strokeWidth={3}
      />
    {/if}
  </span>
</label>

<style>
  .checkbox {

    display: flex;

    width: -webkit-fit-content;

    width: -moz-fit-content;

    width: fit-content;

    cursor: pointer;

    -webkit-user-select: none;

       -moz-user-select: none;

        -ms-user-select: none;

            user-select: none;

    align-items: center;

    font-size: 0.875rem;

    line-height: 1.5rem;

    --tw-text-opacity: 1;

    color: rgb(24 24 27 / var(--tw-text-opacity))
}

  .checkbox.on-dark {

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .label {

    position: absolute;

    margin-left: 1.5rem;

    display: flex;

    align-items: center;

    white-space: nowrap
}

  input {

    position: absolute;

    top: 0px;

    left: 0px;

    height: 0px;

    width: 0px;

    opacity: 0
}

  .checkmark {

    position: absolute;

    box-sizing: content-box;

    height: 1rem;

    width: 1rem;

    cursor: pointer;

    border-radius: 0.125rem;

    border-width: 1px;

    --tw-border-opacity: 1;

    border-color: rgb(113 113 122 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(255 255 255 / var(--tw-bg-opacity))
}

  .checkmark.on-dark {

    --tw-border-opacity: 1;

    border-color: rgb(255 255 255 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity))
}

  input:checked + .checkmark, 
  input.indeterminate + .checkmark {

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .checkbox.disabled, 
  .checkbox.disabled .checkmark {

    cursor: default
}

  .checkbox.disabled.on-dark {

    --tw-text-opacity: 0.8
}

  .checkbox.disabled:not(.on-dark) .checkmark {

    --tw-bg-opacity: 1;

    background-color: rgb(212 212 216 / var(--tw-bg-opacity))
}

  .checkbox.disabled.on-dark .checkmark {

    --tw-border-opacity: 0.8;

    --tw-text-opacity: 0.8
}</style>
