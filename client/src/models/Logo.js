export default class Logo {
    $text;
    $opacity;
    $image_path;
    $color;
    $orientation;
    $width;

    constructor(color, text, orientation) {
        this.$color = color;
        this.$orientation = orientation;
        this.$text = text;

        // setters
        this.color = color;
        this.orientation = orientation;
        this.$updateImagePath(); // path according to orientation and color
        this.opacity = 0; // fully transparent
    }

    get image_path() {
        return this.$image_path;
    }

    get color() {
        return this.$color;
    }

    get orientation() {
        return this.$orientation;
    }

    get width() {
        return this.$width;
    }

    get opacity() {
        return this.$opacity;
    }

    set opacity(opacity) {
        this.$opacity = opacity;

        return this;
    }

    get text() {
        return this.$text;
    }

    set text(text) {
        this.$text.text = text;
    }

    /**
     * Sets logo's color, and sets a proper color for the text.
     * @param {string} color
     * */
    set color(color) {
        this.$color = color;
        this.$text.color = color !== "white" ? "#676c72" : "#FFFFFF";
        this.$updateImagePath();

        return this
    }

    /**
     * Sets the orientation of the logo, and updates the logo's image path.
     * @param {string} orientation "vertical" or "horizontal"
     * */
    set orientation(orientation) {
        if (orientation === "vertical" || orientation === "horizontal") {
            this.$orientation = orientation;
            this.$width = orientation === "vertical" ? "500px" : "1050px";
            this.$updateImagePath();
            // adapt text to the new orientation
            this.$text.changeOrientation(orientation);
        }

        return this
    }

    changeOrientation() {
        this.orientation = this.$orientation === "vertical"? "horizontal": "vertical";
    }

    $updateImagePath() {
        this.$image_path = `/${this.orientation.charAt(0)}-${this.color}.png`;
    }
}