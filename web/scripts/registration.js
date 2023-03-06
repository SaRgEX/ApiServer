import {Error} from "./error.js";

let container = document.body.querySelector(".container");
let button = document.getElementById("button")
let inputsInfo = document.querySelectorAll(".info > input")
let inputsLog = document.querySelectorAll(".logging > input")
let group = document.querySelector("select")
f()
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
            let response = JSON.parse(e.currentTarget.response.replace("}{", ", "))
            let err = new Error(response.error, container, button);
            err.CreateNotification()
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}

function f() {
    var x = new XMLHttpRequest();
    x.onload = function (e) {
        let response = e.currentTarget.response;
        console.log(response)
        // преобразование ответа сервера в массив
        let options = JSON.parse(response).Group.id;
        var select = document.getElementById('group');
        console.log(options)
        // очистка списка
        while (select.options.length > 0) {
            select.options.remove(0);
        }
        var option = document.createElement('option');
        option.text = options;
        console.log(option)
        select.options.add(option);
    };
    x.open("GET", "/group", true);
    x.send();
}