/* When the user clicks on the button, 
toggle between hiding and showing the dropdown content */
function Drop() {
    document.getElementById("Dropdown").classList.toggle("show");
}

// Close the dropdown if the user clicks outside of it
window.onclick = function(e) {
    if (!e.target.matches('.drop-button')) {
        let Dropdown = document.getElementById("Dropdown");
        if (Dropdown.classList.contains('show')) {
            Dropdown.classList.remove('show');
        }
    }
}