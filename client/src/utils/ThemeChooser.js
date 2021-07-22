import Theme from "../models/Theme";

export default class ThemeChooser {
    $theme;
    $light_theme;
    $dark_theme;

    constructor() {
        this.$initLightTheme();
        this.$initDarkTheme();

        this.$theme = this.$light_theme;
    }

    $initLightTheme() {
        this.$light_theme = new Theme("#2c3e50", "#FFFFFF", "#DDDDDD", "#2c3e50", "#FFFFFF", "#000000");
    }

    $initDarkTheme() {
        this.$dark_theme = new Theme("#FFFFFF", "#151D2B", "#2c3e50", "#FFFFFF", "#FFFFFF", "#2288F0");
    }

    getTheme() {
        return this.$theme;
    }

    getLightTheme() {
        return this.$light_theme;
    }

    getDarkTheme() {
        return this.$dark_theme;
    }

    changeTheme() {
        this.$theme = this.$theme === this.$light_theme ? this.$dark_theme : this.$light_theme;
    }
}