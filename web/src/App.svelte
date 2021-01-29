<script>
  import Result from "./Result.svelte";
  import Checkbox from "./Checkbox.svelte";
  import Header from "./Header.svelte";
  import Button from "./Button.svelte";

  let text = "";
  let serverList = [];
  let resultList = [];
  let estimate = false;
  let concurrent = true;
  $: serverList = getServerList(text);
  $: resultList = serverList.map((api) => ({
    api,
    timer: 0,
    message: "wait",
  }));
  let lock = false;

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
      await fetch(`${api}/emby/system/info/public`);
      const timer = new Date() - startT;
      result.timer = estimate ? (timer / 13.5625).toFixed(3) : timer;
      result.message = "ok";
      return result;
    } catch {
      return result;
    }
  }
  async function normalSend() {
    lock = true;
    for (const key in resultList) {
      resultList[key].message = "loading";
      const data = await getResult(resultList[key].api);
      resultList[key] = data;
    }
    resultList = resultList.sort((a, b) => a.timer - b.timer);
    lock = false;
  }
  function concurrentSend() {
    lock = true;
    const tasks = resultList.map((v, k) => {
      resultList[k].message = "loading";
      return getResult(v.api).then((data) => {
        resultList[k] = data;
      });
    });
    Promise.all(tasks).then(() => {
      lock = false;
      resultList = resultList.sort((a, b) => a.timer - b.timer);
    });
  }
  function onStart() {
    if (concurrent) {
      concurrentSend();
    } else {
      normalSend();
    }
  }
</script>

<Header
  value={{
    size: serverList.length,
    estimate,
    concurrent,
  }}
/>
<main>
  <div class="inputContainer">
    <textarea
      rows="5"
      placeholder="粘贴厂妹返回的服务器地址消息然后点击开始按钮"
      class="input"
      bind:value={text}
      disabled={lock}
    />
  </div>

  <Checkbox bind:checked={concurrent} disabled={lock}>启用并发</Checkbox>
  <Checkbox bind:checked={estimate} disabled={lock}>启用估算</Checkbox>
  <Button on:click={onStart} disabled={!serverList.length || lock}>开始</Button>

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

  <h3>注意事项</h3>
  <ul>
    <li>并发: 启用后大幅提升测试速度，但会损失精确度。</li>
    <li>
      估算: 默认是 http 请求延时，更接近使用环境；由于浏览器不能发送 ICMP
      ，使用估算时延迟更接近 ping。
    </li>
    <li>
      web 版本仅用于便携式访问，要获取更好精准度请使用 <a
        href="https://github.com/KeJunMao/better-emby-server">bes</a
      > 的 cli 版本。
    </li>
  </ul>
</main>

<style>
  main {
    padding: 2% 4%;
  }
  .input {
    display: block;
    width: 100%;
    background-color: var(--input-background);
    display: block;
    margin: 0;
    margin-bottom: 0 !important;
    padding: 0.4em 0.25em;
    box-sizing: border-box;
    width: 100%;
    border-radius: 0.3em;
    border: var(--line-size) solid var(--input-background);
    font-size: 110%;
    color: var(--theme-text-color);
  }
  .inputContainer {
    margin-bottom: 1.41em;
  }
  table {
    border-collapse: collapse;
    border-radius: 0.45em;
    background-color: var(--input-background);
    margin-top: 2em;
  }
</style>
