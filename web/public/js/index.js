var hamburger = document.querySelector(".hamburger"), navMenu = document.querySelector("nav"), activeMenu = document.querySelector("nav a"); hamburger.addEventListener("click", mobileMenu); function mobileMenu() { hamburger.classList.toggle("active"); navMenu.classList.toggle("active") } window.onload = function () { for (var a = document.getElementById("nav").getElementsByTagName("a"), b = 0, c = a.length, d = location.href.split("#")[0]; b < c; b++)a[b].href.split("#")[0] == d && (a[b].className += "active") }; var slideIndex = 1; showSlides(slideIndex);
function plusSlides(a) { showSlides(slideIndex += a) } function currentSlide(a) { showSlides(slideIndex = a) } function showSlides(a) { var b = document.getElementsByClassName("mySlides"), c = document.getElementsByClassName("dot"); a > b.length && (slideIndex = 1); 1 > a && (slideIndex = b.length); for (a = 0; a < b.length; a++)b[a].style.display = "none"; for (a = 0; a < c.length; a++)c[a].className = c[a].className.replace(" active-slide", ""); b[slideIndex - 1].style.display = "block"; c[slideIndex - 1].className += " active-slide" };