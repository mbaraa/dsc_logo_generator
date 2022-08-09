<script lang="ts">
  import {createEventDispatcher} from "svelte";

  const dispatch = createEventDispatcher();

  function downloadLogo() {
    dispatch("genlogo");
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
