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
    var stage, layer, canvas, context, aspectRatio, imgStageWidth, imgStageHeight, currentGroup;
    export default {
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

                var imageObj = new Image();
                imageObj.onload = function () {
                    aspectRatio = stage.width()/stage.height();
                    imgStageWidth = imageObj.width/imageObj.height >= aspectRatio ? stage.width() : imageObj.width/imageObj.height*stage.height();
                    imgStageHeight = imageObj.width/imageObj.height >= aspectRatio ? imageObj.height/imageObj.width*stage.width() : stage.height();
                    var yoda = new Konva.Image({
                        x: 0,
                        y: 0,
                        image: imageObj,
                        width: imgStageWidth,
                        height: imgStageHeight,
//                        width: stage.width(),
//                        height: imageObj.height/imageObj.width*stage.width(),
    //                    width: imageObj.width/imageObj.height*stage.height(),
    //                    height: stage.height(),
                    });
                    // add the shape to the layer
                    canvas.width = imgStageWidth;
                    canvas.height = imgStageHeight;

                    var canvasImage = new Konva.Image({
                        image: canvas,
                        x: 0,
                        y: 0
                    });
                    layer.add(yoda);
                    layer.add(canvasImage);
                    layer.draw();
                };
                imageObj.src = 'http://www.bz55.com/uploads/allimg/150306/139-1503061IR6.jpg';

                canvas = document.createElement('canvas');

    //            layer.add(image);
    //            layer.draw();
    //            stage.draw();
                context = canvas.getContext("2d");

                stage.addEventListener("mousedown", mouseDown, false);
                stage.addEventListener("mousemove", mouseXY, false);
                stage.addEventListener("mouseup", mouseUp, false);
            }
        }
    };
    var startX, endX, startY, endY;
    var mouseIsDown, mouseIsInGroup;

    function mouseDown(eve) {
        if (!mouseIsInGroup) {
            mouseIsDown = true;
            var pos = getMousePos(canvas, eve);
            startX = endX = pos.x;
            startY = endY = pos.y;
            drawSquare(); //update
        }
    }

    function mouseXY(eve) {
        if (mouseIsDown && !mouseIsInGroup) {
            var pos = getMousePos(canvas, eve);
            endX = pos.x;
            endY = pos.y;
            drawSquare();
        }
    }

    function mouseUp(eve) {
        if (mouseIsDown && !mouseIsInGroup) {
            mouseIsDown = false;
            var pos = getMousePos(canvas, eve);
            endX = pos.x;
            endY = pos.y;
            drawSquare(); //update on mouse-up

            console.log("mouse up！！！！");

            var rect = new Konva.Rect({
                width: endX - startX,
                height: endY - startY,
                stroke: 'red',
                strokeWidth: 1,
                dash: [5, 5]
            });
            rect.on('mouseover', function () {
                document.body.style.cursor = 'move';
            });
            rect.on('mouseout', function () {
                document.body.style.cursor = 'default';
            });

            var minX = canvas.getBoundingClientRect().left;
            var maxX = canvas.getBoundingClientRect().left + imgStageWidth - rect.width();
            var minY = canvas.getBoundingClientRect().top;
            var maxY = canvas.getBoundingClientRect().top + imgStageHeight - rect.height();
            var rectGroup = new Konva.Group({
                x: startX,
                y: startY,
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
                console.log("x:" + this.getX(), "y:" + this.getY());
            });
            rectGroup.on('mouseover', function () {
                mouseIsInGroup = true;
            });
            rectGroup.on('mouseout', function () {
                mouseIsInGroup = false;
            });
            rectGroup.on('mousedown', function (e) {
                if(e.evt.button == 2){
                    this.destroy();
                    layer.draw();
                    mouseIsInGroup = false;
                }
            });
            rect.on('click', function () {
                for(var i = 0; i < layer.get('Group').length; i++) {
                    let group = layer.get('Group')[i];
                    let rect = group.get('Rect')[0];
                    rect.setStroke('red');
                }
                this.stroke('yellow');
                layer.draw();
                currentGroup = this.getParent();
            });
            rectGroup.add(rect);
            layer.add(rectGroup);
            addAnchor(rectGroup, 0, 0, 'topLeft');
            addAnchor(rectGroup, endX - startX, 0, 'topRight');
            addAnchor(rectGroup, endX - startX, endY - startY, 'bottomRight');
            addAnchor(rectGroup, 0, endY - startY, 'bottomLeft');

            context.clearRect(0, 0, canvas.width, canvas.height);
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

        context.clearRect(0, 0, canvas.width, canvas.height);

        context.beginPath();
        context.rect(startX + offsetX, startY + offsetY, width, height);
        context.lineWidth = 1;
        context.strokeStyle = 'red';
        context.setLineDash([5,5]);
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
        // update anchor positions
        switch (activeAnchor.getName()) {
            case 'topLeft':
                topRight.setY(anchorY);
                bottomLeft.setX(anchorX);
                break;
            case 'topRight':
                topLeft.setY(anchorY);
                bottomRight.setX(anchorX);
                break;
            case 'bottomRight':
                bottomLeft.setY(anchorY);
                topRight.setX(anchorX);
                break;
            case 'bottomLeft':
                bottomRight.setY(anchorY);
                topLeft.setX(anchorX);
                break;
        }
        rect.position(topLeft.position());
        var width = topRight.getX() - topLeft.getX();
        var height = bottomLeft.getY() - topLeft.getY();
        if (width && height) {
            console.log("width:" + width, "height:" + height);
            rect.width(width);
            rect.height(height);
            var minX = canvas.getBoundingClientRect().left;
            var maxX = canvas.getBoundingClientRect().left + imgStageWidth - rect.width();
            var minY = canvas.getBoundingClientRect().top;
            var maxY = canvas.getBoundingClientRect().top + imgStageHeight - rect.height();
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
    }
    function addAnchor(group, x, y, name) {
        var stage = group.getStage();
        var layer = group.getLayer();
        var minX = canvas.getBoundingClientRect().left;
        var maxX = canvas.getBoundingClientRect().right + imgStageWidth;
        var minY = canvas.getBoundingClientRect().top;
        var maxY = canvas.getBoundingClientRect().top + imgStageHeight;
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
                var X=pos.x;
                var Y=pos.y;
                if(X<minX){X=minX;}
                if(X>maxX){X=maxX;}
                if(Y<minY){Y=minY;}
                if(Y>maxY){Y=maxY;}
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
            var layer = this.getLayer();
            switch (this.getName()) {
                case 'topLeft':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'topRight':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'bottomRight':
                    document.body.style.cursor = 'crosshair';
                    break;
                case 'bottomLeft':
                    document.body.style.cursor = 'crosshair';
                    break;
            }
            this.strokeWidth(2);
            layer.draw();
        });
        anchor.on('mouseout', function () {
            var layer = this.getLayer();
            document.body.style.cursor = 'default';
            this.strokeWidth(1);
            layer.draw();
        });
        group.add(anchor);
    }
</script>