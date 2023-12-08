<script>var _a;
import { onMount } from 'svelte';
import { page } from '$app/stores';
import { MAX_PAGE_SIZE, pagination, perPageFromSearchParameter, options, defaultItemsPerPage, } from '../stores/pagination';
import { updateQueryParameters } from '../utilities/update-query-parameters';
import FilterSelect from './select/filter-select.svelte';
import { getFloatStyle } from '../utilities/get-float-style';
import Icon from '$holocene/icon/icon.svelte';
export let items;
export let floatId = undefined;
export let startingIndex = 0;
export let perPageKey = 'per-page';
export let currentPageKey = 'page';
export let itemsPerPage = null;
$: perPage =
    itemsPerPage !== null
        ? String(itemsPerPage)
        : String(perPageFromSearchParameter($page.url.searchParams.get(perPageKey))).toString();
$: {
    if (parseInt(perPage, 10) > parseInt(MAX_PAGE_SIZE, 10)) {
        updateQueryParameters({
            parameter: perPageKey,
            value: MAX_PAGE_SIZE,
            url: $page.url,
        });
    }
    else if (!options.includes(perPage)) {
        updateQueryParameters({
            parameter: perPageKey,
            value: defaultItemsPerPage,
            url: $page.url,
        });
    }
}
$: store = pagination(items, perPage, startingIndex);
$: currentPage =
    (_a = $page.url.searchParams.get(currentPageKey)) !== null && _a !== void 0 ? _a : $store.currentPage;
$: store.jumpToPage(currentPage);
let screenWidth;
let width;
let height;
onMount(() => {
    updateWidth();
    if (startingIndex > 0) {
        handlePageChange();
    }
});
const handlePageChange = () => {
    updateQueryParameters({
        parameter: currentPageKey,
        value: $store.currentPage,
        url: $page.url,
    });
};
const updateWidth = () => {
    var _a;
    if (floatId) {
        width = (_a = document.getElementById(floatId)) === null || _a === void 0 ? void 0 : _a.clientWidth;
    }
};
$: floatStyle = getFloatStyle({ width, height, screenWidth });
</script>

<svelte:window bind:innerWidth={screenWidth} on:resize={updateWidth} />

<div class="pagination relative mb-8 flex flex-col gap-4">
  <div
    class={`flex flex-col items-center gap-4 lg:flex-row ${
      $$slots['action-top-left'] ? 'justify-between' : 'justify-end'
    }`}
  >
    <slot name="action-top-left" />
    <nav
      style={floatStyle}
      bind:clientHeight={height}
      class="flex flex-col justify-end gap-4 md:flex-row"
    >
      <slot name="action-top-center" />
      {#if !itemsPerPage}
        <FilterSelect
          label="Per Page"
          parameter={perPageKey}
          value={perPage}
          {options}
        />
      {/if}
      <div class="flex items-center justify-center gap-1">
        <button
          class="caret"
          disabled={!$store.hasPrevious}
          on:click={() => {
            store.previous();
            handlePageChange();
          }}
        >
          <Icon name="chevron-left" />
        </button>
        <p>
          {$store.length ? $store.startingIndex + 1 : 0}–{$store.endingIndex +
            1} of {$store.length}
        </p>
        <button
          class="caret"
          disabled={!$store.hasNext}
          on:click={() => {
            store.next();
            handlePageChange();
          }}
        >
          <Icon name="chevron-right" />
        </button>
      </div>
      <slot name="action-top-right" />
    </nav>
  </div>
  <slot visibleItems={$store.items} initialItem={$store.initialItem} />
  <nav
    class={`flex ${
      $$slots['action-bottom-left'] ? 'justify-between' : 'justify-end'
    }`}
  >
    <slot name="action-bottom-left" />
    <div class="flex gap-4">
      {#if !itemsPerPage}
        <FilterSelect
          label="Per Page"
          parameter={perPageKey}
          value={perPage}
          {options}
        />
      {/if}
      <div class="flex items-center justify-center gap-1">
        <button
          class="caret"
          disabled={!$store.hasPrevious}
          on:click={() => store.previous()}
        >
          <Icon name="chevron-left" />
        </button>
        <p>
          {$store.length ? $store.startingIndex + 1 : 0}–{$store.endingIndex +
            1} of {$store.length}
        </p>
        <button
          class="caret"
          disabled={!$store.hasNext}
          on:click={() => store.next()}
        >
          <Icon name="chevron-right" />
        </button>
      </div>
      <slot name="action-bottom-right" />
    </div>
  </nav>
</div>

<style>
  .caret {

    --tw-text-opacity: 1;

    color: rgb(113 113 122 / var(--tw-text-opacity))
}

  .caret:disabled {

    cursor: not-allowed;

    --tw-text-opacity: 1;

    color: rgb(212 212 216 / var(--tw-text-opacity))
}</style>
