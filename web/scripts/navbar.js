document.querySelector('.navbar-link').onclick = function(e) {
    console.log(e)
    if(e.target.nodeName !== 'IMG' && e.target.nodeName !== 'SPAN') return;
    console.log(e.target.parentElement)

    e.target.parentElement.lastElementChild.classList.toggle('element-active')
}