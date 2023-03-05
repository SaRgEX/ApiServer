import {Error} from "./error.js";
let container = document.body.querySelector(".container");
let button = document.getElementById("button")
let inputsInfo = document.querySelectorAll(".info > input")
let inputsLog = document.querySelectorAll(".logging > input")
let group = document.querySelector("select")

button.onclick = () => {

    let data = {};

    for (let i = 0; i < inputsLog.length; i++) {
        data[inputsLog[i].name] = inputsLog[i].value
    }
    for (let i = 0; i < inputsInfo.length; i++) {
        data[inputsInfo[i].name] = inputsInfo[i].value
    }
    data[group.name] = group.value
    console.log(data)
    Registration(data)
}

function Registration(data) {
    let xhr = new XMLHttpRequest()
    xhr.open("POST", "/student")
    xhr.onload = function (e) {
        console.log(e.currentTarget.status)
        if (e.currentTarget.status !== 200) {
            console.log(e.currentTarget.response.replace("}{", ", "))
            let response = JSON.parse(e.currentTarget.response.replace("}{", ", "))
            console.log(response)

            console.log(response.error)
            let err = new Error(response.error, container, button);
            err.CreateNotification()
            console.log(err)
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}