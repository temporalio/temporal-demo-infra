<script>import Icon from '$holocene/icon/icon.svelte';
import { copyToClipboard } from '../utilities/copy-to-clipboard';
import { parseWithBigInt, stringifyWithBigInt, } from '../utilities/parse-with-big-int';
export let content;
export let inline = false;
export let language = 'json';
let root;
let isJSON = language === 'json';
const formatJSON = (jsonData) => {
    if (!jsonData)
        return;
    let parsedData;
    try {
        parsedData = parseWithBigInt(jsonData);
    }
    catch (error) {
        parsedData = jsonData;
    }
    return stringifyWithBigInt(parsedData, undefined, inline ? 0 : 2);
};
$: parsedContent = isJSON ? formatJSON(content) : content;
const { copy, copied } = copyToClipboard();
function highlight(root, language, source) {
    root.textContent = source;
    root.classList.forEach((item) => root.classList.remove(item));
    if (language) {
        root.classList.add(`language-${language}`);
    }
    window.Prism.highlightElement(root);
}
$: {
    if (root && window.Prism) {
        highlight(root, language, parsedContent);
    }
}
</script>

{#if parsedContent || parsedContent === null}
  <div
    class="w-full rounded-lg {inline
      ? 'h-auto overflow-auto'
      : 'h-full'} {$$props.class}"
    data-cy={$$props.dataCy}
  >
    <div class="relative h-full">
      <!-- The spacing for this if statement is like this because PRE's honor all whitespace and
      line breaks so we have this peculiar formatting to preserve this components output -->
      <pre
        class="w-full overflow-x-scroll rounded-lg p-4"
        class:h-full={!inline}><code
          bind:this={root}
          class="language-{language}"
          data-cy={$$props['data-cy']}
        /></pre>

      <button
        on:click={(e) => copy(e, parsedContent)}
        class="absolute top-2.5 right-2.5 rounded-md bg-gray-900 opacity-90 hover:bg-white"
      >
        <Icon
          name={$copied ? 'checkmark' : 'copy'}
          class="text-white hover:text-gray-900"
        />
      </button>
    </div>
  </div>
{/if}

<style>
  .inline {
    top: 1.25rem;
    right: 0.5rem
}</style>
