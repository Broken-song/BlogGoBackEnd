 // 获取用户信息
 function getUserInfo () {
    let xhr = new XMLHttpRequest()
        xhr.open('post', rootURL + 'api/auth/info')
        xhr.setRequestHeader("Content-Type", "application/json")
        let authToken = localStorage.getItem('token')
        xhr.setRequestHeader('Authorization', 'Bearer ' + authToken); 
        xhr.onload = function () {
        let response;  
            try {
            response = JSON.parse(xhr.response); 
            if(response.error_code){  
                console.log('Error Code:', response.error_code);
            }else if(response.error_code == 41000)
            {
                logoutFunc
            }
            else{
                if (response.data){
                    localStorage.setItem('uid', response.data.ID)
                    localStorage.setItem('uname', response.data.name)
                    localStorage.setItem('uemail', response.data.email)
                    localStorage.setItem('ucontent', response.data.content)
                    localStorage.setItem('urole', response.data.role)
                }
            }
        } catch (error) {  
            console.error('Error parsing response:', error);  
        }
    }
    xhr.send()
}