<script>
  export let result = {
    api,
    timer: 0,
    message: "loading",
  };
</script>

<tr>
  <td>{result.api}</td>
  <td>{result.timer} ms</td>
  <td class={result.message}>{result.message}</td>
</tr>

<style>
  tr{
    text-align: center;
  }
  .ok {
    background-color: var(--theme-primary-color);
  }
  .error {
    background-color: #cb272a
  }
  .loading {
    background-color: #a6b64c;
  }
</style>
