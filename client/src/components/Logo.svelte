<script lang="ts">
  import HorizontalLogo from "./HorizontalLogo.svelte";
  import VerticalLogo from "./VerticalLogo.svelte";

  export let isHorizontal: boolean;
  export let isColor: boolean;
</script>

<div
  class="flex justify-center md:mt-[35px] md:relative md:left-0 md-right-0 md:translate-y-0 md:top-0 absolute top-[45%] translate-y-[-50%] left-[10px] right-[10px]"
>
  {#if isHorizontal}
    <HorizontalLogo on:chapter-name {isColor} />
  {:else}
    <VerticalLogo on:chapter-name {isColor} />
  {/if}
</div>
