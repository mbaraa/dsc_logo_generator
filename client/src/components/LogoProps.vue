<template>
    <div class="base" :style="{
        color: $store.getters.getTheme.font_color,
    }">
        <div class="prop">
            <label for="opacity">Background</label>&nbsp;
            <button id="opacity" @click="updateOpacity"> {{ makeTitle(bg) }}</button>
            |
            <label> Color Type</label>
            <FontAwesomeIcon :icon="{prefix: 'fas', iconName: 'palette'}"
                             :class="getColorChangerStyleClass()" class="colorChanger"
                             @click="changeColorType"
                             :title="getColorType() + ' : click to change the logo\'s color!'"/>

            |
            <label>Orientation</label>
            <img src="/orientation.png" alt="orientation" class="orientation"
                 @click="changeOrientation" title="click to change the logo's orientation!"/>
        </div>

        <input type="text" @keyup="setLogoText" v-model="logo.text.text" placeholder="University Name"
               class="uniName" @keyup.enter="generateAndDownloadLogo"/>
        &nbsp;
        <button class="genLogo" title="generate and download the current logo"
                @click="generateAndDownloadLogo">
            Download Logo
        </button>

        <!-- Logo goes brr -->
        <VerticalLogo v-if="logo.orientation.charAt(0) === 'v'"/>
        <HorizontalLogo v-if="logo.orientation.charAt(0) === 'h'"/>
    </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";
import VerticalLogo from "./VerticalLogo.vue";
import HorizontalLogo from "./HorizontalLogo.vue";

import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {faPalette} from "@fortawesome/free-solid-svg-icons";
import {library} from "@fortawesome/fontawesome-svg-core";
// font awesome
library.add(faPalette as any);

export default defineComponent({
    name: "LogoProps",
    components: {
        VerticalLogo,
        HorizontalLogo,
        FontAwesomeIcon
    },
    data() {
        return {
            logo: this.$store.getters.getLogo,
            bg: "opaque",
            colors: ["color", "white"],
            colorIndex: 0,
            fontSizeUpdateCounter: 0
        }
    },
    methods: {
        setLogo() {
            this.$store.dispatch("setLogo", this.logo);
        },
        setLogoText() {
            this.updateFontSize()
            this.setLogo();
        },
        updateOpacity() {
            this.bg = this.bg === "opaque" ? "transparent" : "opaque";
        },
        getOpacity() {
            this.logo.opacity = this.bg === "opaque" ? 1 : 0;
        },
        verifyLogoText(): boolean {
            return (this.logo.text !== "");
        },
        verifyLogoData(): boolean {
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
        getLogoOrientation(): number {
            return this.logo.orientation === "vertical" ? 1 : 2;
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
                    a.download = `GDSC ${this.logo.text.text} ${this.logo.orientation} ${this.logo.color}`;
                    a.click();
                })
        },
        updateFontSize() {
            let canvas: HTMLCanvasElement = document.createElement("canvas");
            let context: CanvasRenderingContext2D | null = canvas?.getContext("2d");
            if (context != null) {
                context.font = "ProductSans";

                let width = Number(context?.measureText(this.$store.getters.getLogo.text.text).width);
                let textstringLength = this.$store.getters.getLogo.width.length;

                if (width > Number(0.32 *
                    Number(this.$store.getters.getLogo.width.substring(0, textstringLength - 2)))
                ) {
                    this.fontSizeUpdateCounter = Number((this.fontSizeUpdateCounter + 1) % 7);
                    if (this.fontSizeUpdateCounter === 0) {
                        this.logo.text.size =
                            Number(
                                -0.001 + Number(this.logo.text.size.substring(0, textstringLength - 2))
                            ) + "em";
                    }
                } else {
                    this.logo.text.resetTextSize(this.logo.orientation);
                }

                this.setLogo();
            }
        },
        changeOrientation() {
            this.logo.changeOrientation();
        },
        getColorType(): string {
            return this.colors[this.colorIndex];
        },
        changeColorType() {
            this.colorIndex = (this.colorIndex + 1) % 2;
            this.logo.setColor(this.colors[this.colorIndex]);
        },
        getColorChangerStyleClass(): string {
            switch (this.colors[this.colorIndex]) {
                case "color":
                    return "colorChangerColored";
                case "white":
                default:
                    return "colorChangerWhite";
            }
        },
        makeTitle(str: string): string {
            return str.charAt(0).toUpperCase() + str.substring(1);
        }
    }
});
</script>

<style scoped>
.base {
    font-family: ProductSans;
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
    font-family: ProductSans;
    height: 40px;
    width: 330px;
    font-size: 1.2em;
    border-radius: 5px;
}

.genLogo {
    font-family: ProductSans;
    font-size: 1em;
    cursor: pointer;

    background-color: #4CAF50;
    border: none;
    color: #FFFFFF;
    padding: 12px 15px;
    border-radius: 5px
}

.prop {
    display: block;
    margin: 10px auto;
    font-size: 1.2em;
    width: 550px;
}

.orientation {
    width: 25px;
    height: 25px;
    cursor: pointer;
    padding: 2px;
    border-radius: 5px;
    vertical-align: middle;
}

.orientation:hover {
    background-color: #2C98CA;
}

.colorChanger {
    width: 25px;
    height: 25px;
    cursor: pointer;
    padding: 2px;
    border-radius: 5px;
    vertical-align: middle;
}

.colorChanger:hover {
    background-color: #2C98CA;
}

.colorChangerColored {
    color: #ff006f;
}

.colorChangerGray {
    color: #BFCBD9;
}

.colorChangerWhite {
    color: #000000;
}
</style>
