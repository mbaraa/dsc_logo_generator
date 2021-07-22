<template>
    <div class="nav" :style="{
        backgroundColor: $store.getters.getTheme.top_bar_bg_color,
        color: $store.getters.getTheme.top_bar_font_color,
    }">
        <span>
            DSC Logo Generator
        </span>

        <FontAwesomeIcon class="themeSwitcher" icon="moon"
                         @click="changeTheme" title="click to toggle dark mode"/>
    </div>
</template>

<script>

import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {faMoon} from "@fortawesome/free-solid-svg-icons";
import {library} from "@fortawesome/fontawesome-svg-core";

// font awesome
library.add(faMoon)

export default {
    name: "TopBar",
    components: {
        FontAwesomeIcon
    },
    data() {
        return {
            isDarkMode: false,
            theme: this.$store.getters.getTheme
        }
    },
    methods: {
        changeTheme() {
            this.$store.dispatch("changeTheme")
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
    position: fixed;
    display: inline;
    text-align: right;
    right: 1.5%;
    cursor: pointer;
}

</style>