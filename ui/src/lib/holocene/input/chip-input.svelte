<script>import { onDestroy, afterUpdate } from 'svelte';
import { writable } from 'svelte/store';
import Chip from '$holocene/chip.svelte';
export let id;
export let chips;
export let label = '';
export let placeholder = '';
export let name = id;
export let disabled = false;
export let required = false;
export let hintText = '';
export let validator = () => true;
const values = writable(chips);
let displayValue = '';
let shouldScrollToInput = false;
let inputContainer;
let input;
$: invalid = $values.some((chip) => !validator(chip));
const scrollToInput = () => {
    let rect = input.getBoundingClientRect();
    inputContainer.scrollTo(rect.x, rect.y);
    shouldScrollToInput = false;
};
const unsubscribe = values.subscribe((updatedChips) => {
    shouldScrollToInput = updatedChips.length > chips.length;
    chips = updatedChips;
});
afterUpdate(() => {
    if (shouldScrollToInput) {
        scrollToInput();
    }
});
onDestroy(() => {
    unsubscribe();
});
const handleKeydown = (e) => {
    if ((e.key === ',' || e.key === 'Enter') && displayValue !== '') {
        e.preventDefault();
        values.update((previous) => [...previous, displayValue]);
        displayValue = '';
    }
    const eventTarget = e.target;
    if (e.key === 'Backspace' &&
        eventTarget &&
        eventTarget.value === '' &&
        $values.length > 0) {
        values.update((previous) => previous.slice(0, -1));
    }
};
const handlePaste = (e) => {
    e.preventDefault();
    const clipboardContents = e.clipboardData.getData('text/plain');
    values.update((previous) => [...previous, ...clipboardContents.split(',')]);
};
const handleBlur = () => {
    if (displayValue !== '') {
        values.update((previous) => [...previous, displayValue]);
        displayValue = '';
    }
};
const removeChip = (index) => {
    values.update((previous) => {
        previous.splice(index, 1);
        return previous;
    });
};
</script>

<label class={$$props.class} for={id}>
  {#if label}
    {label}{#if required}*{/if}
  {/if}
  <div bind:this={inputContainer} class="input-container" class:invalid>
    {#if $values.length > 0}
      {#each $values as chip, i (`${chip}-${i}`)}
        {@const valid = validator(chip)}
        <Chip
          on:remove={() => removeChip(i)}
          intent={valid ? 'default' : 'warning'}>{chip}</Chip
        >
      {/each}
    {/if}
    <input
      data-lpignore="true"
      autocomplete="off"
      {disabled}
      {placeholder}
      {id}
      {name}
      {required}
      multiple
      bind:this={input}
      bind:value={displayValue}
      on:blur={handleBlur}
      on:keydown={handleKeydown}
      on:paste={handlePaste}
    />
  </div>
  {#if invalid && hintText}
    <span class="hint">
      {hintText}
    </span>
  {/if}
</label>

<style>
  label {

    margin-bottom: 2.5rem;

    font-size: 0.875rem;

    line-height: 1.25rem;

    font-weight: 500;

    --tw-text-opacity: 1;

    color: rgb(24 24 27 / var(--tw-text-opacity))
}

  .input-container {

    display: flex;

    max-height: 5rem;

    min-height: 2.5rem;

    width: 100%;

    flex-direction: row;

    flex-wrap: wrap;

    gap: 0.25rem;

    overflow-y: scroll;

    border-radius: 0.25rem;

    border-width: 1px;

    --tw-border-opacity: 1;

    border-color: rgb(24 24 27 / var(--tw-border-opacity));

    --tw-bg-opacity: 1;

    background-color: rgb(255 255 255 / var(--tw-bg-opacity));

    padding: 0.5rem;

    font-size: 0.875rem;

    line-height: 1.25rem;

    --tw-text-opacity: 1;

    color: rgb(24 24 27 / var(--tw-text-opacity))
}

  .input-container:focus-within {

    --tw-border-opacity: 1;

    border-color: rgb(29 78 216 / var(--tw-border-opacity))
}

  .input-container .invalid {

    --tw-border-opacity: 1;

    border-color: rgb(185 28 28 / var(--tw-border-opacity))
}

  input {

    display: inline-block;

    width: 100%;

    border-radius: 0.25rem;

    --tw-bg-opacity: 1;

    background-color: rgb(255 255 255 / var(--tw-bg-opacity))
}

  input:focus {

    outline: 2px solid transparent;

    outline-offset: 2px
}

  .hint {

    font-size: 0.75rem;

    line-height: 1rem;

    --tw-text-opacity: 1;

    color: rgb(185 28 28 / var(--tw-text-opacity))
}</style>
