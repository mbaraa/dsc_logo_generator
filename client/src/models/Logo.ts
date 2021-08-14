import Text from "@/models/Text";

export default class Logo {
    $text: Text;
    $opacity: number;
    $image_path: string | undefined;
    $color: string;
    $orientation: string;
    $width: string | undefined;

    constructor(color: string, text: Text, orientation: string) {
        this.$color = color;
        this.$orientation = orientation;
        this.$text = text;

        // setters
        this.$color = color;
        this.$orientation = orientation;
        this.$updateImagePath(); // path according to orientation and color
        this.$opacity = 0; // fully transparent
        this.$width = "500px";
    }

    get image_path(): string {
        return <string>this.$image_path;
    }

    get color(): string {
        return this.$color;
    }

    get orientation(): string {
        return this.$orientation;
    }

    get width(): string {
        return <string>this.$width;
    }

    get opacity(): number {
        return this.$opacity;
    }

    set opacity(opacity: number) {
        this.$opacity = opacity;
    }

    get text(): Text {
        return this.$text;
    }

    set text(text: Text) {
        this.$text = text;
    }

    setContent(text: string): void {
        this.$text.text = text;
    }

    /**
     * Sets logo's color, and sets a proper color for the text.
     * @param {string} color
     * */
    setColor(color: string): void {
        this.$color = color;
        this.$text.color = color !== "white" ? "#676c72" : "#FFFFFF";
        this.$updateImagePath();
    }

    /**
     * Sets the orientation of the logo, and updates the logo's image path.
     * @param {string} orientation "vertical" or "horizontal"
     * */
    setOrientation(orientation: string): void {
        if (orientation === "vertical" || orientation === "horizontal") {
            this.$orientation = orientation;
            this.$width = orientation === "vertical" ? "500px" : "1050px";
            this.$updateImagePath();
            // adapt text to the new orientation
            this.$text.changeOrientation(orientation);
        }
    }

    changeOrientation(): void {
        this.setOrientation(this.$orientation === "vertical" ? "horizontal" : "vertical");
    }

    $updateImagePath(): void {
        this.$image_path = `/${this.orientation.charAt(0)}-${this.color}.png`;
    }
}