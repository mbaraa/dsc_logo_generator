<script lang="ts">
    import {createEventDispatcher} from "svelte";

    export let textColor: string;

    const dispatch = createEventDispatcher();

    let chapterName = "Type Your Chapter's Name Here!";

    function handleChangeChapterName(e: Event) {
        console.log(chapterName);
        dispatch("chapter-name", chapterName);
    }
</script>

<input autofocus={true} style={`color: ${textColor}; border-bottom: 1px solid ${textColor}`} class={`bg-[transparent] border-b-[1px] p-[2px] rounded-[8px] text-center`} bind:value={chapterName} on:keyup={handleChangeChapterName} />