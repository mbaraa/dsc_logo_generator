<script lang="ts">
  import ChapterNameText from "./ChapterNameText.svelte";

  export let isColor: boolean;

  $: textColor = isColor ? "#676c72" : "white";
</script>

<div class="relative">
  <img
    class="rounded-[8px] border-[2px] border-[#F6F6F6]"
    alt="logo"
    src={`/logos/${isColor ? "h-color" : "h-white"}.png`}
  />
  <div
    class={`absolute w-full text-center text-[20px] left-[45%] translate-x-[-47.5%] top-[50%] z-[50] `}
  >
    <ChapterNameText on:chapter-name {textColor} />
  </div>
</div>
