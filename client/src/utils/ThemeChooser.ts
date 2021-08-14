import Theme from "../models/Theme";

export default class ThemeChooser {
    $theme: Theme;
    $light_theme: Theme;
    $dark_theme: Theme;

    constructor() {
        this.$light_theme = new Theme("#2c3e50", "#FFFFFF", "#BCBCBC", "#2c3e50", "#FFFFFF", "#000000");
        this.$dark_theme = new Theme("#FFFFFF", "#2d333b", "#22272e", "#FFFFFF", "#FFFFFF", "#2288F0");

        this.$theme = this.$light_theme;
    }

    getTheme(): Theme {
        return this.$theme;
    }

    getLightTheme(): Theme {
        return this.$light_theme;
    }

    getDarkTheme(): Theme {
        return this.$dark_theme;
    }

    changeTheme(): void {
        this.$theme = this.$theme === this.$light_theme ? this.$dark_theme : this.$light_theme;
    }
}