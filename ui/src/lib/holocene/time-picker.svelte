<script>import ToggleButtons from './toggle-button/toggle-buttons.svelte';
import ToggleButton from './toggle-button/toggle-button.svelte';
import Input from './input/input.svelte';
export let hour = '';
export let minute = '';
export let second = '';
export let half = 'AM';
</script>

<div class="flex gap-2">
  <Input
    id="hour"
    bind:value={hour}
    placeholder="00"
    suffix="hrs"
    maxLength={2}
    error={Boolean(parseInt(hour) > 12)}
  />
  <Input
    id="minute"
    required
    bind:value={minute}
    placeholder="00"
    suffix="min"
    maxLength={2}
    error={Boolean(parseInt(hour) > 59)}
  />
  <Input
    id="second"
    bind:value={second}
    placeholder="00"
    suffix="sec"
    maxLength={2}
    error={Boolean(parseInt(hour) > 59)}
  />
  <ToggleButtons>
    <ToggleButton active={half === 'AM'} on:click={() => (half = 'AM')}
      >AM</ToggleButton
    >
    <ToggleButton active={half === 'PM'} on:click={() => (half = 'PM')}
      >PM</ToggleButton
    >
  </ToggleButtons>
</div>
