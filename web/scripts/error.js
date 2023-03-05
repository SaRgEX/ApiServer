import {Fade} from "./fade.js";

export class Error {
    text
    container
    button

    constructor(text, container, button) {
        this.text = text
        this.container = container
        this.button = button
    }

    CreateProgressBar() {
        let progressBar = document.createElement("progress");
        progressBar.classList.add("notification-progress");
        progressBar.max = 100
        progressBar.value = 100
        this.Timer(progressBar)
        return progressBar;
    }

    CreateMessage() {
        let message = document.createElement("div")
        message.classList.add("notification-message")
        message.innerText = this.text
        message.appendChild(this.CreateProgressBar())
        return message;
    }

    CreateNotification() {
        let notification = document.createElement("div")
        notification.className = "notification"
        notification.appendChild(this.CreateMessage())
        let fade = new Fade(notification)
        fade.In()
        document.body.appendChild(notification)
        return notification
    }

    Timer(progressBar) {
        let interval = setInterval(() => {
            if (progressBar.value > 0) {
                progressBar.value--;
            } else {
                clearInterval(interval)
            }
        }, 30)

    }

    ErrorData() {
        this.button.disabled = "true";
        this.container.setAttribute("style", "border-color: #c03232;" +
            " box-shadow: 4px 4px 20px #ff0000")
    }
}