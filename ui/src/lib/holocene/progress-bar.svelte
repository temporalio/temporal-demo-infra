<script>export let width = '100%';
</script>

<div class="meter">
  <span style="width: {width}" />
</div>

<style>
  .meter {
    box-shadow: inset 0 -1px 1px rgb(255 255 255 / 30%);
    position: absolute;
    left: 0px;
    right: 0px;
    box-sizing: content-box;
    height: 5px;
    background-image: linear-gradient(to right, var(--tw-gradient-stops));
    --tw-gradient-from: #bfdbfe;
    --tw-gradient-to: rgb(191 219 254 / 0);
    --tw-gradient-stops: var(--tw-gradient-from), var(--tw-gradient-to);
    --tw-gradient-to: #2563eb;

    z-index: 5;
  }

  .meter > span {
    position: relative;
    display: block;
    height: 100%;
    overflow: hidden;
    --tw-bg-opacity: 1;
    background-color: rgb(59 130 246 / var(--tw-bg-opacity));

    background-image: linear-gradient(
      center bottom,
      rgb(96 165 250) 32%,
      rgb(59 130 246) 77%
    );
    box-shadow: inset 0 2px 9px rgb(255 255 255 / 25%),
      inset 0 -2px 6px rgb(0 0 0 / 40%);
}

  .meter > span::after {
    position: absolute;
    top: 0px;
    left: 0px;
    bottom: 0px;
    right: 0px;
    overflow: hidden;

    content: '';
    background-image: linear-gradient(
      -45deg,
      rgb(255 255 255 / 20%) 25%,
      transparent 25%,
      transparent 50%,
      rgb(255 255 255 / 20%) 50%,
      rgb(255 255 255 / 20%) 75%,
      transparent 75%,
      transparent
    );
    z-index: 1;
    background-size: 25px 25px;
    -webkit-animation: move 1.5s linear infinite;
            animation: move 1.5s linear infinite;
    border-top-right-radius: 8px;
    border-bottom-right-radius: 8px;
    border-top-left-radius: 20px;
    border-bottom-left-radius: 20px;
}

  @-webkit-keyframes move {
    0% {
      background-position: 0 0;
    }

    100% {
      background-position: 25px 25px;
    }
  }

  @keyframes move {
    0% {
      background-position: 0 0;
    }

    100% {
      background-position: 25px 25px;
    }
  }</style>
