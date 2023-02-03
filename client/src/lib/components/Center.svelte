<script lang="ts">
    import config from "../config";

    import DownloadLogo from "./DownloadLogo.svelte";
    import Logo from "./Logo.svelte";
    import LogoOptions from "./LogoOptions.svelte";

    let isColor = true;
    let isHorizontal = false;
    let isTransparent = false;
    let chapterName = "";

    function handleOnChangeColor(e: CustomEvent) {
        isColor = e.detail as boolean;
    }

    function handleOnSetHorizontal(e: CustomEvent) {
        isHorizontal = e.detail as boolean;
    }

    function handleOnSetTransparent(e: CustomEvent) {
        isTransparent = e.detail as boolean;
    }

    function handleChangeChapterName(e: CustomEvent) {
        chapterName = e.detail as string;
    }

    $: color = isColor ? "color" : "white";
    $: opacity = isTransparent ? 0 : 1;
    $: _orientation = isHorizontal ? 2 : 1;

    async function downloadLogo() {
        const url = `${config.backendAddress}/api/genlogo/?uni_name=${chapterName}&img_color=${color}&opacity=${opacity}&logo_type=${_orientation}`;
        await fetch(url, {
            method: "GET",
            mode: "cors",
        })
            .then((resp) => resp.json())
            .then((data) => {
                let a = document.createElement("a");

                a.href = `data:image/png;base64,${data["image"]}`;
                a.download = `GDSC ${chapterName} ${
                    _orientation === 2 ? "Horizontal" : "Vertical"
                } ${color}.png`;
                a.click();
            });
    }
</script>

<div class="my-[28px] mx-[53px]">
    <LogoOptions
        on:change-color={handleOnChangeColor}
        on:set-horizontal={handleOnSetHorizontal}
        on:set-transparent={handleOnSetTransparent}
    />
    <Logo
        {isColor}
        {isHorizontal}
        on:chapter-name={handleChangeChapterName}
        on:genlogoenter={downloadLogo}
    />
    <div class="mb-[60px]">
        <DownloadLogo on:genlogo={downloadLogo} />
    </div>
</div>
