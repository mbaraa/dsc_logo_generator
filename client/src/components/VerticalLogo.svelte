<script lang="ts">
  import ChapterNameText from "./ChapterNameText.svelte";

  export let isColor: boolean;

  $: textColor = isColor ? "#676c72" : "white";
</script>

<div class="relative">
  <img
    class="rounded-[8px] border-[2px] border-[#F6F6F6]"
    alt="logo"
    src={`/logos/${isColor ? "v-color" : "v-white"}.png`}
  />
  <div
    class={`absolute w-full text-center text-[20px] left-[50%] translate-x-[-50%] top-[58%] z-[50]`}
  >
    <ChapterNameText on:chapter-name {textColor} />
  </div>
</div>
