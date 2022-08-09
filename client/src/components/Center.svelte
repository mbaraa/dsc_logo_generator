<script lang="ts">
  import Logo from "./Logo.svelte";
import LogoOptions from "./LogoOptions.svelte";

let isColor = true;
let isHorizontal = false;
let isTransparent = false;

  function handleOnChangeColor(e: CustomEvent) {
    isColor = !isColor
  }

  function handleOnSetHorizontal(e: CustomEvent) {
    isHorizontal = !isHorizontal
  }

  function handleOnSetTransparent(e: CustomEvent) {
    isTransparent = !isTransparent;
  }

</script>

<div class="my-[28px] mx-[53px]">
  <LogoOptions on:change-color={handleOnChangeColor} on:set-horizontal={handleOnSetHorizontal} on:set-transparent={handleOnSetTransparent}/>
  <Logo {isColor} {isHorizontal} />
</div>
