<script>import { monthNames } from '../utilities/calendar';
export let months;
const onClick = (e, month) => {
    if (e.metaKey && month !== '*') {
        if (months.includes('*')) {
            months = months.filter((m) => m !== '*');
        }
        if (months.includes(month)) {
            months = months.filter((m) => m !== month);
        }
        else {
            months = [...months, month];
        }
    }
    else {
        months = [month];
    }
};
</script>

<div class="flex flex-wrap gap-4 text-center">
  {#each monthNames as { label, value }}
    <button
      class="cell"
      class:active={months.includes(value)}
      on:click|preventDefault={(e) => onClick(e, value)}
    >
      {label}
    </button>
  {/each}
</div>

<style>
  .cell {

    cursor: pointer;

    border-radius: 0.25rem;

    border-width: 1px;

    --tw-bg-opacity: 1;

    background-color: rgb(212 212 216 / var(--tw-bg-opacity));

    padding: 0.5rem
}

.cell:hover {

    --tw-bg-opacity: 1;

    background-color: rgb(29 78 216 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}

  .active {

    --tw-bg-opacity: 1;

    background-color: rgb(24 24 27 / var(--tw-bg-opacity));

    --tw-text-opacity: 1;

    color: rgb(255 255 255 / var(--tw-text-opacity))
}</style>
