<script lang="ts">
    import {createEventDispatcher} from "svelte";

    const dispatch = createEventDispatcher();

    let orientation = false;

    function rotateLogo() {
        dispatch("set-horizontal", !orientation)
        orientation = !orientation
    }
</script>

<button on:click={rotateLogo} class="py-[10px] px-[14px] border-[1px] border-[#F6F6F6] rounded-[8px]">
    <img class="inline w-[20px] h-[20px]" alt="Rotate Logo" src="/images/rotate.png"/>
</button>