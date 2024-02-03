function reachTopScoll () {
    var top = getTop()
    if(top > 80){
        oBtn.style.marginTop = '0px'
    }else{
        oBtn.style.marginTop = '100px'
    }
}
function reachTopClick () {
    var top = getTop()
    var timer = setInterval(function(){
       document.documentElement.scrollTop -= 20
        if(getTop() <= 0){
            clearInterval(timer)
        } 
    },10)    
}
function getTop(){
    return document.documentElement.scrollTop || document.body.scrollTop
}