let button = document.getElementById("button")
let pHeader = document.querySelector(".popup-header")
button.onclick = () => {
    let inputs = document.querySelectorAll(".container > input")

    let data = {}

    for (let i = 0; i < inputs.length; i++) {
        data[inputs[i].name] = inputs[i].value
    }
    Authorization(data)
}

function Authorization(data) {
    let xhr = new XMLHttpRequest()
    xhr.open("POST", "/session")
    xhr.onload = function (e) {
        console.log(e.currentTarget.status)
        if (e.currentTarget.status !== 200) {
            let response = JSON.parse(e.currentTarget.response)
            CreateNotification(response.error)
            if (pHeader.children.length > 1) {
                pHeader.firstChild.remove()
            }
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}

function CreateMessage(msg) {
    let message = document.createElement("div")
    message.classList.toggle("notification-message")
    message.innerText = msg
    return message;
}

function CreateNotification(msg) {
    let notification = document.createElement("div")
    notification.className = "notification"
    notification.appendChild(CreateMessage(msg))
    FadeIn(notification)
    document.body.appendChild(notification)
}

function FadeIn(msg) {
    setTimeout(() => {
        msg.opacity = 0
        setTimeout(() => {
            msg.remove()
        }, 1000)
    }, 5000)
}