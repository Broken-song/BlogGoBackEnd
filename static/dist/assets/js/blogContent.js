 function getContent() {
    // 博客内容
    let currentPage = localStorage.getItem('page');
    // let itemsPerPage = parseInt(document.getElementById('limit').value);  
    let itemsPerPage = 5;
    // 发送请求给后端  
    let xhr = new XMLHttpRequest();  
    xhr.open('POST', '/api/data_get');
    xhr.setRequestHeader("Content-Type", "application/json")
    const content = {
        page: parseInt(currentPage),
        limit: itemsPerPage
    }
    const jsonContent = JSON.stringify(content);
    // 处理后端返回的数据并显示在页面上
    xhr.onreadystatechange = function() {  
        if (xhr.readyState === 4 && xhr.status === 200) {    
            let data = JSON.parse(xhr.response);
            let pagecount = (data.data.total_count / itemsPerPage) - 1
            if((data.data.total_count % itemsPerPage) != 0)
            {
                pagecount += 1;
            }
            for (let i = 1; i < pagecount; i++) {
                if(page.options[page.options.length - 1].value < pagecount){
                    let pageOption = document.createElement('option')
                    pageOption.text = i + 1
                    pageOption.value = i + 1
                    page.appendChild(pageOption)
                }
            }
            let pageBody = document.querySelector('.mainbody')
            pageBody.innerHTML = ""
            for (let i = 0; i < data.data.page_count; i++) {
                let item = data.data.content[i]; 
                let pageBlogCard = document.createElement('div')
                pageBlogCard.className = 'blogcard'
                pageBody.append(pageBlogCard)
                let pageContentBox = document.createElement('div')
                pageContentBox.className = "cont_box"
                let title = document.createElement('text')
                title.id = "title"
                title.textContent = item.title
                let content = document.createElement('text')
                content.id = "content"
                content.textContent = item.content
                let tag = document.createElement('text')
                tag.id = "tag"
                tag.textContent = item.tag
                let time = document.createElement('text')
                time.id = "time"
                time.textContent = "创建时间：" + formatDateTime(item.created_at) + " 最近改动时间：" + formatDateTime(item.updated_at)
                pageContentBox.append(title, content, tag, time)
                pageBlogCard.append(pageContentBox)
            }
        }  
    }
    xhr.send(jsonContent)
}