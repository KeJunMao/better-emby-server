<script>
  import Result from "./Result.svelte";
  let text = "";
  let serverList = [];
  let resultList = [];
  let lock = false;
  let concurrent = true;

  function getServerList(text) {
    return (
      text
        .split("\n")
        .filter((v) => v.includes("http"))
        .map(
          (v) =>
            v.match(
              /(https?|http):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/
            )[0]
        ) || []
    );
  }
  async function getResult(api) {
    const result = {
      api,
      timer: 0,
      message: "error",
    };
    const startT = new Date();
    try {
      const resp = await fetch(`${api}/emby/system/info/public`);
      const data = await resp.json();
      const timer = new Date() - startT;
      result.timer = timer;
      result.message = "ok";
      return result;
    } catch {
      return result;
    }
  }
  async function normalSend() {
    lock = true;
    for (const key in resultList) {
      const data = await getResult(resultList[key].api);
      resultList[key] = data;
    }
    resultList = resultList.sort((a, b) => a.timer - b.timer);
    lock = false;
  }
  function concurrentSend() {
    lock = true;
    const tasks = resultList.map((v, k) => {
      return getResult(v.api).then((data) => {
        resultList[k] = data;
      });
    });
    Promise.all(tasks).then(() => {
      lock = false;
      resultList = resultList.sort((a, b) => a.timer - b.timer);
    });
  }
  function onTextChange() {
    serverList = getServerList(text);
    resultList = serverList.map((api) => ({
      api,
      timer: 0,
      message: "loading",
    }));
    if (concurrent) {
      concurrentSend();
    } else {
      normalSend();
    }
  }
</script>

<main>
  <textarea
    placeholder="粘贴厂妹返回的服务器地址消息"
    class="input"
    bind:value={text}
    on:change={onTextChange}
    disabled={lock}
  />
  <label>
    <input type="checkbox" bind:checked={concurrent} />
    并发请求 - 大幅提高请求速度，精确度有所降低
  </label>

  <p>有效服务器地址 {serverList.length} 条</p>
  {#if serverList.length > 0}
    <table width="100%">
      <thead>
        <th>接口</th>
        <th>耗时</th>
        <th>状态</th>
      </thead>
      <tbody>
        {#each resultList as result}
          <Result {result} />
        {/each}
      </tbody>
    </table>
  {/if}
</main>

<style>
  .input {
    display: block;
    width: 100%;
  }
</style>
