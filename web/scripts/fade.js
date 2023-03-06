export class Fade {
    message
    constructor(message) {
        this.message = message
    }

    In() {
        this.message.style.opacity = "1"
        setTimeout(() => {
            this.message.style.opacity = "0"
            setTimeout(() => {
                this.message.remove();
                this.removeDiv()
            }, 1000)
        }, 2000);
    }

    Out() {
        this.message.style.opacity = "0"
        setTimeout(() => {
            this.message.style.opacity = "1"
            setTimeout(() => {
            }, 1000)
        }, 2000);
    }

    removeDiv() {
        let notification = document.querySelectorAll(".notification > .notification-message");
        if (notification.length === 0) {
            document.querySelector(".notification").remove()
        }
    }
}