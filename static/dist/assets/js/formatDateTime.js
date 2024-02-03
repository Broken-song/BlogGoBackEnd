function formatDateTime(dateTimeString) {  
    // 创建一个 Date 对象  
    const date = new Date(dateTimeString);  
    
    // 获取日期和时间的各个组成部分  
    const year = date.getFullYear();  
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份是从 0 开始的，所以需要 +1，并且可能是单个数字，所以需要补零  
    const day = String(date.getDate()).padStart(2, '0'); // 天可能是单个数字，所以需要补零  
    const hours = String(date.getHours()).padStart(2, '0');  
    const minutes = String(date.getMinutes()).padStart(2, '0');  
    const seconds = String(date.getSeconds()).padStart(2, '0');  
    
    // 返回格式化的日期时间字符串  
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;  
  }  