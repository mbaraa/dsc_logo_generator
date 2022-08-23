<script lang="ts">
  import HorizontalLogo from "./HorizontalLogo.svelte";
  import VerticalLogo from "./VerticalLogo.svelte";

  export let isHorizontal: boolean;
  export let isColor: boolean;
</script>

<div
  title="This logo is just a preview, download the logo to see the correct one!"
  class="flex justify-center md:mt-[35px] md:relative md:left-0 md-right-0 md:translate-y-0 md:top-0 absolute top-[45%] translate-y-[-50%] left-[10px] right-[10px]"
>
  {#if isHorizontal}
    <HorizontalLogo on:chapter-name on:genlogoenter {isColor} />
  {:else}
    <VerticalLogo on:chapter-name on:genlogoenter {isColor} />
  {/if}
</div>
