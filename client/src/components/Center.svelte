<script lang="ts">
  import ChapterNameText from "./ChapterNameText.svelte";
  import DownloadLogo from "./DownloadLogo.svelte";
  import Logo from "./Logo.svelte";
  import LogoOptions from "./LogoOptions.svelte";

  let isColor = true;
  let isHorizontal = false;
  let isTransparent = false;
  let chapterName = "Type Your Chapter Name Here!";

  function handleOnChangeColor(e: CustomEvent) {
    isColor = e.detail as boolean;
  }

  function handleOnSetHorizontal(e: CustomEvent) {
    isHorizontal = e.detail as boolean;
  }

  function handleOnSetTransparent(e: CustomEvent) {
    isTransparent = e.detail as boolean;
  }

  function handleChangeChapterName(e: CustomEvent) {
    chapterName = e.detail as string;
  }
</script>

<div class="my-[28px] mx-[53px]">
  <LogoOptions
    on:change-color={handleOnChangeColor}
    on:set-horizontal={handleOnSetHorizontal}
    on:set-transparent={handleOnSetTransparent}
  />
  <Logo {isColor} {isHorizontal} on:chapter-name={handleChangeChapterName} />
  <DownloadLogo {chapterName} {isColor} {isHorizontal} {isTransparent} />
</div>
