function loginStatus() {
    if (localStorage.getItem('token')){
        localStorage.setItem('isLogin', 1)
    }else{
        localStorage.setItem('isLogin', 0)
    }
    if(localStorage.getItem('isLogin') == 1){
        isLogin.setAttribute("style","display: block;")
        login_btn.forEach((element, index) => {
            login_btn[index].setAttribute("style","display: none;")
        });
        // 获取用户信息
        if(!localStorage.getItem('uid')){
            loginStatus.onload = getUserInfo()
        }
        // 获取头像链接
        if(!localStorage.getItem('avatar_url')){
            let xhr = new XMLHttpRequest()
            xhr.open('post', rootURL + 'api/avatar_get')
            let avatarGet = new FormData()
            avatarGet.append('business','avatar')
            avatarGet.append('uid',localStorage.getItem('uid'))
            xhr.onload = function () {
                let response;  
                try {  
                    response = JSON.parse(xhr.response); 
                    if(response.error_code){  
                        console.log('Error Code:', response.error_code);
                    }else{
                        if (response.data){
                            localStorage.setItem('avatar_url', response.data)
                        }
                    }
                } catch (error) {  
                    console.error('Error parsing response:', error);  
                }  
            }
            xhr.send(avatarGet)
        }
    }else{
        isLogin.setAttribute("style","display: none;")
        login_btn.forEach((element, index) => {
            login_btn[index].removeAttribute("style","display: none;")
        });
    }
}