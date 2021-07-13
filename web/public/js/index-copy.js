var hambur = document.querySelector(".hamburger");
var navMenu = document.querySelector("nav");
var activeMenu = document.querySelector("nav a");

hambur.addEventListener("click", mobileMenu);

function mobileMenu() {
    hambur.classList.toggle("active");
    navMenu.classList.toggle("active");
}

window.onload = function () {
    var all_links = document.getElementById("nav").getElementsByTagName("a"),
        i = 0, len = all_links.length,
        full_path = location.href.split('#')[0]; //Ignore hashes?

    // Loop through each link.
    for (; i < len; i++) {
        if (all_links[i].href.split("#")[0] == full_path) {
            all_links[i].className += "active";
        }
    }
}

// When the user scrolls the page, execute myFunction
window.onscroll = function() {onSticky(), scrollFunction()};

// Get the navbar
var navbar = document.querySelector("header");

// Get the offset position of the navbar
var sticky = navbar.offsetTop = 10;

// Add the sticky class to the navbar when you reach its scroll position. Remove "sticky" when you leave the scroll position
function onSticky() {
  if (window.pageYOffset >= sticky) {
    navbar.classList.add("sticky")
  } else {
    navbar.classList.remove("sticky");
  }
}

// Get the button:
mybutton = document.getElementById("myBtn");

// When the user scrolls down 20px from the top of the document, show the button
// window.onscroll = function() {scrollFunction()};

function scrollFunction() {
  if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
    mybutton.style.display = "block";
  } else {
    mybutton.style.display = "none";
  }
}

// When the user clicks on the button, scroll to the top of the document
function topFunction() {
  document.body.scrollTop = 0; // For Safari
  document.documentElement.scrollTop = 0; // For Chrome, Firefox, IE and Opera
}