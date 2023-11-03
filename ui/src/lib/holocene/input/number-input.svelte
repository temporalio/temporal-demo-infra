<script>export let id;
export let value;
export let label = '';
export let units = '';
export let placeholder = '';
export let name = id;
export let disabled = false;
export let theme = 'light';
export let autocomplete = false;
export let hintText = '';
export let max = null;
export let spellcheck = null;
let valid = true;
function validateNumber(value) {
    if (!value.match('^[0-9]+$')) {
        valid = false;
        hintText = max ? `Enter a number between 1 - ${max}` : 'Enter a number';
    }
    else {
        greatThanMax(value);
    }
}
function greatThanMax(value) {
    if (max && parseInt(value) > max) {
        valid = false;
        hintText = `Enter a number between 1 - ${max}`;
    }
    else {
        setValidAndClearHint();
    }
}
function setValidAndClearHint() {
    valid = true;
    hintText = '';
}
$: {
    if (value) {
        validateNumber(value);
    }
    else {
        setValidAndClearHint();
    }
}
</script>

<div class={$$props.class}>
  {#if label}
    <label for={id}>{label}</label>
  {/if}
  <div class="flex items-center gap-2">
    <div class="input-container {theme}" class:disabled class:invalid={!valid}>
      <input
        class="m-2 block w-full bg-white text-center focus:outline-none"
        class:disabled
        {disabled}
        data-lpignore="true"
        {placeholder}
        {id}
        {name}
        autocomplete={autocomplete ? 'on' : 'off'}
        {spellcheck}
        bind:value
        on:input
        on:change
        on:focus
        on:blur
      />
    </div>
    <div class="units">{units}</div>
  </div>
</div>
{#if hintText}
  <span class="mt-1 text-xs text-red-700">{hintText}</span>
{/if}

<style>
  /* Base styles */
  label {
    margin-bottom: 2.5rem;
    font-family: Poppins, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500
}

  .units {
    font-family: Poppins, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500
}

  .input-container {
    position: relative;
    box-sizing: border-box;
    display: flex;
    height: 2.5rem;
    width: 4rem;
    align-items: center;
    border-radius: 0.25rem;
    border-width: 1px;
    --tw-border-opacity: 1;
    border-color: rgb(24 24 27 / var(--tw-border-opacity));
    font-size: 0.875rem;
    line-height: 1.25rem
}

  .input-container:focus-within {
    --tw-border-opacity: 1;
    border-color: rgb(29 78 216 / var(--tw-border-opacity))
}

  .input-container.disabled {
    border-width: 1px
}

  .icon-container {
    margin-left: 0.5rem;
    display: flex;
    align-items: center;
    justify-content: center
}

  .copy-icon-container {
    display: flex;
    height: 100%;
    width: 2.25rem;
    cursor: pointer;
    align-items: center;
    justify-content: center;
    border-top-right-radius: 0.25rem;
    border-bottom-right-radius: 0.25rem;
    border-left-width: 1px
}

  .input-container.invalid {
    --tw-border-opacity: 1;
    border-color: rgb(185 28 28 / var(--tw-border-opacity));
    --tw-text-opacity: 1;
    color: rgb(185 28 28 / var(--tw-text-opacity))
}

  .count {
    visibility: hidden;
    margin-right: 0.5rem;
    font-family: Poppins, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500;
    --tw-text-opacity: 1;
    color: rgb(24 24 27 / var(--tw-text-opacity))
}

  /* Light theme styles */
  .input-container.light, 
  .input-container.light .icon-container, 
  .input-container.light input {
    --tw-bg-opacity: 1;
    background-color: rgb(255 255 255 / var(--tw-bg-opacity))
}

  .input-container.light .icon-container {
    --tw-text-opacity: 1;
    color: rgb(161 161 170 / var(--tw-text-opacity))
}

  .input-container.light.disabled {
    --tw-border-opacity: 1;
    border-color: rgb(82 82 91 / var(--tw-border-opacity));
    --tw-bg-opacity: 1;
    background-color: rgb(250 250 250 / var(--tw-bg-opacity));
    --tw-text-opacity: 1;
    color: rgb(82 82 91 / var(--tw-text-opacity))
}

  .input-container.light.disabled input {
    --tw-bg-opacity: 1;
    background-color: rgb(250 250 250 / var(--tw-bg-opacity))
}

  .input-container.light.disabled .copy-icon-container {
    --tw-border-opacity: 1;
    border-color: rgb(82 82 91 / var(--tw-border-opacity));
    --tw-bg-opacity: 1;
    background-color: rgb(228 228 231 / var(--tw-bg-opacity))
}

  /* Dark theme styles */
  .input-container.dark, 
  .input-container.dark .icon-container, 
  .input-container.dark input, 
  .input-container.dark .copy-icon-container {
    --tw-bg-opacity: 1;
    background-color: rgb(24 24 27 / var(--tw-bg-opacity));
    --tw-text-opacity: 1;
    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .input-container.dark input::-moz-placeholder {
    --tw-text-opacity: 1;
    color: rgb(228 228 231 / var(--tw-text-opacity))
}

  .input-container.dark input:-ms-input-placeholder {
    --tw-text-opacity: 1;
    color: rgb(228 228 231 / var(--tw-text-opacity))
}

  .input-container.dark input::placeholder {
    --tw-text-opacity: 1;
    color: rgb(228 228 231 / var(--tw-text-opacity))
}

  .input-container.dark.disabled, 
  .input-container.dark.disabled .copy-icon-container, 
  .input-container.dark.disabled input {
    --tw-border-opacity: 1;
    border-color: rgb(24 24 27 / var(--tw-border-opacity));
    --tw-bg-opacity: 1;
    background-color: rgb(24 24 27 / var(--tw-bg-opacity))
}</style>
