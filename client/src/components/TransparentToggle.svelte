<script lang="ts">
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  let transparent = false;

  $: text = transparent ? "Transparent" : "Opaque";

  function toggleTransparent() {
    dispatch("set-transparent", !transparent);
    transparent = !transparent;
  }
</script>

<button
    title={transparent? "Your will see the transparency when you download the logo!": ""}
  on:click={toggleTransparent}
  class={`rounded-[8px] py-[12px] px-[20px] text-[${
    transparent ? "white" : "#1F2436"
  }] bg-[${transparent ? "#4385F3" : "white"}]`}>{text}</button
>
