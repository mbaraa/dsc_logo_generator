<script lang="ts">
  import { createEventDispatcher } from "svelte";
  export let options: { i: number; name: string; action: any }[];
  let currentOption = 0;

  const dispatch = createEventDispatcher();

  function changeColor() {
    dispatch("change-color", options[0].name !== options[currentOption].name);
    currentOption = (currentOption + 1) % options.length;
  }
</script>

{#each options as option}
  <button
    on:click={changeColor}
    style={`color: ${currentOption === option.i ? "white" : "#1F2436"}`}
    class={`font-bold first:rounded-tl-[8px] first:rounded-bl-[8px] option py-[12px] px-[20px] bg-[${currentOption === option.i ? "#4385F3" : "white"}]`}
    >{option.name}</button
  >
{/each}

<style>
  .option:nth-child(2) {
    border-radius: 0 8px 8px 0;
  }
</style>
