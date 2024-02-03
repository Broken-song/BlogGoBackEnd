// 导航栏滚动隐藏
function navbarHide () {
    const header = document.querySelector('header')
    if(window.pageYOffset > 0)
    {
        header.setAttribute("style", "top: -61.5px;")
    }else{
        header.setAttribute("style", "")
    }
}