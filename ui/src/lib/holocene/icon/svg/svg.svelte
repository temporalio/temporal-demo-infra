<script>export let width = 24;
export let height = 24;
export let viewBox = '0 0 24 24';
</script>

<svg
  {width}
  {height}
  {viewBox}
  fill="none"
  xmlns="http://www.w3.org/2000/svg"
  class={$$props.class}
>
  <slot />
</svg>
