<script>import Button from '$holocene/button.svelte';
export let showOn = 150; // pixels
export let scrollToContainer = false;
let hidden = true;
function getScrollContainer() {
    return scrollToContainer
        ? document.getElementById('scroll-container')
        : document.documentElement || document.body;
}
function scrollToTop() {
    var _a;
    (_a = getScrollContainer()) === null || _a === void 0 ? void 0 : _a.scrollIntoView();
}
function handleOnScroll() {
    const scrollContainer = getScrollContainer();
    if (!scrollContainer)
        return;
    hidden = Boolean(scrollContainer.getBoundingClientRect().top + showOn > 0);
}
</script>

<svelte:window on:scroll={handleOnScroll} />

<div id="scroll-container" class={$$props.class}>
  <slot />
  <div class="back-to-top" class:hidden>
    <Button
      class="!py-1.5 !px-1"
      icon="arrow-up"
      variant="secondary"
      on:click={scrollToTop}
    />
  </div>
</div>

<style>
  .back-to-top {

    position: fixed;

    right: 1.25rem;

    bottom: 1.25rem;

    z-index: 50
}

  .back-to-top.hidden {

    visibility: hidden
}</style>
