// 登出
function logoutFunc (){
    // let currentURL = window.location.href; 
    // let rootURL = currentURL.split('/')[0];  
    let xhr = new XMLHttpRequest()
    xhr.open('post', rootURL + 'api/auth/logout')
    xhr.setRequestHeader("Content-Type", "application/json")
    let authToken = localStorage.getItem('token')
    xhr.setRequestHeader('Authorization', 'Bearer ' + authToken); 
    xhr.onload = function () {
        let response;  
            try {  
            response = JSON.parse(xhr.response); 
            if(response.error_code){  
                console.log('Error Code:', response.error_code);
                if(response.error_code == 41000){
                    localStorage.clear()
                    location.reload()
                }
            }
        } catch (error) {  
            console.error('Error parsing response:', error);  
        }  
    }
    xhr.send()
}