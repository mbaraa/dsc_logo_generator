<script lang="ts">
  import ChapterNameText from "./ChapterNameText.svelte";

  export let isColor: boolean;

  $: textColor = isColor ? "#676c72" : "white";
</script>

<div class="relative">
  <img
    class="rounded-[8px] border-[2px] border-[#F6F6F6] md:w-[800px]"
    alt="logo"
    src={`/logos/${isColor ? "h-color" : "h-white"}.png`}
  />
  <div
    class={`text-[20px] absolute w-full text-center left-[-2%] md:left-[-10.5%] md:text-[25px] top-[55%] z-[50] `}
  >
    <ChapterNameText on:chapter-name {textColor} />
  </div>
</div>
