<template>
    <div title="the downloaded logo may vary than this one!" :style="{
    'background-image': 'url('+ String($store.getters.getLogo.image_path) +')',
    'box-shadow': $store.getters.getTheme.border_color+ ' 1px 1px 5px 2px',
    }" class="logo">
        <UniNameText style="font-size: .8em; top: 295px"/>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import UniNameText from "./UniNameText.vue";

export default defineComponent({
    name: "VerticalLogo",
    components: {
        UniNameText
    }
});
</script>

<style scoped>
.logo {
    width: 500px;
    height: 500px;
    position: relative;
    margin: 20px auto;
    border-radius: 5px;
}
</style>