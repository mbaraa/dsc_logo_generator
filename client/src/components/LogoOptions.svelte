<script lang="ts">
  import RoatateLogo from "./RoatateLogo.svelte";
  import SelectColor from "./SelectColor.svelte";
  import TransparentToggle from "./TransparentToggle.svelte";

  let colorOptions = [
    { i: 0, name: "Colored", action: "foo" },
    { i: 1, name: "White", action: "bar" },
  ];
</script>

<div class="hidden sm:flex justify-center">
  <div class="grid grid-cols-3 border-[1px] border-[#F6F6F6] rounded-[8px] mr-[5px]">
    <SelectColor options={colorOptions} on:change-color />
    <TransparentToggle on:set-transparent />
  </div>
  <RoatateLogo on:set-horizontal />
</div>
