export default class Logo {
    text;
    $opacity;
    $image_path;
    $text_color;
    $color;
    $orientation;

    constructor(color, text, orientation) {
        this.$color = color;
        this.$orientation = orientation;
        this.text = text;

        // setters
        this.color = color;
        this.orientation = orientation;
        this.$updateImagePath(); // path according to orientation and color
        this.opacity = 0; // fully transparent
    }

    get image_path() {
        return this.$image_path;
    }

    get text_color() {
        return this.$text_color;
    }

    get color() {
        return this.$color;
    }

    get orientation() {
        return this.$orientation;
    }

    get opacity() {
        return this.$opacity;
    }

    set opacity(opacity) {
        this.$opacity = opacity;

        return this;
    }

    /**
     * Sets logo's color, and sets a proper color for the text.
     * @param {string} color
     * */
    set color(color) {
        this.$color = color;
        this.$text_color = color !== "white" ? "#676c72" : "#FFFFFF";
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
            this.$updateImagePath();
        }

        return this
    }

    $updateImagePath() {
        this.$image_path = `/${this.orientation.charAt(0)}-${this.color}.png`;
    }
}