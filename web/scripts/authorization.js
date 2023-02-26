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
        if (e.currentTarget.status != 200) {
            let response = JSON.parse(e.currentTarget.response)
            console.log(response)
            CreateP(response.error)
            if (pHeader.children.length > 1) {
                pHeader.firstChild.remove()
            }
        } else {
            window.location.href = '/'
        }
    }
    xhr.send(JSON.stringify(data))
}

function CreateP(text) {
    let elem = document.createElement("p")
    let elemText = document.createTextNode(text)
    elem.appendChild(elemText)
    pHeader.appendChild(elem)
    FadeIn(elem, 3000, 60)
}

function FadeIn(elem, t, f) {
    let fps = f || 50;
    let time = t || 500;
    let steps = time / (1000 / fps);
    let op = 1;
    let d0 = op / steps;
    let timer = setInterval(function(){
        op -= d0;
        elem.style.opacity = op;
        steps--;
        if(steps <= 0){
            elem.remove()
            clearInterval(timer);
            elem.style.display = 'none';
        }
    }, (1000 / fps));
}