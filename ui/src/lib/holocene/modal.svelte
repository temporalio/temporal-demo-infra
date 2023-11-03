<script>import Icon from '$holocene/icon/icon.svelte';
import { createEventDispatcher } from 'svelte';
import Button from '$holocene/button.svelte';
export let open = false;
export let hideConfirm = false;
export let confirmText = 'Confirm';
export let cancelText = 'Cancel';
export let confirmType = 'primary';
export let confirmDisabled = false;
export let large = false;
export let loading = false;
let modalElement;
const dispatch = createEventDispatcher();
const cancelModal = () => {
    dispatch('cancelModal');
};
const handleKeyboardNavigation = (event) => {
    if (!open) {
        return;
    }
    if (event.key === 'Escape') {
        cancelModal();
        return;
    }
    const focusable = modalElement.querySelectorAll('button');
    const firstFocusable = focusable[0];
    const lastFocusable = focusable[focusable.length - 1];
    if (event.key === 'Tab') {
        if (event.shiftKey) {
            if (document.activeElement === firstFocusable) {
                lastFocusable.focus();
                event.preventDefault();
            }
        }
        else if (document.activeElement === lastFocusable) {
            firstFocusable.focus();
            event.preventDefault();
        }
    }
};
$: {
    if (open && modalElement) {
        modalElement.focus();
    }
}
</script>

<svelte:window on:keydown={handleKeyboardNavigation} />
{#if open}
  <div class="modal">
    <div on:click={cancelModal} class="overlay" />
    <div
      bind:this={modalElement}
      class="body"
      class:large
      tabindex="-1"
      role="alertdialog"
      aria-labelledby="modal-title"
      aria-describedby="modal-description"
    >
      {#if !loading}
        <button
          aria-label={cancelText}
          class="float-right m-4"
          on:click={cancelModal}
        >
          <Icon
            name="close"
            class="cursor-pointer rounded-full hover:bg-gray-900 hover:text-white"
          />
        </button>
      {/if}
      <div id="modal-title" class="title">
        <slot name="title">
          <h3>Title</h3>
        </slot>
      </div>
      <div id="modal-content" class="content">
        <slot name="content">
          <span>Content</span>
        </slot>
      </div>
      <div class="flex items-center justify-end space-x-2 p-6">
        <Button
          thin
          variant="secondary"
          disabled={loading}
          on:click={cancelModal}>{cancelText}</Button
        >
        {#if !hideConfirm}
          <Button
            thin
            variant={confirmType}
            {loading}
            disabled={confirmDisabled || loading}
            dataCy="confirm-modal-button"
            on:click={() => dispatch('confirmModal')}>{confirmText}</Button
          >
        {/if}
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {

    position: fixed;

    top: 0px;

    left: 0px;

    z-index: 50;

    display: flex;

    height: 100%;

    width: 100%;

    cursor: default;

    align-items: center;

    justify-content: center;

    padding: 2rem
}

@media (min-width: 1024px) {

    .modal {

        padding: 0px
    }
}

  .overlay {

    position: fixed;

    height: 100%;

    width: 100%;

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity));

    opacity: 0.5
}

  .body {

    z-index: 50;

    margin-left: auto;

    margin-right: auto;

    width: 100%;

    max-width: 32rem;

    overflow-y: auto;

    border-radius: 0.5rem;

    --tw-bg-opacity: 1;

    background-color: rgb(255 255 255 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(24 24 27 / var(--tw-text-opacity));

    --tw-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);

    --tw-shadow-colored: 0 20px 25px -5px var(--tw-shadow-color), 0 8px 10px -6px var(--tw-shadow-color);

    box-shadow: var(--tw-ring-offset-shadow, 0 0 #0000), var(--tw-ring-shadow, 0 0 #0000), var(--tw-shadow)
}

  @media (min-width: 768px) {

    .body {

        height: -webkit-max-content;

        height: -moz-max-content;

        height: max-content
    }
}

  @media (min-width: 1024px) {

    .large {

        width: 50%
    }
}

  .title {

    --tw-bg-opacity: 1;

    background-color: rgb(255 255 255 / var(--tw-bg-opacity));

    padding-left: 2rem;

    padding-right: 2rem;

    padding-top: 2rem;

    padding-bottom: 0px;

    font-size: 1.5rem;

    line-height: 2rem
}

  .content {

    white-space: normal;

    padding: 2rem
}</style>
