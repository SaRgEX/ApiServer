import {Error} from "./error.js";
let button = document.getElementById("button")
let cont = document.body.querySelector(".container");
button.onclick = () => {
    let inputs = document.querySelectorAll(".container-body > form > input")

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
        console.log(data)
        if (e.currentTarget.status !== 200) {
            let response = JSON.parse(e.currentTarget.responseText)
            let err = new Error(response.error, cont, button);
            err.CreateNotification()
            console.log(err)
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}