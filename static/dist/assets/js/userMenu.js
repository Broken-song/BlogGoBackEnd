function menuPosition () {
    let navbar = document.querySelector('.navbar')
    let navbarRect = navbar.getBoundingClientRect() 
    let menuLeft = navbarRect.width + navbarRect.left - 78
    userMenu.style.left = menuLeft + "px"
}