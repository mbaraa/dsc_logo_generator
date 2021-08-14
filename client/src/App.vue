<template>
    <div :style="{
    backgroundColor: $store.getters.getTheme.bg_color,
}" style="width: 100%; height: 100vh">
        <TopBar/>
        <LogoProps/>
        <br/>
        <Footer/>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import Footer from "./components/Footer.vue";
import TopBar from "./components/TopBar.vue";
import LogoProps from "./components/LogoProps.vue";

export default defineComponent({
    name: 'App',
    components: {
        LogoProps,
        TopBar,
        Footer,
    }
});
</script>

<style>

@font-face {
    font-family: ProductSans;
    src: url("/ProductSans-Regular.ttf")
}

#app {
    height: 100%;
    font-family: ProductSans;
}

</style>
