<template>
    <div title="the downloaded logo may vary than this one!" :style="{
    'background-image': 'url('+ String($store.getters.getLogo.image_path) +')',
    'box-shadow': $store.getters.getTheme.border_color+ ' 1px 1px 5px 2px',
    color: $store.getters.getLogo.text_color,
    }" class="logo">
        <UniNameText style="left: 214px; position: absolute"/>
    </div>
</template>

<script>
import UniNameText from "./UniNameText.vue";

export default {
    name: "HorizontalLogo",
    components: {
        UniNameText
    }
}
</script>

<style scoped>
.logo {
    width: 1050px;
    height: 210px;
    position: relative;
    margin: 20px auto;
    border-radius: 5px;
}
</style>