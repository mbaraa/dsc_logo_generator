<script lang="ts">
  import config from "../config";

  export let isHorizontal: boolean;
  export let isColor: boolean;
  export let isTransparent: boolean;
  export let chapterName: string;

  $: color = isColor ? "color" : "white";
  $: opacity = isTransparent ? 0 : 1;
  $: _orientation = isHorizontal ? 2 : 1;

  async function downloadLogo() {
    const url = `${config.backendAddress}/api/genlogo/?uni_name=${chapterName}&img_color=${color}&opacity=${opacity}&logo_type=${_orientation}`;
    await fetch(url, {
      method: "GET",
      mode: "cors",
    })
      .then((resp) => resp.json())
      .then((data) => {
        let a = document.createElement("a");

        a.href = `data:image/png;base64,${data["image"]}`;
        a.download = `GDSC ${chapterName} ${
          _orientation === 2 ? "Horizontal" : "Vertical"
        } ${color}`;
        a.click();
      });
  }
</script>

<div class="flex justify-center mt-[20px]">
  <button
    class="rounded-[8px] bg-[#4385F3] p-[13px] font-bold text-white md:relative md:left-0 md-right-0 md:translate-y-0 md:top-0 absolute top-[76%] translate-y-[-50%]"
    on:click={downloadLogo}
  >
    <img
      class="inline w-[12px] h-[12px] mb-[3px]"
      alt="Download"
      src="/images/download.png"
    /> Download Logo
  </button>
</div>
