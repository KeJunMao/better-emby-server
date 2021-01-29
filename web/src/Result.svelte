<script>
  export let result = {
    api,
    timer: 0,
    message: "loading",
  };
</script>

<tr class={result.message}>
  <td>{result.api}</td>
  <td>{result.timer} ms</td>
  <td>{result.message}</td>
</tr>

<style>
  tr{
    text-align: center;
  }
  .ok {
    background-color: #c0ffc0;
  }
  .error {
    background-color: #c07820;
  }
  .loading {
    background-color: aliceblue;
  }
</style>
