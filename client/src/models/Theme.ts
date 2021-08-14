export default class Theme {
    font_color: string;
    bg_color: string;
    top_bar_bg_color: string;
    top_bar_font_color: string;
    footer_bg_color: string;
    border_color: string;

    constructor(font_color :string, bg_color: string, top_bg_color: string, top_font_color:string,
                footer_bg_color:string, border_color:string) {

        this.font_color = font_color;
        this.bg_color = bg_color;
        this.top_bar_bg_color = top_bg_color;
        this.top_bar_font_color = top_font_color;
        this.footer_bg_color = footer_bg_color;
        this.border_color = border_color;
    }
}