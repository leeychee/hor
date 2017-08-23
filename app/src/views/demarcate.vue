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
  import Vue from 'vue'
  import VueResource from 'vue-resource'
  import iView from 'iview';
  import Bus from '../libs/bus';
  Vue.use(VueResource);
  Vue.use(iView);

  var stage,
      layer,
      drawCanvas,
      context,
      aspectRatio,//画布区域宽高比
      zoomRatio, //图片缩放比
      currentImageId,
      imageObj,
      imgStageWidth,
      imgStageHeight,
      currentGroup,
      echoGroups,//图片中已标记框框
      imageIds = [],//图片id数组
      gIndex = -1,//游标
      minRectSize = 60,
      opType,//demarcate:tag(default),review:review
      chance,//color'variable middle transmit
      setType = 'car',//set vehile type 
      setUser = 'sun'//set user

  var startX, endX, startY, endY;
  var mouseIsDown,
      mouseIsInGroup,
      mouseIsOnCircle,
      mouseIsOnCurrentGroup;

  var demarcate = {
    name: 'demarcate',
    data() {
      return {
//                stageWidth: 5/6*window.innerWidth,
//                stageHeight: window.innerHeight - 60,
//                aspectRatio : this.stageWidth/this.stageHeight,
      };
    },
    props: ['minSize', 'userName', 'preType', 'type'],
    created: function () {
      console.log("type:" + this.type);
      if (this.type == "r") {
        opType = "review";
      } else {
        opType = "tag";
      }
    },
    watch: {
      minSize: function (val) {
        minRectSize = val;
      },
      preType: function (val) {
        setType = val;
      },
      userName: function (val) {
        setUser = val;
      },
      '$route' (to, from) {
        // 对路由变化作出响应...
        console.log("to:", to, ",from:", from);
      }
    },
    mounted() {
      this.stageCanvas();
    },
    beforeDestroy() {

    },
    methods: {
      stageCanvas: function () {
        stage = new Konva.Stage({ //create a stage
          container: 'main',
          width: 22 / 24 * window.innerWidth,
          height: window.innerHeight - 40,
        });
        // add canvas element
        layer = new Konva.Layer(); //创建一个层
        stage.add(layer);

        imageObj = new Image();

        var yoda = new Konva.Image({
          x: 0,
          y: 0,
        });

        var canvasImage = new Konva.Image({
          x: 0,
          y: 0
        });
        layer.add(yoda);
        layer.add(canvasImage);

        drawCanvas = document.createElement('canvas');
        context = drawCanvas.getContext("2d");

        stage.addEventListener("mousedown", mouseDown, false);
        stage.addEventListener("mousemove", mouseXY, false);
        stage.addEventListener("mouseup", mouseUp, false);

        imageIds = [];
        gIndex = -1;

        if (imageIds[gIndex + 1]) {
          Vue.http.get("/image/" + imageIds[gIndex + 1]).then(res => {
            let o = res.body;
            echoGroups = o.objects;
            currentImageId = o.id;
            gIndex++;
            imageObj.src = "/f/" + o.path;
            this.$emit('updateImgName', o.path);
            this.$emit('updateUser', o.demname);
            console.log("/image/:id : ", res.body);
          }, err => {
            console.log("error /image/:id : ", err);
          });
        } else {
          Vue.http.get("/images/_next", {params: {"type": opType}}).then(resp => {
            console.log(resp.body);
            let o = resp.body;
            echoGroups = o.objects;
            currentImageId = o.id;
            imageIds.push(currentImageId);
            gIndex = imageIds.length - 1;
            console.log(imageIds);
            imageObj.src = "/f/" + o.path;
            this.$emit('updateImgName', o.path);
            this.$emit('updateUser', o.demname);
          }, error => {
            if (error.ok == false) {
              switch (error.status) {
                case 404:
                  this.$Message.warning("没有更多图片了");
                  break;
                default:
                  this.$Message.error("服务器好像出了点问题！");
                  break;
              }
            }
            console.log("image error: ", error);
          });
        }
        imageObj.onload = function () {
          stage.removeEventListener("mousedown", mouseDown, false);
          stage.addEventListener("mousedown", mouseDown, false);
          aspectRatio = stage.width() / stage.height();
          imgStageWidth = imageObj.width / imageObj.height >= aspectRatio ? stage.width() : imageObj.width / imageObj.height * stage.height();
          imgStageHeight = imageObj.width / imageObj.height >= aspectRatio ? imageObj.height / imageObj.width * stage.width() : stage.height();
          zoomRatio = imageObj.width / imgStageWidth;

          yoda.width(imgStageWidth);
          yoda.height(imgStageHeight);
          yoda.image(imageObj);
          drawCanvas.width = imgStageWidth;
          drawCanvas.height = imgStageHeight;
          canvasImage.image(drawCanvas);

          if (layer.get('Group').length > 0) {
            for (var i = layer.get('Group').length - 1; i >= 0; i--) {
              layer.get('Group')[i].destroy();
            }
            currentGroup = null;
            mouseIsInGroup = false;
            mouseIsOnCircle = false;
          }
          if (echoGroups) {
            for (var j = 0; j < echoGroups.length; j++) {
              var group = echoGroups[j];
              group.x = group.x / zoomRatio;
              group.y = group.y / zoomRatio;
              group.w = group.w / zoomRatio;
              group.h = group.h / zoomRatio;
              console.log("each group:", group);
              if(group.type === "people"){ 
                  chance = 'blue';
                }else if(group.type === "bike"){
                  chance = 'green';
                }else{
                  chance = 'red';
                }
              var rect = new Konva.Rect({
                width: group.w,
                height: group.h,
                stroke: chance,//next or precious img show rect'color
                strokeWidth: 1,
                type: group.type,   
                user: group.user
              });
              rect.on('mouseover', function () {
                document.body.style.cursor = 'default';
              });
              rect.on('mouseout', function () {
                document.body.style.cursor = 'default';
              });

              var minX = drawCanvas.getBoundingClientRect().left;
              var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth - rect.width();
              var minY = drawCanvas.getBoundingClientRect().top;
              var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight - rect.height();
              var rectGroup = new Konva.Group({
                x: group.x,
                y: group.y,
                draggable: false,
                dragBoundFunc: function (pos) {
                  var X = pos.x;
                  var Y = pos.y;
                  if (X < minX) {
                    X = minX;
                  }
                  if (X > maxX) {
                    X = maxX;
                  }
                  if (Y < minY) {
                    Y = minY;
                  }
                  if (Y > maxY) {
                    Y = maxY;
                  }
                  return ({x: X, y: Y});
                }
              });
              rectGroup.on('dragmove', function () {
//                console.log("x:" + this.getX(), "y:" + this.getY());
                console.log("originX:", Math.round(this.getX() * zoomRatio), "originY:", Math.round(this.getY() * zoomRatio));
              });
              rectGroup.on('mouseover', function () {
                mouseIsInGroup = true;
              });
              rectGroup.on('mouseout', function () {
                mouseIsInGroup = false;
              });
              rectGroup.on('mousedown', function (e) {
                if (e.evt.button == 2) {
                  if (currentGroup == this) {
                    currentGroup = null;
                    this.destroy();
                    layer.draw();
                    mouseIsInGroup = false;
                    mouseIsOnCircle = false;
                  }
                }
              });
              rect.on("click", function (e) {
                if (e.evt.button == 0) {
                  var w1 = startX <= endX ? endX - startX : startX - endX;
                  var h1 = startY <= endY ? endY - startY : startY - endY;
                  if (w1 < minRectSize && h1 < minRectSize) {
                    for (var i = 0; i < layer.get('Group').length; i++) {
                      var group = layer.get('Group')[i];
                      var rect = group.get('Rect')[0];
                      if(rect.attrs.type === "people"){ 
                        rect.setStroke('blue');
                      }else if(rect.attrs.type === "bike"){
                        rect.setStroke('green');
                      }else{
                        rect.setStroke('red');
                      }
                      rect.off("mouseover");
                      rect.off("mouseout");
                      group.draggable(false);
                      group.off("mouseover");
                      group.off("mouseout");
                      group.on('mouseover', function () {
                        mouseIsInGroup = true;
                      });
                      group.on('mouseout', function () {
                        mouseIsInGroup = false;
                      });
                    }
                    this.stroke('yellow');
                    this.on('mouseover', function () {
                      document.body.style.cursor = 'move';
                    });
                    this.on('mouseout', function () {
                      document.body.style.cursor = 'default';
                    });
                    currentGroup = this.getParent();
                    currentGroup.moveToTop();
                    document.body.style.cursor = 'move';
                    mouseIsOnCurrentGroup = true;
                    currentGroup.draggable(true);
                    stage.removeEventListener("mousedown", mouseDown, false);
                    currentGroup.on('mouseover', function () {
                      console.log("Mouse is on current group");
                      mouseIsOnCurrentGroup = true;
                      this.draggable(true);
                      stage.removeEventListener("mousedown", mouseDown, false);
                    });
                    currentGroup.on('mouseout', function () {
                      console.log("Mouse is out of current group");
                      mouseIsOnCurrentGroup = false;
                      this.draggable(false);
                      stage.addEventListener("mousedown", mouseDown, false);
                    });
                    layer.draw();
                  }
                }
              });
              rectGroup.add(rect);
              layer.add(rectGroup);
              addAnchor(rectGroup, 0, 0, 'topLeft');
              addAnchor(rectGroup, group.w, 0, 'topRight');
              addAnchor(rectGroup, 0, group.h, 'bottomLeft');
              addAnchor(rectGroup, group.w, group.h, 'bottomRight');
              context.clearRect(0, 0, drawCanvas.width, drawCanvas.height);
              layer.draw();
            }
          }
          layer.draw();
        };
      }
    }
  };

  function mouseDown(eve) {
    if (!mouseIsOnCircle) {
      mouseIsDown = true;
      var pos = getMousePos(drawCanvas, eve);
      startX = endX = pos.x;
      startY = endY = pos.y;
      drawSquare(eve); //update
    }
  }

  function mouseXY(eve) {
    if (mouseIsDown && !mouseIsOnCircle) {
      var pos = getMousePos(drawCanvas, eve);
      endX = pos.x;
      endY = pos.y;
      drawSquare(eve);
    }
  }

  function mouseUp(eve) {
    var pos = getMousePos(drawCanvas, eve);
    endX = pos.x;
    endY = pos.y;
    var w = startX <= endX ? endX - startX : startX - endX;
    var h = startY <= endY ? endY - startY : startY - endY;
    if (mouseIsDown) {
      mouseIsDown = false;
      drawSquare(eve); //update on mouse-up
      if (w >= minRectSize || h >= minRectSize) {
        switch(setType){
          case "people": chance = 'blue';
          break;
          case "bike": chance = 'green';
          break;
          default: chance = 'red';
        }
        var rect = new Konva.Rect({
          width: w,
          height: h, 
          stroke: chance,
          strokeWidth: 1,
          type: setType,
          user: setUser
        });
        rect.on('mouseover', function () {
          document.body.style.cursor = 'default';
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
          draggable: false,
          dragBoundFunc: function (pos) {
            var X = pos.x;
            var Y = pos.y;
            if (X < minX) {
              X = minX;
            }
            if (X > maxX) {
              X = maxX;
            }
            if (Y < minY) {
              Y = minY;
            }
            if (Y > maxY) {
              Y = maxY;
            }
            return ({x: X, y: Y});
          }
        });
        rectGroup.on('dragmove', function () {
//          console.log("x:" + this.getX(), "y:" + this.getY());
          console.log("originX:", Math.round(this.getX() * zoomRatio), "originY:", Math.round(this.getY() * zoomRatio));
        });
        rectGroup.on('mouseover', function () {
          mouseIsInGroup = true;
        });
        rectGroup.on('mouseout', function () {
          mouseIsInGroup = false;
        });
        rectGroup.on('mousedown', function (e) {
          if (e.evt.button == 2) {
            if (currentGroup == this) {
              currentGroup = null;
              this.destroy();
              layer.draw();
              mouseIsInGroup = false;
              mouseIsOnCircle = false;
            }
          }
        });
        rect.on("click", function (e) {
          if (e.evt.button == 0) {
            var w1 = startX <= endX ? endX - startX : startX - endX;
            var h1 = startY <= endY ? endY - startY : startY - endY;
            if (w1 < minRectSize && h1 < minRectSize) {
              for (var i = 0; i < layer.get('Group').length; i++) {
                var group = layer.get('Group')[i];
                var rect = group.get('Rect')[0];
                if(rect.attrs.type === "people"){ 
                  rect.setStroke('blue');
                }else if(rect.attrs.type === "bike"){
                  rect.setStroke('green');
                }else{
                  rect.setStroke('red');
                }
                rect.off("mouseover");
                rect.off("mouseout");
                group.draggable(false);
                group.off("mouseover");
                group.off("mouseout");
                group.on('mouseover', function () {
                  mouseIsInGroup = true;
                });
                group.on('mouseout', function () {
                  mouseIsInGroup = false;
                });
              }
              this.stroke('yellow');
              this.on('mouseover', function () {
                document.body.style.cursor = 'move';
              });
              this.on('mouseout', function () {
                document.body.style.cursor = 'default';
              });
              currentGroup = this.getParent();
              currentGroup.moveToTop();
              document.body.style.cursor = 'move';
              mouseIsOnCurrentGroup = true;
              currentGroup.draggable(true);
              stage.removeEventListener("mousedown", mouseDown, false);
              currentGroup.on('mouseover', function () {
                mouseIsOnCurrentGroup = true;
                this.draggable(true);
                stage.removeEventListener("mousedown", mouseDown, false);
              });
              currentGroup.on('mouseout', function () {
                mouseIsOnCurrentGroup = false;
                this.draggable(false);
                stage.addEventListener("mousedown", mouseDown, false);
              });
              layer.draw();
            }
          }
        });
        rectGroup.add(rect);
        layer.add(rectGroup);
        addAnchor(rectGroup, 0, 0, 'topLeft');
        addAnchor(rectGroup, w, 0, 'topRight');
        addAnchor(rectGroup, 0, h, 'bottomLeft');
        addAnchor(rectGroup, w, h, 'bottomRight');

        //make drawing group as current group
        for (var i = 0; i < layer.get('Group').length; i++) {
          var group = layer.get('Group')[i];
          var rect = group.get('Rect')[0];
          if(rect.attrs.type === "people"){ 
            rect.setStroke('blue');
          }else if(rect.attrs.type === "bike"){
            rect.setStroke('green');
          }else{
           rect.setStroke('red');
          }
          rect.off("mouseover");
          rect.off("mouseout");
          group.draggable(false);
          group.off("mouseover");
          group.off("mouseout");
          group.on('mouseover', function () {
            mouseIsInGroup = true;
          });
          group.on('mouseout', function () {
            mouseIsInGroup = false;
          });
        }
        rectGroup.get('Rect')[0].setStroke('yellow');
        layer.draw();
        currentGroup = rectGroup;
        currentGroup.moveToTop();
        currentGroup.get('Rect')[0].on('mouseover', function () {
          document.body.style.cursor = 'move';
        });
        currentGroup.get('Rect')[0].on('mouseout', function () {
          document.body.style.cursor = 'default';
        });
        currentGroup.on('mouseover', function () {
          mouseIsOnCurrentGroup = true;
          this.draggable(true);
          stage.removeEventListener("mousedown", mouseDown, false);
        });
        currentGroup.on('mouseout', function () {
          mouseIsOnCurrentGroup = false;
          this.draggable(false);
          stage.addEventListener("mousedown", mouseDown, false);
        });
      }

      context.clearRect(0, 0, drawCanvas.width, drawCanvas.height);
      layer.draw();
    }
  }

  function drawSquare(e) {
    var pos = getMousePos(drawCanvas, e);
    endX = pos.x;
    endY = pos.y;
    var minX = drawCanvas.getBoundingClientRect().left;
    var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth;
    var minY = drawCanvas.getBoundingClientRect().top;
    var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight;
//    console.log("maxX: ", maxX, " maxY: ", maxY);
//    console.log("endX: ", endX, " endY: ", endY);
    if (endX <= minX || endX >= maxX - 2 || endY <= minY || endY >= maxY - 1) {
      mouseUp(e);
      return;
    }

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
    switch(setType){
      case "people": context.strokeStyle = 'blue';
      break;
      case "bike": context.strokeStyle = 'green';
      break;
      default: context.strokeStyle = 'red';
    }
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

//        console.log("anchorX:"+anchorX,"anchorY:"+anchorY);
    switch (activeAnchor.getName()) {
      case 'topLeft':
        group.setX(group.getX() + anchorX);
        group.setY(group.getY() + anchorY);
        topLeft.position({x: 0, y: 0});
        topRight.position({x: topRight.getX() - anchorX, y: 0});
        bottomLeft.position({x: 0, y: bottomLeft.getY() - anchorY});
        bottomRight.position({x: bottomRight.getX() - anchorX, y: bottomRight.getY() - anchorY});
        break;
      case 'topRight':
        group.setY(group.getY() + anchorY);
        topLeft.position({x: 0, y: 0});
        topRight.position({x: anchorX, y: 0});
        bottomLeft.position({x: 0, y: bottomLeft.getY() - anchorY});
        bottomRight.position({x: anchorX, y: bottomRight.getY() - anchorY});
        break;
      case 'bottomLeft':
        group.setX(group.getX() + anchorX);
        topLeft.position({x: 0, y: 0});
        topRight.position({x: topRight.getX() - anchorX, y: 0});
        bottomLeft.position({x: 0, y: anchorY});
        bottomRight.position({x: bottomRight.getX() - anchorX, y: anchorY});
        break;
      case 'bottomRight':
        topLeft.position({x: 0, y: 0});
        topRight.position({x: anchorX, y: 0});
        bottomLeft.position({x: 0, y: anchorY});
        bottomRight.position({x: anchorX, y: anchorY});
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
    console.log("originWidth:", Math.round(width * zoomRatio), "originHeight:", Math.round(height * zoomRatio));
    rect.width(width);
    rect.height(height);
    var minX = drawCanvas.getBoundingClientRect().left;
    var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth - rect.width();
    var minY = drawCanvas.getBoundingClientRect().top;
    var maxY = drawCanvas.getBoundingClientRect().top + imgStageHeight - rect.height();
    group.dragBoundFunc(function (pos) {
      var X = pos.x;
      var Y = pos.y;
      if (X < minX) {
        X = minX;
      }
      if (X > maxX) {
        X = maxX;
      }
      if (Y < minY) {
        Y = minY;
      }
      if (Y > maxY) {
        Y = maxY;
      }
      return ({x: X, y: Y});
    });
  }
  function addAnchor(group, x, y, name) {
    var stage = group.getStage();
    var layer = group.getLayer();
    var minX = drawCanvas.getBoundingClientRect().left;
    var maxX = drawCanvas.getBoundingClientRect().left + imgStageWidth;
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
            if (X < minX) {
              X = minX;
            }
            if (X > p.getX() + r.width()) {
              X = p.getX() + r.width();
            }
            if (Y < minY) {
              Y = minY;
            }
            if (Y > p.getY() + r.height()) {
              Y = p.getY() + r.height();
            }
            break;
          case 'topRight':
            if (X < p.getX()) {
              X = p.getX();
            }
            if (X > maxX) {
              X = maxX;
            }
            if (Y < minY) {
              Y = minY;
            }
            if (Y > p.getY() + r.height()) {
              Y = p.getY() + r.height();
            }
            break;
          case 'bottomLeft':
            if (X < minX) {
              X = minX;
            }
            if (X > p.getX() + r.width()) {
              X = p.getX() + r.width();
            }
            if (Y < p.getY()) {
              Y = p.getY();
            }
            if (Y > maxY) {
              Y = maxY;
            }
            break;
          case 'bottomRight':
            if (X < p.getX()) {
              X = p.getX();
            }
            if (X > maxX) {
              X = maxX;
            }
            if (Y < p.getY()) {
              Y = p.getY();
            }
            if (Y > maxY) {
              Y = maxY;
            }
            break;
        }
        return ({x: X, y: Y});
      }
    });
    anchor.on('dragmove', function () {
      update(this);
      layer.draw();
    });
    anchor.on('mousedown touchstart', function () {
//            group.setDraggable(false);
      this.moveToTop();
    });
    anchor.on('dragend', function () {
//            group.setDraggable(true);
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
      mouseIsOnCircle = true;
      layer.draw();
    });
    anchor.on('mouseout', function () {
      document.body.style.cursor = 'default';
      this.strokeWidth(1);
      mouseIsOnCircle = false;
      layer.draw();
    });
    group.add(anchor);
  }
  addEventListener('keydown', check, true);
  document.oncontextmenu = function () {
    return false;
  };
  function check(e) {
    var key = e.which || e.keyCode;
    if (currentGroup) {
      var rect = currentGroup.get('Rect')[0];
      var rectRatio = rect.width() / rect.height();
      var wOffset = rectRatio > 1 ? rectRatio : 1;
      var hOffset = rectRatio > 1 ? 1 : 1 / rectRatio;
      console.log("wOffset:", wOffset, " hOffset:", hOffset);
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
        case 87://w
          if ((minX + 1) <= currentGroup.getX()) {
            topLeft.setX(topLeft.getX() - 1);
            bottomLeft.setX(bottomLeft.getX() - 1);
            update(topLeft);
            layer.draw();
          }
          if ((minY + 1) <= currentGroup.getY()) {
            topLeft.setY(topLeft.getY() - 1);
            topRight.setY(topRight.getY() - 1);
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
        case 83://s
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
        case 88://x
          currentGroup.destroy();
          currentGroup = null;
          layer.draw();
          mouseIsInGroup = false;
          mouseIsOnCircle = false;
          stage.addEventListener("mousedown", mouseDown, false);
          break;
      }
      if (currentGroup) {
        console.log("originX:", Math.round(currentGroup.getX() * zoomRatio), "originY:", Math.round(currentGroup.getY() * zoomRatio));
      }
      layer.draw();
    }
    //下一张
    if (key == 68) {//d
      var tags = [];
      if (layer.get('Group').length > 0) {
        for (var i = 0; i < layer.get('Group').length; i++) {
          var group = layer.get('Group')[i];
          var rect = group.get('Rect')[0];
          tags.push({
            "x": Math.round(group.getX() * zoomRatio),
            "y": Math.round(group.getY() * zoomRatio),
            "w": Math.round(rect.width() * zoomRatio),
            "h": Math.round(rect.height() * zoomRatio),
            "type": rect.attrs.type,  //commit img type
            "user": setUser
          });
        }
        console.log("tags:", tags);
      }
      Vue.http.post("/image/" + currentImageId + "/_" + opType, {"objects": tags}).then(res => {
        console.log("after tag: ", res.body);
        let obj = res.body;
        if (obj.status == "ok") {
        }
        if (imageIds[gIndex + 1]) {
          Vue.http.get("/image/" + imageIds[gIndex + 1]).then(res => {
            let o = res.body;
            echoGroups = o.objects;
            currentImageId = o.id;
            gIndex++;
            imageObj.src = "/f/" + o.path;
            Bus.$emit('updateImgName', o.path);
            Bus.$emit('updateUser', o.demname);
            console.log("/image/:id : ", res.body);
          }, err => {
            console.log("error /image/:id : ", err);
          });
        } else {
          Vue.http.get("/images/_next", {params: {"type": opType}}).then(resp => {
            console.log(resp.body);
            let o = resp.body;
            echoGroups = o.objects;
            currentImageId = o.id;
            Bus.$emit('updateCounts', imageIds.length);
            imageIds.push(currentImageId);
            gIndex = imageIds.length - 1;
            console.log(imageIds);
            imageObj.src = "/f/" + o.path;
            Bus.$emit('updateImgName', o.path);
            Bus.$emit('updateUser', o.demname);
          }, error => {
            Bus.$emit('updateCounts', imageIds.length);
            if (error.ok == false) {
              switch (error.status) {
                case 404:
                  iView.Message.warning("没有更多图片了！");
                  break;
                default:
                  iView.$Message.error("服务器好像出了点问题！");
                  break;
              }
            }
            console.log("image error: ", error);
          });
        }
      }, err => {
        console.log("tag error: ", err);
      });
    }
    if (key == 65) {//a
      var tags = [];
      if (layer.get('Group').length > 0) {
        for (var i = 0; i < layer.get('Group').length; i++) {
          var group = layer.get('Group')[i];
          var rect = group.get('Rect')[0];
          tags.push({
            "x": Math.round(group.getX() * zoomRatio),
            "y": Math.round(group.getY() * zoomRatio),
            "w": Math.round(rect.width() * zoomRatio),
            "h": Math.round(rect.height() * zoomRatio),
            "type": rect.attrs.type,
            "user": setUser
          });
        }
        console.log("tags:", tags);
      }
      Vue.http.post("/image/" + currentImageId + "/_" + opType, {"objects": tags}).then(res => {
        console.log("after tag: ", res.body);
        let obj = res.body;
        if (obj.status == "ok") {
        }
        if (imageIds[gIndex - 1]) {
          Vue.http.get("/image/" + imageIds[gIndex - 1]).then(res => {
            let o = res.body;
            echoGroups = o.objects;
            currentImageId = o.id;
            gIndex--;
            imageObj.src = "/f/" + o.path;
            Bus.$emit('updateImgName', o.path);
            Bus.$emit('updateUser', o.demname);
            console.log("/image/:id : ", res.body);
          }, err => {
            console.log("error /image/:id : ", err);
          });
        } else {
          iView.Message.warning("已经是第一张图片了！");
        }
      }, err => {
        console.log("tag error: ", err);
      });
    }
    if (key == 81) {//q
      if (layer.get('Group').length > 0) {
        if (currentGroup) {
          for (let i = 0; i < layer.get('Group').length; i++) {
            var group = layer.get('Group')[i];
            var rect = group.get('Rect')[0];
            if(rect.attrs.type === "people"){ 
                rect.setStroke('blue');
              }else if(rect.attrs.type === "bike"){
                rect.setStroke('green');
              }else{
                rect.setStroke('red');
            }
            rect.off("mouseover");
            rect.off("mouseout");
            group.draggable(false);
            group.off("mouseover");
            group.off("mouseout");
            group.on('mouseover', function () {
              mouseIsInGroup = true;
            });
            group.on('mouseout', function () {
              mouseIsInGroup = false;
            });
          }
          let j = layer.get('Group').indexOf(currentGroup);
          if (j < layer.get('Group').length - 1) {
            let nextGroup = layer.get('Group')[j + 1];
            nextGroup.get('Rect')[0].stroke('yellow');
            layer.draw();
            currentGroup = nextGroup;
          } else {
            let nextGroup = layer.get('Group')[0];
            nextGroup.get('Rect')[0].stroke('yellow');
            layer.draw();
            currentGroup = nextGroup;
          }
        } else {
          for (let i = 0; i < layer.get('Group').length; i++) {
            var group = layer.get('Group')[i];
            var rect = group.get('Rect')[0];
            if(rect.attrs.type === "people"){ 
                rect.setStroke('blue');
              }else if(rect.attrs.type === "bike"){
                rect.setStroke('green');
              }else{
                rect.setStroke('red');
              }
            rect.off("mouseover");
            rect.off("mouseout");
            group.draggable(false);
            group.off("mouseover");
            group.off("mouseout");
            group.on('mouseover', function () {
              mouseIsInGroup = true;
            });
            group.on('mouseout', function () {
              mouseIsInGroup = false;
            });
          }
          let firstGroup = layer.get('Group')[0];
          firstGroup.get('Rect')[0].stroke('yellow');
          layer.draw();
          currentGroup = firstGroup;
        }

        currentGroup.moveToTop();
        currentGroup.get('Rect')[0].on('mouseover', function () {
          document.body.style.cursor = 'move';
        });
        currentGroup.get('Rect')[0].on('mouseout', function () {
          document.body.style.cursor = 'default';
        });
        currentGroup.on('mouseover', function () {
          mouseIsOnCurrentGroup = true;
          this.draggable(true);
          stage.removeEventListener("mousedown", mouseDown, false);
        });
        currentGroup.on('mouseout', function () {
          mouseIsOnCurrentGroup = false;
          this.draggable(false);
          stage.addEventListener("mousedown", mouseDown, false);
        });
        layer.draw();

      }

    }
//                alert(e.keyCode);
  }
  export default demarcate;
  Array.prototype.max = function () {
    return Math.max.apply({}, this)
  };
  Array.prototype.min = function () {
    return Math.min.apply({}, this)
  };
</script>