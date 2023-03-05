export class Fade {
    container
    constructor(container) {
        this.container = container
    }

    In() {
        this.container.style.opacity = "1"
        setTimeout(() => {
            this.container.style.opacity = "0"
            setTimeout(() => {
                this.container.remove();
                this.container.removeAttribute("style")
            }, 1000)
        }, 2000);
    }

    Out() {
        this.container.style.opacity = "0"
        setTimeout(() => {
            this.container.style.opacity = "1"
            setTimeout(() => {
            }, 1000)
        }, 2000);
    }
}