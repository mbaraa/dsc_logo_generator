<script lang="ts">
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  let transparent = false;

  $: text = "";

  function toggleTransparent() {
    dispatch("set-transparent", transparent);
    transparent = transparent;
  }
</script>

<select bind:value={text} on:change={toggleTransparent} placeholder="Trasparency" class="font-bold rounded-tr-[8px] rounded-br-[8px] py-[12px] px-[15px] text-black bg-white">
    <option value="" selected style="display:none">Transparency</option>
    <option>Opaque</option>
    <option
    title="Your will see the transparency when you download the logo!"
    >Transparent</option>
</select>

