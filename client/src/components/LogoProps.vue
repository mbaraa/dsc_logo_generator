<template>
    <div class="base" :style="{
        color: $store.getters.getTheme.font_color,
    }">
        <label for="opacity">Transparent Background</label>
        <input class="opacity" id="opacity" type="checkbox" value=false
               v-model="opacity">

        |<label for="colors"> Logo Color Type</label>&nbsp;
        <select id="colors" name="colors" @change="setLogoColor" v-model="logo.color">
            <option value="color">Colored</option>
            <option value="gray">Gray</option>
            <option value="white">White</option>
        </select>

        <input type="text" @keyup="setLogoText" v-model="logo.text" placeholder="University Name"
               class="uniName prop"/>

        <button class="genLogo" title="generate and download the current logo" @click="generateAndDownloadLogo">Download Logo
        </button>
        <!--        <button class="openHorizontal" id="openHorizontal" onclick="window.location.href='horizontal_index.html'">Switch To Horizontal</button>-->
        <!-- Logo goes brr -->
        <Logo/>

    </div>
</template>

<script>
import Logo from "./Logo";
export default {
    name: "LogoProps",
    components: {Logo},
    data() {
        return {
            logo: this.$store.getters.getLogo,
            opacity: false,
        }
    },
    methods: {
        setLogo() {
            this.$store.dispatch("setLogo", this.logo);
        },
        setLogoColor() {
            this.setLogo();
        },
        setLogoText() {
            this.setLogo();
        },
        updateOpacity() {
            this.logo.opacity = this.opacity? 0: 1;
        },
        verifyLogoText() {
            return (this.logo.text !== "");
        },
        verifyLogoTextLength() {
            return (this.logo.text.length <= 66);
        },
        verifyLogoData() {
            if (!this.verifyLogoText()) {
                window.alert("Hmm... a nameless university!");
                return false;
            }
            if (!this.verifyLogoTextLength()) {
                window.alert("The generated logo may differ from this one!");
            }
            return true;
        },
        generateAndDownloadLogo() {
            if (this.verifyLogoData()) {
                this.updateOpacity();
                this.setLogo();

                this.getLogoFromServer();
            }
        },
        async getLogoFromServer() {
            const url = `http://127.0.0.1:1105/logo-gen/api/gen?uni_name=${this.logo.text}&img_color=${this.logo.color}&opacity=${this.logo.opacity}&logo_type=1`;
            await fetch(url)
                .then(resp => resp.json())
                .then(data => {
                    let a = document.createElement("a");

                    a.href = `data:image/png;base64,${data["image"]}`;
                    a.download = `DSC ${this.logo.text} ${this.logo.orientation}`;
                    a.click();
                })
        },
    }
}
</script>

<style scoped>
.base {
    position: relative;
    text-align: center;
    margin: auto;
    width: auto;
    height: auto;
    overflow-x: hidden;
    overflow-y: auto;
    padding-top: 20px;

    font-family: ProductSans;
    font-size: 1.1em;
}

.uniName {
    height: 30px;
    width: 500px;
    font-size: 1.2em;

}

.genLogo {
    font-size: 1.15em;
}

.prop {
    display: block;
    margin: 10px auto;
}
</style>