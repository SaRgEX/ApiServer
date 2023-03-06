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

    Timer(progressBar) {
        let interval = setInterval(() => {
            if (progressBar.value > 0) {
                progressBar.value--;
            } else {
                clearInterval(interval)
            }
        }, 30)

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
        let fade = new Fade(message)
        fade.In()
        return message;
    }

    CreateNotification() {
        let notification = document.body.querySelector(".notification")
        if(!notification) {
            notification = document.createElement("div")
            notification.className = "notification"
            console.log(notification)
        }
        notification.appendChild(this.CreateMessage())
        document.body.appendChild(notification)
        return notification
    }

    ErrorData() {
        this.button.disabled = "true";
        this.container.setAttribute("style", "border-color: #c03232;" +
            " box-shadow: 4px 4px 20px #ff0000")
    }

}