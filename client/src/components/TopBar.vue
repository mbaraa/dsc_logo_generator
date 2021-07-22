<template>
    <div class="nav" :style="{
        backgroundColor: $store.getters.getTheme.top_bar_bg_color,
        color: $store.getters.getTheme.top_bar_font_color,
    }">
        <span>
            DSC Logo Generator
        </span>

        <DarkModeToggle class="themeSwitcher" title="click to toggle dark mode"/>
    </div>
</template>

<script>
import DarkModeToggle from "./DarkModeToggle";

export default {
    name: "TopBar",
    components: {
        DarkModeToggle,
    },
    data() {
        return {
            isDarkMode: false,
            theme: this.$store.getters.getTheme
        }
    },
}
</script>

<style scoped>
.nav {
    font-family: ProductSans;
    font-size: 2em;
    font-weight: bold;

    position: sticky;
    padding: 15px;

}

.themeSwitcher {
    position: absolute;
    display: inline;
    text-align: right;
    right: 1.5%;
}

</style>