<script>import { dataConverterPort } from '../stores/data-converter-config';
import { codecEndpoint, passAccessToken, } from '../stores/data-encoder-config';
import { validateHttpOrHttps, validateHttps } from '../utilities/is-http';
import Modal from '$holocene/modal.svelte';
import CodecEndpointSettings from './codec-endpoint-settings.svelte';
import DataConverterPortSettings from './data-converter-port-settings.svelte';
export let showSettings;
export let onClose;
let endpoint = $codecEndpoint !== null && $codecEndpoint !== void 0 ? $codecEndpoint : '';
let port = $dataConverterPort !== null && $dataConverterPort !== void 0 ? $dataConverterPort : '';
let passToken = $passAccessToken !== null && $passAccessToken !== void 0 ? $passAccessToken : false;
$: error = '';
$: {
    if (passToken && !validateHttps(endpoint)) {
        error = 'Endpoint must be https:// if passing access token';
    }
    else if (endpoint && !validateHttpOrHttps(endpoint)) {
        error = 'Endpoint must start with http:// or https://';
    }
    else {
        error = '';
    }
}
const onCancel = () => {
    endpoint = $codecEndpoint !== null && $codecEndpoint !== void 0 ? $codecEndpoint : '';
    port = $dataConverterPort !== null && $dataConverterPort !== void 0 ? $dataConverterPort : '';
    passToken = $passAccessToken !== null && $passAccessToken !== void 0 ? $passAccessToken : false;
    onClose();
};
const onConfirm = () => {
    error = '';
    $codecEndpoint = endpoint;
    $passAccessToken = passToken;
    $dataConverterPort = port;
    onClose();
};
</script>

<Modal
  open={showSettings}
  on:cancelModal={onCancel}
  on:confirmModal={onConfirm}
  cancelText="Cancel"
  confirmDisabled={Boolean(error)}
>
  <h3 slot="title" data-cy="data-encoder-title">Data Encoder</h3>
  <div slot="content">
    <CodecEndpointSettings bind:endpoint bind:passToken {error} />
    <DataConverterPortSettings bind:port />
    <small data-cy="data-encoder-info"
      >If both are set, the Remote Codec Endpoint will be used.</small
    >
  </div>
</Modal>
