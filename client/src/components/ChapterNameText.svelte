<script lang="ts">
  import { createEventDispatcher } from "svelte";

  export let textColor: string;
  export let _class: string;

  const dispatch = createEventDispatcher();

  let chapterName = "";

  function handleChangeChapterName(e: Event) {
    dispatch("chapter-name", chapterName);
  }

  function handleEnter(e: KeyboardEvent) {
    if (e.key == "Enter") {
      dispatch("genlogoenter");
    }
  }
</script>

<!-- svelte-ignore a11y-autofocus -->
<input
  autofocus={true}
  style={`color: ${textColor}; border-bottom: 1px solid ${textColor}`}
  class={`bg-[transparent] border-b-[1px] p-[2px] rounded-[8px] ${_class}`}
  bind:value={chapterName}
  on:keyup={handleChangeChapterName}
  on:keypress={handleEnter}
  placeholder="Type Your Chapter's Name Here!"
/>
