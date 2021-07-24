<template>
    <div class="base" :style="{
        color: $store.getters.getTheme.font_color,
    }">
        <div class="prop">
            <label for="opacity">Transparent Background</label>
            <input id="opacity" type="checkbox" value=false v-model="opacity">
            |
            <label for="colors"> Logo Color Type</label>&nbsp;
            <select id="colors" name="colors" @change="setLogoColor" v-model="logo.color">
                <option value="color">Colored</option>
                <option value="gray">Gray</option>
                <option value="white">White</option>
            </select>
            |
            <label>Orientation</label>
            <button @click="changeOrientation">Change</button>
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
            opacity: false,
        }
    },mounted() {
        console.log(this.logo)
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
            this.logo.opacity = this.opacity ? 0 : 1;
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
        getLogoOrientation() {
            return this.logo.orientation === "vertical"? 1:2;
        },
        async getLogoFromServer() {
            const url = `http://127.0.0.1:1105/api/genlogo/?uni_name=${this.logo.text.text}&img_color=${this.logo.color}&opacity=${this.logo.opacity}&logo_type=${this.getLogoOrientation()}`;
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
                this.logo.text.resetTextSize();
            }

            this.setLogo();
        },
        changeOrientation() {
            this.logo.changeOrientation();
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
    width: 320px;
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
    font-size: 1.1em;
    width: 500px;
}
</style>