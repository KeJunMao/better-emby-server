<script>
  export let value = {
    size: 0,
    estimate: false,
    concurrent: false,
  };
</script>

<header>
  <img class="logo" src="./logowhite.png" alt="" /> BES - 有效服务器地址 {value.size}
  条 - {#if !value.concurrent}
    非{/if}并发{#if value.estimate}
    估算{/if}模式
</header>

<style>
  .logo {
    height: 1.8em;
    margin-right: 0.5em;
  }
  header {
    padding: 0.8em 0;
    display: flex;
    align-items: center;
    flex-grow: 1;
    font-weight: bolder;
  }
</style>
