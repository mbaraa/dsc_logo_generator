<template>
    <div :style="{
    backgroundColor: $store.getters.getTheme.bg_color,
}" style="width: 100%; height: 100vh">
        <TopBar/>
        <LogoProps/>
        <Footer/>
    </div>
</template>

<script>
import Footer from "./components/Footer.vue";
import TopBar from "./components/TopBar";
import LogoProps from "./components/LogoProps";

export default {
    name: 'App',
    components: {
        LogoProps,
        TopBar,
        Footer,
    }
}
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
