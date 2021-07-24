<template>
    <div class="text" :style="{
        top: $store.getters.getLogo.text.topToLogo,
    }">
        <p :style="{
            fontSize: $store.getters.getLogo.text.size, /* to match the values in the backend :) */
            color: $store.getters.getLogo.text.color,

        }" id="uniName">
            {{ $store.getters.getLogo.text.text }}
        </p>
    </div>
</template>

<script>
export default {
    name: "UniNameText",
}
</script>

<style scoped>
.text {
    font-family: ProductSans;
    position: relative;
    text-align: center;
    margin: auto;
    width: auto;
}
</style>