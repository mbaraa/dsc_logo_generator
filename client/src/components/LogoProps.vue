<template>
    <div class="base" :style="{
        color: $store.getters.getTheme.font_color,
    }">
        <div class="prop">
            <label for="opacity">Background</label>&nbsp;
            <button id="opacity" @click="updateOpacity"> {{ makeTitle(bg) }} </button>
            |
            <label for="colors"> Color Type</label>&nbsp;
            <select id="colors" name="colors" @change="setLogoColor" v-model="logo.color">
                <option value="color">Colored</option>
                <option value="gray">Gray</option>
                <option value="white">White</option>
            </select>
            |
            <label>Orientation</label>&nbsp;
            <button @click="changeOrientation">{{ makeTitle(logo.orientation) }}</button>
        </div>

        <input type="text" @keyup="setLogoText" v-model="logo.text.text" placeholder="University Name"
               class="uniName" @keyup.enter="generateAndDownloadLogo"/>
        &nbsp;
        <button class="genLogo" title="generate and download the current logo"
                @click="generateAndDownloadLogo" :style="{
                    backgroundColor: $store.getters.getTheme.top_bar_bg_color,
                }" style="box-shadow: rgba(0, 0, 0, 0) 0 0; border-radius: 5px">
            Download Logo
        </button>

        <!--        <button class="openHorizontal" id="openHorizontal" onclick="window.location.href='horizontal_index.html'">Switch To Horizontal</button>-->
        <!-- Logo goes brr -->
        <VerticalLogo v-if="logo.orientation.charAt(0) === 'v'"/>
        <HorizontalLogo v-if="logo.orientation.charAt(0) === 'h'"/>
    </div>
</template>

<script>
import VerticalLogo from "./VerticalLogo.vue";
import HorizontalLogo from "./HorizontalLogo.vue";

export default {
    name: "LogoProps",
    components: {
        VerticalLogo,
        HorizontalLogo
    },
    data() {
        return {
            logo: this.$store.getters.getLogo,
            bg: "opaque",
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
            this.updateFontSize()
            this.setLogo();
        },
        updateOpacity() {
            this.bg = this.bg === "opaque"? "transparent": "opaque";
        },
        getOpacity() {
            this.logo.opacity = this.bg === "opaque" ? 1: 0;
        },
        verifyLogoText() {
            return (this.logo.text !== "");
        },
        verifyLogoData() {
            if (!this.verifyLogoText()) {
                window.alert("Hmm... a nameless university!");
                return false;
            }
            return true;
        },
        generateAndDownloadLogo() {
            if (this.verifyLogoData()) {
                this.getOpacity();
                this.setLogo();

                this.getLogoFromServer();
            }
        },
        getLogoOrientation() {
            return this.logo.orientation === "vertical"? 1:2;
        },
        async getLogoFromServer() {
            const url = `/api/genlogo/?uni_name=${this.logo.text.text}&img_color=${this.logo.color}&opacity=${this.logo.opacity}&logo_type=${this.getLogoOrientation()}`;
            await fetch(url, {
                method: "GET",
                mode: "cors",
            })
                .then(resp => resp.json())
                .then(data => {
                    let a = document.createElement("a");

                    a.href = `data:image/png;base64,${data["image"]}`;
                    a.download = `DSC ${this.logo.text.text} ${this.logo.orientation} ${this.logo.color}`;
                    a.click();
                })
        },
        updateFontSize() {
            let canvas = document.createElement("canvas");
            let context = canvas.getContext("2d");
            context.font = "ProductSans";

            let width = Number(context.measureText(this.$store.getters.getLogo.text.text).width);
            let textStringLength = this.$store.getters.getLogo.width.length;

            if (width > Number(0.32 *
                Number(this.$store.getters.getLogo.width.substring(0, textStringLength - 2)))
            ) {
                this.logo.text.size =
                    Number(
                        -0.04 + Number(this.logo.text.size.substring(0, textStringLength - 2))
                    ) + "em";
            } else {
                this.logo.text.resetTextSize(this.logo.orientation);
            }

            this.setLogo();
        },
        changeOrientation() {
            this.logo.changeOrientation();
        },
        makeTitle(str) {
            return str.charAt(0).toUpperCase() + str.substring(1);
        }
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
}

.uniName {
    height: 40px;
    width: 330px;
    font-size: 1.2em;
    border-radius: 5px;
}

.genLogo {
    font-size: 1.15em;
    height: 44px;
    cursor: pointer;
}

.prop {
    display: block;
    margin: 10px auto;
    font-size: 1em;
    width: 550px;
}
</style>