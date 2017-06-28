<style scoped>
    .resize-drag {
      color: red;
      font-size: 20px;
      font-family: sans-serif;
      border-radius: 0px;
      padding: 20px;
      margin: 30px 20px;
      border: dotted yellow 1px;
      width: 120px;
      /* This makes things *much* easier */
      box-sizing: border-box;
    }

    .resize-container {
      width: 100%;
      height: 240px;
    }
</style>
<template>
    <div id="main" style="height: 100%;">
    </div>
</template>
<script>
    import Konva from 'konva';
    var stage,
        layer,
        drawCanvas,
        context,
        aspectRatio,//画布区域宽高比
        zoomRatio, //图片缩放比
        imageObj,
        imgStageWidth,
        imgStageHeight,
        currentGroup;
    var demarcate =  {
        name: 'demarcate',
        data() {
            return {
//                stageWidth: 5/6*window.innerWidth,
//                stageHeight: window.innerHeight - 60,
//                aspectRatio : this.stageWidth/this.stageHeight,
            };
        },
        mounted() {
            this.stageCanvas();
        },
        beforeDestroy() {

        },
        methods: {
            stageCanvas: function () {
                stage = new Konva.Stage({
                    container: 'main',
                    width: 5/6*window.innerWidth,
                    height: window.innerHeight - 60,
                });
                // add canvas element
                layer = new Konva.Layer();
                stage.add(layer);

                imageObj = new Image();

                var yoda = new Konva.Image({
                    x: 0,
                    y: 0,
//                    image: imageObj,
//                    width: imgStageWidth,
//                    height: imgStageHeight,
                });

                var canvasImage = new Konva.Image({
//                    image: drawCanvas,
                    x: 0,
                    y: 0
                });
                layer.add(yoda);
                layer.add(canvasImage);
    //            layer.add(image);
    //            layer.draw();
    //            stage.draw();

                drawCanvas = document.createElement('canvas');
                context = drawCanvas.getContext("2d");

                stage.addEventListener("mousedown", mouseDown, false);
                stage.addEventListener("mousemove", mouseXY, false);
                stage.addEventListener("mouseup", mouseUp, false);

                imageObj.src = 'http://www.bz55.com/uploads/allimg/150306/139-1503061IR6.jpg';
                imageObj.onload = function () {
                    aspectRatio = stage.width()/stage.height();
                    imgStageWidth = imageObj.width/imageObj.height >= aspectRatio ? stage.width() : imageObj.width/imageObj.height*stage.height();
                    imgStageHeight = imageObj.width/imageObj.height >= aspectRatio ? imageObj.height/imageObj.width*stage.width() : stage.height();
                    zoomRatio = imageObj.width/imgStageWidth;

                    yoda.width(imgStageWidth);
                    yoda.height(imgStageHeight);
                    yoda.image(imageObj);
                    drawCanvas.width = imgStageWidth;
                    drawCanvas.height = imgStageHeight;
                    canvasImage.image(drawCanvas);

                    if (layer.get('Group').length > 0) {
                        for(var i = layer.get('Group').length - 1; i >= 0 ; i--) {
                            layer.get('Group')[i].destroy();;
                        }
                        currentGroup = null;
                        mouseIsInGroup = false;
                    }
                    layer.draw();
                };
            }
        }
    };
    var startX, endX, startY, endY;
    var mouseIsDown, mouseIsInGroup;

    function mouseDown(eve) {
        if (!mouseIsInGroup) {
            mouseIsDown = true;
            var pos = getMousePos(drawCanvas, eve);
            startX = endX = pos.x;
            startY = endY = pos.y;
            drawSquare(); //update
        }
    }

    function mouseXY(eve) {
        if (mouseIsDown && !mouseIsInGroup) {
            var pos = getMousePos(drawCanvas, eve);
            endX = pos.x;
            endY = pos.y;
            drawSquare();
        }
    }

    function mouseUp(eve) {
        if (mouseIsDown && !mouseIsInGroup) {
            mouseIsDown = false;
            var pos = getMousePos(drawCanvas, eve);
            endX = pos.x;
            endY = pos.y;
            drawSquare(); //update on mouse-up

            var rect = new Konva.Rect({
                width: startX <= endX ? endX - startX : startX - endX,
                height: startY <= endY ? endY - startY : startY -endY,
                stroke: 'red',
                strokeWidth: 1
//                dash: [5, 5]
            });
            rect.on('mouseover', function () {
                document.body.style.cursor = 'move';
            });
            rect.on('mouseout', function () {
                document.body.style.cursor = 'default';
            });

            var minX = drawCanvas.getBoundingClientRect().left;
            var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth - rect.width();
            var minY = drawCanvas.getBoundingClientRect().top;
            var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight - rect.height();
            var rectGroup = new Konva.Group({
                x: startX <= endX ? startX : endX,
                y: startY <= endY ? startY : endY,
                draggable: true,
                dragBoundFunc: function (pos) {
                    var X=pos.x;
                    var Y=pos.y;
                    if(X<minX){X=minX;}
                    if(X>maxX){X=maxX;}
                    if(Y<minY){Y=minY;}
                    if(Y>maxY){Y=maxY;}
                    return({x:X, y:Y});
                }
            });
            rectGroup.on('dragmove', function () {
//                console.log("x:" + this.getX(), "y:" + this.getY());
                console.log("originX:", this.getX()*zoomRatio, "originY:", this.getY()*zoomRatio);
            });
            rectGroup.on('mouseover', function () {
                mouseIsInGroup = true;
            });
            rectGroup.on('mouseout', function () {
                mouseIsInGroup = false;
            });
            rectGroup.on('mousedown', function (e) {
                if(e.evt.button == 2){
                    if (currentGroup == this){
                        currentGroup = null;
                    }
                    this.destroy();
                    layer.draw();
                    mouseIsInGroup = false;
                }
            });
            rect.on('click', function () {
                for(var i = 0; i < layer.get('Group').length; i++) {
                    var group = layer.get('Group')[i];
                    var rect = group.get('Rect')[0];
                    rect.setStroke('red');
                }
                this.stroke('yellow');
                layer.draw();
                currentGroup = this.getParent();
            });
            rectGroup.add(rect);
            layer.add(rectGroup);
                addAnchor(rectGroup, 0, 0, 'topLeft');
                addAnchor(rectGroup, startX <= endX ? endX - startX : startX - endX, 0, 'topRight');
                addAnchor(rectGroup, 0, startY <= endY ? endY - startY : startY -endY, 'bottomLeft');
                addAnchor(rectGroup, startX <= endX ? endX - startX : startX - endX, startY <= endY ? endY - startY : startY -endY, 'bottomRight');

            context.clearRect(0, 0, drawCanvas.width, drawCanvas.height);
            layer.draw();
        }
    }

    function drawSquare() {
        // creating a square
        var w = endX - startX;
        var h = endY - startY;
        var offsetX = (w < 0) ? w : 0;
        var offsetY = (h < 0) ? h : 0;
        var width = Math.abs(w);
        var height = Math.abs(h);

        context.clearRect(0, 0, drawCanvas.width, drawCanvas.height);

        context.beginPath();
        context.rect(startX + offsetX, startY + offsetY, width, height);
        context.lineWidth = 1;
        context.strokeStyle = 'red';
//        context.setLineDash([5,5]);
        context.stroke();
        layer.draw();
    }

    function getMousePos(canvas, evt) {
        var rect = canvas.getBoundingClientRect();
        return {
            x: evt.layerX - rect.left,
            y: evt.layerY - rect.top
        };
    }

    function update(activeAnchor) {
        var group = activeAnchor.getParent();
        var topLeft = group.get('.topLeft')[0];
        var topRight = group.get('.topRight')[0];
        var bottomRight = group.get('.bottomRight')[0];
        var bottomLeft = group.get('.bottomLeft')[0];
        var rect = group.get('Rect')[0];
        var anchorX = activeAnchor.getX();
        var anchorY = activeAnchor.getY();
        var tmp;

//        console.log("anchorX:"+anchorX,"anchorY:"+anchorY);
        switch (activeAnchor.getName()) {
            case 'topLeft':
                group.setX(group.getX() + anchorX);
                group.setY(group.getY() + anchorY);
                topLeft.position({x:0, y:0});
                topRight.position({x:topRight.getX() - anchorX, y:0});
                bottomLeft.position({x:0, y:bottomLeft.getY() - anchorY});
                bottomRight.position({x:bottomRight.getX() - anchorX, y:bottomRight.getY() - anchorY});
                break;
            case 'topRight':
                group.setY(group.getY() + anchorY);
                topLeft.position({x:0, y:0});
                topRight.position({x:anchorX, y:0});
                bottomLeft.position({x:0, y:bottomLeft.getY()-anchorY});
                bottomRight.position({x:anchorX, y:bottomRight.getY()-anchorY});
                break;
            case 'bottomLeft':
                group.setX(group.getX() + anchorX);
                topLeft.position({x:0, y:0});
                topRight.position({x:topRight.getX() - anchorX, y:0});
                bottomLeft.position({x:0, y:anchorY});
                bottomRight.position({x:bottomRight.getX() - anchorX, y:anchorY});
                break;
            case 'bottomRight':
                topLeft.position({x:0, y:0});
                topRight.position({x:anchorX, y:0});
                bottomLeft.position({x:0, y:anchorY});
                bottomRight.position({x:anchorX, y:anchorY});
                break;
        }


        rect.position(topLeft.position());
        var width = topRight.getX() - topLeft.getX();
        var height = bottomLeft.getY() - topLeft.getY();
//        console.log("topLeft.getX():"+topLeft.getX(),"topLeft.getY():"+topLeft.getY());
//        console.log("topRight.getX():"+topRight.getX(),"topRight.getY():"+topRight.getY());
//        console.log("bottomLeft.getX():"+bottomLeft.getX(),"bottomLeft.getY():"+bottomLeft.getY());
//        console.log("bottomRight.getX():"+bottomRight.getX(),"bottomRight.getY():"+bottomRight.getY());
        console.log("width:" + width, "height:" + height);
        console.log("originWidth:", width*zoomRatio, "originHeight:", height*zoomRatio);
        rect.width(width);
        rect.height(height);
        var minX = drawCanvas.getBoundingClientRect().left;
        var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth - rect.width();
        var minY = drawCanvas.getBoundingClientRect().top;
        var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight - rect.height();
        group.dragBoundFunc(function (pos) {
            var X=pos.x;
            var Y=pos.y;
            if(X<minX){X=minX;}
            if(X>maxX){X=maxX;}
            if(Y<minY){Y=minY;}
            if(Y>maxY){Y=maxY;}
            return({x:X, y:Y});
        });
    }
    function addAnchor(group, x, y, name) {
        var stage = group.getStage();
        var layer = group.getLayer();
        var minX = drawCanvas.getBoundingClientRect().left;
        var maxX = drawCanvas.getBoundingClientRect().right + imgStageWidth;
        var minY = drawCanvas.getBoundingClientRect().top;
        var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight;
        var anchor = new Konva.Circle({
            x: x,
            y: y,
            stroke: '#666',
            fill: '#ddd',
            strokeWidth: 1,
            radius: 4,
            name: name,
            draggable: true,
            dragOnTop: false,
            dragBoundFunc: function (pos) {
                var X = pos.x;
                var Y = pos.y;
                var p = this.getParent();
                var r = p.get('Rect')[0];
//                console.log("r:",r.width(),r.height());
                switch (this.getName()) {
                    case 'topLeft':
                        if(X<minX){X=minX;}
                        if(X>p.getX()+r.width()){X=p.getX()+r.width();}
                        if(Y<minY){Y=minY;}
                        if(Y>p.getY()+r.height()){Y=p.getY()+r.height();}
                        break;
                    case 'topRight':
                        if(X<p.getX()){X=p.getX();}
                        if(X>maxX){X=maxX;}
                        if(Y<minY){Y=minY;}
                        if(Y>p.getY()+r.height()){Y=p.getY()+r.height();}
                        break;
                    case 'bottomLeft':
                        if(X<minX){X=minX;}
                        if(X>p.getX()+r.width()){X=p.getX()+r.width();}
                        if(Y<p.getY()){Y=p.getY();}
                        if(Y>maxY){Y=maxY;}
                        break;
                    case 'bottomRight':
                        if(X<p.getX()){X=p.getX();}
                        if(X>maxX){X=maxX;}
                        if(Y<p.getY()){Y=p.getY();}
                        if(Y>maxY){Y=maxY;}
                        break;
                }
                return({x:X, y:Y});
            }
        });
        anchor.on('dragmove', function () {
            update(this);
            layer.draw();
        });
        anchor.on('mousedown touchstart', function () {
            group.setDraggable(false);
            this.moveToTop();
        });
        anchor.on('dragend', function () {
            group.setDraggable(true);
            layer.draw();
        });
        // add hover styling
        anchor.on('mouseover', function () {
            switch (this.getName()) {
                case 'topLeft':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'topRight':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'bottomLeft':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'bottomRight':
                    document.body.style.cursor = 'crosshair';
                    break;
            }
            this.strokeWidth(2);
            layer.draw();
        });
        anchor.on('mouseout', function () {
            document.body.style.cursor = 'default';
            this.strokeWidth(1);
            layer.draw();
        });
        group.add(anchor);
    }
    window.addEventListener('keydown',check,true);
    document.oncontextmenu = function(){return false;};
    function check(e) {
        var key = event.which || event.keyCode;
        if (currentGroup) {
            var rect = currentGroup.get('Rect')[0];
            var rectRatio = rect.width()/rect.height();
            var wOffset = rectRatio > 1 ? rectRatio : 1 ;
            var hOffset = rectRatio > 1 ? 1 : rectRatio;
            var topLeft = currentGroup.get('.topLeft')[0];
            var topRight = currentGroup.get('.topRight')[0];
            var bottomRight = currentGroup.get('.bottomRight')[0];
            var bottomLeft = currentGroup.get('.bottomLeft')[0];
            var minX = drawCanvas.getBoundingClientRect().left;
            var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth - rect.width();
            var minY = drawCanvas.getBoundingClientRect().top;
            var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight - rect.height();
            switch (key) {
                case 37://←
                    if ((minX + 1) <= currentGroup.getX())
                        currentGroup.setX(currentGroup.getX() - 1);
                    break;
                case 38://↑
                    if ((minY + 1) <= currentGroup.getY())
                        currentGroup.setY(currentGroup.getY() - 1);
                    break;
                case 39://→
                    if (currentGroup.getX() <= (maxX - 1))
                        currentGroup.setX(currentGroup.getX() + 1);
                    break;
                case 40://↓
                    if (currentGroup.getY() <= (maxY - 1))
                        currentGroup.setY(currentGroup.getY() + 1);
                    break;
                case 67://c
                    if ((minX + 1) <= currentGroup.getX()) {
                        topLeft.setX(topLeft.getX() - 1);
                        bottomLeft.setX(bottomLeft.getX() - 1);
                        update(topLeft);
                        layer.draw();
                    }
                    if ((minY + 1) <= currentGroup.getY()) {
                        topLeft.setY(topLeft.getY() - 1);
                        topRight.setY(topRight.getY()  -1);
                        update(topLeft);
                        layer.draw();
                    }
                    if (currentGroup.getX() <= (maxX - 1)) {
                        topRight.setX(topRight.getX() + 1);
                        bottomRight.setX(bottomRight.getX() + 1);
                        update(topLeft);
                        layer.draw();
                    }
                    if (currentGroup.getY() <= (maxY - 1)) {
                        bottomLeft.setY(bottomLeft.getY() + 1);
                        bottomRight.setY(bottomRight.getY() + 1);
                        update(topLeft);
                        layer.draw();
                    }
                    break;
                case 68://d
                    if (topLeft.getX() + wOffset * 2 <= topRight.getX()) {
                        topLeft.setX(topLeft.getX() + wOffset);
                        topRight.setX(topRight.getX() - wOffset);
                        bottomLeft.setX(bottomLeft.getX() + wOffset);
                        bottomRight.setX(bottomRight.getX() - wOffset);
                        update(topLeft);
                        update(topRight);
                        update(bottomLeft);
                        update(bottomRight);
                        layer.draw();
                    }
                    if (topLeft.getY() + hOffset * 2 <= bottomLeft.getY()) {
                        topLeft.setY(topLeft.getY() + hOffset);
                        topRight.setY(topRight.getY() + hOffset);
                        bottomLeft.setY(bottomLeft.getY() - hOffset);
                        bottomRight.setY(bottomRight.getY() - hOffset);
                        update(topLeft);
                        update(topRight);
                        update(bottomLeft);
                        update(bottomRight);
                        layer.draw();
                    }
                    break;
                case 65://a
                    break;
            }
            console.log("originX:", currentGroup.getX()*zoomRatio, "originY:", currentGroup.getY()*zoomRatio);
            layer.draw();
        }
        if (key == 83) {//s
            //Load next image here!!
            imageObj.src = 'https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1499218671&di=fd164f94176ec4ee982db63abffe0df9&imgtype=jpg&er=1&src=http%3A%2F%2Fbizhi.zhuoku.com%2F2013%2F07%2F22%2Fhua%2Fyihuayishijie68.jpg';
        }
//                alert(e.keyCode);
    }
    export default demarcate;
    Array.prototype.max = function(){
      return Math.max.apply({},this)
    };
    Array.prototype.min = function(){
      return Math.min.apply({},this)
    };
</script>