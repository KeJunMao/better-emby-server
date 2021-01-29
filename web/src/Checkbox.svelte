<script>
  export let checked = false
</script>

<label class="checkboxContainer emby-checkbox-label">
  <input type="checkbox" class="checkbox" bind:checked={checked} />
  <span class="checkboxLabel">
    <slot/>
  </span>
</label>

<style>
  .checkboxContainer {
    display: flex;
    align-items: center;
    cursor: pointer;
    margin-bottom: 1.41em;
  }
  .checkbox {
    display: none;
  }
  .checkbox:checked + span:before {
    border-color: var(--theme-primary-color);
    background-color: var(--theme-primary-color);
  }
  .checkboxLabel {
    position: relative;
    margin: 0;
    display: inline-flex;
    align-items: flex-start;
    font-size: .9em;
  }
  .checkboxLabel::before {
    content: "";
    display: inline-block;
    height: 1.56em;
    width: 1.56em;
    border: 0.14em solid;
    border-radius: 0.3em;
    position: static;
    margin-right: 0.5em;
    margin-top: -0.18em;
    flex-shrink: 0;
  }
  .checkbox:checked + .checkboxLabel::after {
    display: block;
  }
  .checkbox + .checkboxLabel::after {
    display: none;
  }
  .checkboxLabel::after {
    position: absolute;
    content: "";
    display: inline-block;
    height: 0.36em;
    width: 1em;
    border-left: 0.18em solid;
    border-bottom: 0.18em solid;
    transform: rotate(-45deg);
    top: 0.32em;
    left: 0.33em;
    color: #fff;
  }
</style>
