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
    data[group.name] = Number(group.value)
    console.log(data)
    Registration(data)
}

function Registration(data) {
    let xhr = new XMLHttpRequest()
    xhr.open("POST", "/student")
    xhr.onload = function (e) {
        console.log(e.currentTarget.status)
        if (e.currentTarget.status !== 201) {
            let response = JSON.parse(e.currentTarget.response)
            let err = new Error(response.error, container, button);
            err.CreateNotification()
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}

(function () {
    let x = new XMLHttpRequest();
    x.open("GET", "/group", true);
    console.log(x)
    x.onload = function (e) {
        let response = e.currentTarget.response;
        // преобразование ответа сервера в массив
        let options = JSON.parse(response).Group;
        let select = document.getElementById('group');
        // очистка списка
        while (select.options.length > 0) {
            select.options.remove(0);
        }
        for (let i = 0; i <= options.length; i++) {
            select.options.add(new Option(options[i].id, options[i].id));
            select.selectedIndex = -1
        }
    };
    x.send();
})()