<template>
    <div title="the downloaded logo may vary than this one!" :style="{
    'background-image': 'url('+ String($store.getters.getLogo.image_path) +')',
    'box-shadow': $store.getters.getTheme.border_color+ ' 1px 1px 5px 2px',
    color: $store.getters.getLogo.text_color,
    }" class="logo">
        <p class="text">{{ $store.getters.getLogo.text }}</p>
    </div>
</template>

<script>

export default {
    name: "Logo",
}
</script>

<style scoped>
.logo {
    width: 500px;
    height: 500px;
    position: relative;
    margin: 20px auto;
    border-radius: 5px;
}

.text {
    font-family: ProductSans;
    position: relative;
    text-align: center;
    margin: auto;

    font-size: 1.58em; /* to match the values */
    top: 60.625%; /* in the backend :) */
}
</style>