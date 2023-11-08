//下拉菜单
 function dropMenu(){
    el.classList.toggle('closeable');
    dropDownMenu.classList.toggle('open')
    const isOpen = dropDownMenu.classList.contains('open')
}
//导航栏滚动隐藏
 function navHide(){
    if(window.scrollY > 0)
    {
        header.setAttribute("style", "top: -61.5px;")
    }else{
        header.setAttribute("style", "")
    }
}
// p = function() {
//     var t = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
//       , e = s.startScrollTop;
//     if (t <= 50)
//         return s.headerClass = "fixed-header",
//         void (s.startScrollTop = t);
//     e > t ? s.headerClass = "fixed-header" : e <= t && (s.headerClass = "hide-header"),
//     s.startScrollTop = t
// };