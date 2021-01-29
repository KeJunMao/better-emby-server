<script>
  export let disabled = false;
</script>

<button class="button" on:click {disabled}>
  <slot />
</button>

<style>
  .button {
    display: block;
    width: 100%;
    font-weight: 500;
    border-radius: 0.42em;
    background: rgba(85, 85, 85, 0.3);
    backdrop-filter: saturate(1.8) blur(1.5em);
    border: none;
    padding: 0.66em 1em;
    background: var(--input-background);
    color: var(--theme-text-color);
    margin: 0.25em 0;
  }
  .button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
</style>
