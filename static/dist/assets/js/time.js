
function myTime(){
    let time=new Date();
    let hh=time.getHours();  //时
    let mm=time.getMinutes();  //分
    let ss=time.getSeconds();  //秒
    // Math.floor() 向下取整
    document.getElementById("1").innerText=Math.floor(hh/10);
    document.getElementById("2").innerText=hh%10;
    document.getElementById("4").innerText=Math.floor(mm/10);
    document.getElementById("5").innerText=mm%10;
    document.getElementById("7").innerText=Math.floor(ss/10);
    document.getElementById("8").innerText=ss%10;
}
// 一秒执行一次
setInterval(myTime,1000);