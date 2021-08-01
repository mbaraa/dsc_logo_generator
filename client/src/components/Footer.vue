<template>
    <footer class="base" :style="{
        color: $store.getters.getTheme.font_color,
        backgroundColor: $store.getters.getTheme.top_bar_bg_color,
    }">
        Made with
        <FontAwesomeIcon icon="laugh"/>
        by <b>GDSC-ASU (jo)</b> &nbsp;
        <a class="linkedin" href="https://www.linkedin.com/company/gdscasu/" target="_blank">
            <FontAwesomeIcon :icon="{prefix: 'fab', iconName: 'linkedin'}" class="linkedin"/>
        </a> &nbsp;
        <a class="facebook" href="https://web.facebook.com/GDSCASU" target="_blank">
            <FontAwesomeIcon :icon="{prefix: 'fab', iconName: 'facebook'}" class="facebook"/>
        </a> &nbsp;

        <a class="dsc" href="https://dsc.community.dev/applied-science-private-university/" target="_blank">
                <FontAwesomeIcon :icon="{prefix: 'fab', iconName: 'google'}"/>
        </a> &nbsp;
        •
        <a href="https://github.com/mbaraa/dsc_logo_generator" target="_blank">
            <FontAwesomeIcon icon="code"/>
            <FontAwesomeIcon :icon="{prefix: 'fab', iconName: 'github'}"/>
        </a> •
    </footer>
</template>

<script>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {faLaugh, faCode} from "@fortawesome/free-solid-svg-icons";
import {faFacebook, faGithub, faLinkedin, faGoogle} from "@fortawesome/free-brands-svg-icons";
import {library} from "@fortawesome/fontawesome-svg-core";
// font awesome
library.add(faFacebook, faGithub, faLinkedin, faLaugh, faCode, faGoogle);


export default {
    name: "Footer",
    components: {
        FontAwesomeIcon
    },
}
</script>

<style scoped>

a {
    color: #ff00ff;
    text-decoration: unset;
}

.linkedin {
    color: #0e76a8;
}

.facebook {
    color: #00acee;
}

.dsc {
    color: #FFFFFF;
}

.base {
    font-family: ProductSans;
    position: fixed;
    bottom: 0;
    margin: auto;
    width: 100%;
    text-align: center;
}
</style>