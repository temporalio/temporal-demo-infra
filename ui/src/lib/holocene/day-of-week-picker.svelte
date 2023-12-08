<script>import { genericWeekDays, weekDays } from '../utilities/calendar';
export let daysOfWeek;
const onClick = (e, day) => {
    if (e.metaKey) {
        // For preventing mixing Generic and Sunday-Friday
        if (genericWeekDays.find((g) => g.value === day)) {
            daysOfWeek = [day];
        }
        else if (daysOfWeek.find((d) => genericWeekDays.find((g) => g.value === d))) {
            daysOfWeek = [day];
        }
        // For Sunday-Friday
        if (daysOfWeek.includes(day)) {
            daysOfWeek.filter((d) => d !== day);
        }
        else {
            daysOfWeek = [...daysOfWeek, day];
        }
    }
    else {
        daysOfWeek = [day];
    }
};
</script>

<div class="flex flex-col gap-4 text-center">
  <div class="flex gap-2 text-center">
    {#each genericWeekDays as { label, value }}
      <button
        class="cell"
        class:active={daysOfWeek.includes(value)}
        on:click|preventDefault={(e) => onClick(e, value)}>{label}</button
      >
    {/each}
  </div>
  <div class="flex flex-wrap gap-2 text-center">
    {#each weekDays as { label, value }}
      <button
        class="cell"
        class:active={daysOfWeek.includes(value)}
        on:click|preventDefault={(e) => onClick(e, value)}>{label}</button
      >
    {/each}
  </div>
</div>

<style>
  .cell {

    width: auto;

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
