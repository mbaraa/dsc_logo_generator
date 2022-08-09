<script lang="ts">
    export let isColor: boolean;
</script>

<img class="rounded-[8px] border-[2px] border-[#F6F6F6]" alt="logo" src={`/logos/${isColor? "v-color": "v-white"}.png`}/>