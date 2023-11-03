<script>import SimpleSplitButton from './simple-split-button.svelte';
export let Hst;
</script>

<Hst.Story>
  <Hst.Variant title="A Split Button">
    <SimpleSplitButton id="button-1" label="Edit">
      <div>Relative</div>
      <div>UTC</div>
      <div>Local</div>
    </SimpleSplitButton>
  </Hst.Variant>
</Hst.Story>
