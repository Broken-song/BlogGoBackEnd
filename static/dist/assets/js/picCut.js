// 头像上传
var handler = {
    init: function($container){
        //需要把dragover的默认行为禁掉，不然会跳页
        $container.on("dragover", function(event){
            event.preventDefault();
        });
        $container.on("drop", function(event){
            event.preventDefault();
            //这里获取拖过来的图片文件，为一个File对象
            var file = event.originalEvent.dataTransfer.files[0];
            handler.handleDrop($(this), file);
        });
        $container.on("change", "input[type=file]", function(event){
            if(!this.value) return;
            var file = this.files[0];
            handler.handleDrop($(this).closest(".container"), file);
            this.value = "";
        });
    },
    handleDrop: function($container, file){
        var $img =  $container.find("img");
        handler.readImgFile(file, $img, $container);
    },
    readImgFile: function(file, $img, $container){
        var reader = new FileReader(file);
        //检验用户是否选则是图片文件
        if(file.type.split("/")[0] !== "image"){
            util.toast("You should choose an image file");
            return;
        }   
        reader.onload = function(event) {
            var base64 = event.target.result;
            handler.compressAndUpload($img, base64, file,  $container);
        }   
        reader.readAsDataURL(file);
        //获取图片base64内容
        var base64 = event.target.result;
        //如果图片大于1MB，将body置半透明
        if(file.size > ONE_MB){
            $("body").css("opacity", 0.5);
        }
        //因为这里图片太大会被卡一下，整个页面会不可操作
        $img.attr("src", baseUrl);
        //还原
        if(file.size > ONE_MB){
            $("body").css("opacity", 1);
        }
        //然后再调一个压缩和上传的函数
        handler.compressAndUpload($img, file, $container);
    },
    readImgFile: function(file, $img, $container){
        EXIF.getData(file, function(){
            var orientation = this.exifdata.Orientation,
            rotateDeg = 0;
            //如果不是ios拍的照片或者是横拍的，则不用处理，直接读取
            if(typeof orientation === "undefined" || orientation === 1){ 
                //原本的readImgFile，添加一个rotateDeg的参数
                handler.doReadImgFile(file, $img, $container, rotateDeg);
            }   
            //否则用canvas旋转一下
            else{
                rotateDeg = orientation === 6 ? 90*Math.PI/180 : 
                orientation === 8 ? -90*Math.PI/180 :
                orientation === 3 ? 180*Math.PI/180 : 0;
                handler.doReadImgFile(file, $img, $container, rotateDeg);
            }   
        });
    },
    compress: function(img, maxWidth, mimeType){
        //设定图片最大压缩宽度为1500px
        var maxWidth = 1500;
        var resultImg = handler.compress($img[0], maxWidth, file.type);
        //创建一个canvas对象
        var cvs = document.createElement('canvas');
        var width = img.naturalWidth,
        height = img.naturalHeight,
        imgRatio = width / height;
        //如果图片维度超过了给定的maxWidth 1500，
        //为了保持图片宽高比，计算画布的大小
        if(width > maxWidth){
            width = maxWidth;
            height = width / imgRatio;
        }   
        cvs.width = width;
        cvs.height = height;
        //把大图片画到一个小画布
        var ctx = cvs.getContext("2d").drawImage(img, 0, 0, img.naturalWidth, img.naturalHeight, 0, 0, width, height);
        //图片质量进行适当压缩
        var quality = width >= 1500 ? 0.5 :
        width > 600 ? 0.6 : 1;
        //导出图片为base64
        var newImageData = cvs.toDataURL(mimeType, quality);
        var resultImg = new Image();
        resultImg.src = newImageData;
        return resultImg;
    }
}