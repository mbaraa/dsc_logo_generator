export default class Theme {
    font_color;
    bg_color;
    top_bar_bg_color;
    top_bar_font_color;
    footer_bg_color
    border_color;

    constructor(font_color, bg_color, top_bg_color, top_font_color, footer_bg_color, border_color) {
        this.font_color = font_color;
        this.bg_color = bg_color;
        this.top_bar_bg_color = top_bg_color;
        this.top_bar_font_color = top_font_color;
        this.footer_bg_color = footer_bg_color;
        this.border_color = border_color;
    }
}