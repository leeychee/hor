<style scoped>
    .layout{
        width: 100%;
        position: absolute;
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        /*position: relative;*/
        height: 100%;
        border-radius: 4px;
        overflow: hidden;
    }
    .layout-breadcrumb{
        padding: 10px 15px 0;
    }
    .layout-content{
        min-height: 200px;
        margin: 15px;
        overflow: hidden;
        background: #dfe5e1;
        border-radius: 4px;
        height: 80%;
    }
    .layout-content-main{
        padding: 10px;
    }
    .layout-copy{
        text-align: center;
        padding: 10px 0 20px;
        color: #9ea7b4;
    }
    .layout-menu-left{
        background: #464c5b;
    }
    .layout-header{
        height: 60px;
        background: #fff;
        box-shadow: 0 1px 1px rgba(0,0,0,.1);
    }
    .layout-logo-left {
        width: 90%;
        height: 30px;
        background: #5b6270;
        border-radius: 3px;
        margin: 15px auto;
        padding-left: 8px;
        padding-top: 6px;
        color: white;
    }
    .layout-ceiling-main a{
        color: #9ba7b5;
    }
    .layout-text {
        font-size: 1.4em;
    }
    .layout-hide-text .layout-text{
        display: none;
    }
    .ivu-col{
        transition: width .2s ease-in-out;
    }
    .layout .ivu-row-flex {
        height: 100%;
    }
    /*.ivu-table td{*/
        /*text-align: center;*/
        /*width: 10px;*/
    /*}*/
    .ivu-table td.table-info-column{
        background-color: #2db7f5;
        color: #fff;
    }
    .right-toggle-btn {
        float: right;
        margin-right: 30%;
    }
    .left-hide-right-toggle-btn {
        float: right;
        margin-right: 27%;
    }

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
    <div class="layout" :class="{'layout-hide-text': spanLeft < 4}">
        <Row type="flex">
            <i-col :span="spanLeft" class="layout-menu-left">
                <Menu active-name="1" theme="dark" width="auto">
                    <div class="layout-logo-left">LOGO—HOR</div>
                    <Menu-item name="1">
                        <Icon type="ios-navigate" :size="iconSize"></Icon>
                        <span class="layout-text">Demarcate</span>
                    </Menu-item>
                    <Menu-item name="2">
                        <Icon type="ios-keypad" :size="iconSize"></Icon>
                        <span class="layout-text">Review</span>
                    </Menu-item>
                    <Menu-item name="3">
                        <Icon type="ios-analytics" :size="iconSize"></Icon>
                        <span class="layout-text">Export</span>
                    </Menu-item>
                </Menu>
            </i-col>
            <i-col :span="spanMiddle">
                <div class="layout-header">
                    <i-button type="text" @click="toggleLeftClick">
                        <Icon type="navicon" size="32"></Icon>
                    </i-button>
                    <h1 style="display: inline;position: absolute;margin-top: 6px;">Demarcate</h1>

                    <!--<i-button type="text" class="right-toggle-btn" :class="{'left-hide-right-toggle-btn': spanLeft < 4}" >-->
                        <!--<Icon type="navicon" size="32"></Icon>-->
                    <!--</i-button>-->
                </div>
                <!--<div class="layout-breadcrumb">-->
                    <!--<Breadcrumb>-->
                        <!--<Breadcrumb-item href="#">首页</Breadcrumb-item>-->
                        <!--<Breadcrumb-item href="#">应用中心</Breadcrumb-item>-->
                        <!--<Breadcrumb-item>某应用</Breadcrumb-item>-->
                    <!--</Breadcrumb>-->
                <!--</div>-->
                <div class="layout-content">
                    <div id="main" style="height: 100%;">
                    </div>
                    <!--<div class="resize-container">-->
                      <!--<div class="resize-drag">-->
                         <!--Resize from any edge or corner-->
                      <!--</div>-->
                    <!--</div>-->
                </div>
                <Row type="flex" justify="center" class="code-row-bg" style="height: 10%;">
                    <Col span="3"><Button type="info" size="large">Submit</Button></Col>
                    <Col span="3"><Button type="warning" size="large">&nbsp;Next&nbsp;</Button></Col>
                </Row>
                <!--<div class="layout-copy">-->
                    <!--2000-2017 &copy; SunCreate.-->
                <!--</div>-->
            </i-col>
            <!--<i-col :span="spanRight" :offset="18" style="position: fixed;z-index: 9999;">-->
                <!--<Table width="auto" height="auto" :columns="columns2" :data="data3"></Table>-->
            <!--</i-col>-->
        </Row>
    </div>
</template>
<script>
//    import Rect from '../libs/rect';
//    var interact = require('interactjs');
    import Konva from 'konva';
    var stage, layer, canvas, context;
    export default {
        data () {
            return {
                spanLeft: 4,
                spanMiddle: 20,
                spanRight: 6,
                columns2: [
                    {
                        title: '序号',
                        key: 'index',
                        width: 60,
                        fixed: 'left',
                        className: 'table-info-column'
                    },
                    {
                        title: 'x',
                        key: 'xAxis',
                        width: 60
                    },
                    {
                        title: 'y',
                        key: 'yAxis',
                        width: 60
                    },
                    {
                        title: 'w',
                        key: 'width',
                        width: 60
                    },
                    {
                        title: 'h',
                        key: 'height',
                        width: 60
                    },
                    {
                        title: '操作',
                        key: 'action',
                        fixed: 'right',
                        width: 120,
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'text',
                                        size: 'small'
                                    }
                                }, '编辑'),
                                h('Button', {
                                    props: {
                                        type: 'text',
                                        size: 'small'
                                    }
                                }, '删除')
                            ]);
                        }
                    }
                ],
                data3: [
                    {
                        index: '1',
                        xAxis: 10,
                        yAxis: 10,
                        width: 50,
                        height: 50
                    },
                    {
                        index: '2',
                        xAxis: 30,
                        yAxis: 10,
                        width: 30,
                        height: 50
                    },
                    {
                        index: '3',
                        xAxis: 15,
                        yAxis: 100,
                        width: 50,
                        height: 50
                    },
                    {
                        index: '4',
                        xAxis: 210,
                        yAxis: 10,
                        width: 20,
                        height: 120
                    }
                ]
            }
        },
        computed: {
            iconSize () {
                return this.spanLeft === 4 ? 24 : 32;
            }
        },
        methods: {
            toggleLeftClick () {
                if (this.spanLeft === 4) {
                    this.spanLeft = 2;
                    this.spanMiddle = 22;
                    this.spanRight = 6;
                } else {
                    this.spanLeft = 4;
                    this.spanMiddle = 20;
                    this.spanRight = 6;
                }
            }
        },
        mounted: function () {
            stage = new Konva.Stage({
                container: 'main',
                width: 5/6*window.innerWidth -30,
                height: .8*window.innerHeight,
            });
            // add canvas element
            layer = new Konva.Layer();
            stage.add(layer);

            var imageObj = new Image();
            imageObj.onload = function () {
                var yoda = new Konva.Image({
                    x: 0,
                    y: 0,
                    image: imageObj,
//                    width: stage.width(),
//                    height: imageObj.height/imageObj.width*stage.width(),
                    width: imageObj.width/imageObj.height*stage.height(),
                    height: stage.height(),
                });
                // add the shape to the layer
                layer.add(yoda);
                layer.add(image);
                layer.draw();
            };
            imageObj.src = 'http://www.bz55.com/uploads/allimg/150306/139-1503061IR6.jpg';

            canvas = document.createElement('canvas');
            canvas.width = stage.width();
            canvas.height = stage.height();

            var image = new Konva.Image({
                image: canvas,
                x: 0,
                y: 0
            });
//            layer.add(image);
//            layer.draw();
//            stage.draw();
            context = canvas.getContext("2d");

            stage.addEventListener("mousedown", mouseDown, false);
            stage.addEventListener("mousemove", mouseXY, false);
            stage.addEventListener("mouseup", mouseUp, false);
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
                stroke: 'white',
                strokeWidth: 2
            });
            rect.on('mouseover', function () {
                document.body.style.cursor = 'move';
            });
            rect.on('mouseout', function () {
                document.body.style.cursor = 'default';
            });

            var rectGroup = new Konva.Group({
                x: startX,
                y: startY,
                draggable: true
            });
            rectGroup.on('dragmove', function () {
                console.log("x:" + this.getX(), "y:" + this.getY());
            });
            rectGroup.on('mouseover', function () {
                mouseIsInGroup = true;
                console.log("mouse over this group");
            });
            rectGroup.on('mouseout', function () {
                mouseIsInGroup = false;
                console.log("mouse out this group!!!");
            });
            layer.add(rectGroup);
            rectGroup.add(rect);
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
        context.lineWidth = 2;
        context.strokeStyle = 'white';
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
        }
    }
    function addAnchor(group, x, y, name) {
        var stage = group.getStage();
        var layer = group.getLayer();
        var anchor = new Konva.Circle({
            x: x,
            y: y,
            stroke: '#666',
            fill: '#ddd',
            strokeWidth: 1,
            radius: 4,
            name: name,
            draggable: true,
            dragOnTop: false
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
                    document.body.style.cursor = 'nw-resize';
                    break;
                case 'topRight':
                    document.body.style.cursor = 'ne-resize';
                    break;
                case 'bottomRight':
                    document.body.style.cursor = 'se-resize';
                    break;
                case 'bottomLeft':
                    document.body.style.cursor = 'sw-resize';
                    break;
            }
            this.setStrokeWidth(4);
            layer.draw();
        });
        anchor.on('mouseout', function () {
            var layer = this.getLayer();
            document.body.style.cursor = 'default';
            this.setStrokeWidth(2);
            layer.draw();
        });
        group.add(anchor);
    }
</script>