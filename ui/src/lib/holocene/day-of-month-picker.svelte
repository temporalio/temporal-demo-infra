<script>const daysInMonth = Array.from({ length: 31 }).map((_, i) => i + 1);
export let daysOfMonth;
const onClick = (e, day) => {
    if (e.metaKey) {
        if (daysOfMonth.includes(day)) {
            daysOfMonth.filter((d) => d !== day);
        }
        else {
            daysOfMonth = [...daysOfMonth, day];
        }
    }
    else {
        daysOfMonth = [day];
    }
};
</script>

<div class="text-center">
  <div
    class="grid grid-cols-7 gap-y-4 gap-x-4 rounded border border border-gray-900 px-0 py-4 text-center md:gap-x-8 md:px-2 xl:gap-x-16 xl:px-4"
  >
    {#each daysInMonth as day}
      <button
        class="cell"
        class:active={daysOfMonth.includes(day)}
        on:click|preventDefault={(e) => onClick(e, day)}>{day}</button
      >
    {/each}
  </div>
</div>

<style>
  .cell {

    height: 2.5rem;

    width: 100%;

    cursor: pointer;

    border-radius: 0.25rem;

    padding: 0.25rem
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
