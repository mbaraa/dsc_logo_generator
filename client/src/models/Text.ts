export default class Text {
    text: string;
    color: string;
    size: string;
    width: number | string;
    fontFamily: string;
    topToLogo: string;

    constructor(text: string, color: string, fontFamily: string, size: string, topToLogo: string) {
        this.text = text;
        this.color = color;
        this.fontFamily = fontFamily;
        this.size = size;
        this.topToLogo = topToLogo;
        this.width = "1em";
    }

    changeOrientation(orientation: string): void {
        this.resetTextSize(orientation);
        this.setTop(orientation);
    }

    resetTextSize(orientation: string): void {
        this.size = orientation === "vertical" ? "1.58em" : "3.12em";
    }

    setTop(orientation: string): void {
        this.topToLogo = orientation === "vertical" ? "60.625%" : "45%";
    }
}