export default class Text {
    text;
    color;
    size;
    width;
    fontFamily;
    topToLogo;

    constructor(text, color, fontFamily, size, topToLogo) {
        this.text = text;
        this.color = color;
        this.fontFamily = fontFamily;
        this.size = size;
        this.topToLogo = topToLogo;
    }

    changeOrientation(orientation) {
        this.resetTextSize(orientation);
        this.setTop(orientation);
    }

    resetTextSize(orientation) {
        this.size = orientation === "vertical"? "1.58em": "3.12em";
    }

    setTop(orientation) {
        this.topToLogo = orientation === "vertical"? "60.625%": "45%";
    }
}