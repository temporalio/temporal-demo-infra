<script context="module">import { writable } from 'svelte/store';
import { v4 } from 'uuid';
const toasts = writable([]);
const push = (toast) => {
    const toastWithDefaults = {
        id: v4(),
        duration: 3000,
        xPosition: 'right',
        yPosition: 'top',
        variant: 'primary',
        ...toast,
    };
    toasts.update((ts) => [...ts, toastWithDefaults]);
    const timeoutId = setTimeout(() => {
        pop(toastWithDefaults.id);
        clearTimeout(timeoutId);
    }, toastWithDefaults.duration);
};
const pop = (id) => {
    toasts.update((ts) => ts.filter((t) => t.id !== id));
};
export const toaster = {
    push,
    pop,
    toasts,
};
</script>

<script>import ToastComponent from './toast.svelte';
export let pop;
export let toasts;
const dismissToast = (event) => {
    pop(event.detail.id);
};
$: topRightToasts = $toasts.filter((toast) => toast.yPosition === 'top' && toast.xPosition === 'right');
$: bottomRightToasts = $toasts
    .filter((toast) => toast.yPosition === 'bottom' && toast.xPosition === 'right')
    .reverse();
$: bottomLeftToasts = $toasts
    .filter((toast) => toast.yPosition === 'bottom' && toast.xPosition === 'left')
    .reverse();
$: topLeftToasts = $toasts.filter((toast) => toast.yPosition === 'top' && toast.xPosition === 'left');
</script>

<div class="toast-container top-5 right-5">
  {#each topRightToasts as { message, variant, id } (id)}
    <ToastComponent {variant} {id} on:dismiss={dismissToast}>
      {message}
    </ToastComponent>
  {/each}
</div>
<div class="toast-container bottom-5 right-5">
  {#each bottomRightToasts as { message, variant, id } (id)}
    <ToastComponent {variant} {id} on:dismiss={dismissToast}>
      {message}
    </ToastComponent>
  {/each}
</div>
<div class="toast-container bottom-5 left-5">
  {#each bottomLeftToasts as { message, variant, id } (id)}
    <ToastComponent {variant} {id} on:dismiss={dismissToast}>
      {message}
    </ToastComponent>
  {/each}
</div>
<div class="toast-container top-5 left-5">
  {#each topLeftToasts as { message, variant, id } (id)}
    <ToastComponent {variant} {id} on:dismiss={dismissToast}>
      {message}
    </ToastComponent>
  {/each}
</div>

<style>
  .toast-container {
    position: fixed;
    z-index: 50;
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 0.5rem
}</style>
