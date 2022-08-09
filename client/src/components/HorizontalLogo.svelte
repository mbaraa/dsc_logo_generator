<script lang="ts">
  import ChapterNameText from "./ChapterNameText.svelte";

  export let isColor: boolean;

  $: textColor = isColor ? "#676c72" : "white";
</script>

<div class="relative w-auto">
  <img
    class="rounded-[8px] border-[2px] border-[#F6F6F6] md:w-[800px]"
    alt="logo"
    src={`/logos/${isColor ? "h-color" : "h-white"}.png`}
  />
  <div
    class={`text-[20px] md:text-[25px] z-[50] relative bottom-[32%] left-[13%] sm:left-[14.1%]`}
  >
    <ChapterNameText _class="text-left pl-[20px]" on:chapter-name {textColor} />
  </div>
</div>
