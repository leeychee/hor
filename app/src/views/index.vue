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
                        <img src="http://192.168.99.2:9000/f/P_20170601100228_994.jpg">
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
    import Rect from '../libs/rect';
    var interact = require('interactjs');
    import Konva from 'konva';
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
            Rect.init("main");
            interact('.resize-drag')
                .draggable({
                    // enable inertial throwing
                    inertia: true,
                    // keep the element within the area of it's parent
                    restrict: {
                      restriction: "parent",
                      endOnly: true,
                      elementRect: { top: 0, left: 0, bottom: 1, right: 1 }
                    },
                    // enable autoScroll
                    autoScroll: true,

                    // call this function on every dragmove event
                    onmove: dragMoveListener,
                    // call this function on every dragend event
                    onend: function (event) {
                      var textEl = event.target.querySelector('p');

                      textEl && (textEl.textContent =
                        'moved a distance of '
                        + (Math.sqrt(event.dx * event.dx +
                                     event.dy * event.dy)|0) + 'px');
                    }
                })
                .resizable({
                    preserveAspectRatio: true,
                    edges: { left: true, right: true, bottom: true, top: true }
                })
                .on('resizemove', function (event) {
                    var target = event.target,
                        x = (parseFloat(target.getAttribute('data-x')) || 0),
                        y = (parseFloat(target.getAttribute('data-y')) || 0);

                    // update the element's style
                    target.style.width  = event.rect.width + 'px';
                    target.style.height = event.rect.height + 'px';

                    // translate when resizing from top or left edges
                    x += event.deltaRect.left;
                    y += event.deltaRect.top;

                    target.style.webkitTransform = target.style.transform =
                        'translate(' + x + 'px,' + y + 'px)';

                    target.setAttribute('data-x', x);
                    target.setAttribute('data-y', y);
                    target.textContent = Math.round(event.rect.width) + '×' + Math.round(event.rect.height);
                })
                .on('mouseover', function (event) {
                    Rect.container.onmousedown = null;
                    console.log("mouse come in!");
                })
                .on('mouseout', function (event) {
                    Rect.container.onmousedown = Rect.start;
                    console.log("mouse leave!");
                })
                ;
        }
    };
    function dragMoveListener (event) {
        var target = event.target,
            // keep the dragged position in the data-x/data-y attributes
            x = (parseFloat(target.getAttribute('data-x')) || 0) + event.dx,
            y = (parseFloat(target.getAttribute('data-y')) || 0) + event.dy;

        // translate the element
        target.style.webkitTransform =
        target.style.transform =
          'translate(' + x + 'px, ' + y + 'px)';

        // update the posiion attributes
        target.setAttribute('data-x', x);
        target.setAttribute('data-y', y);
      }

      // this is used later in the resizing and gesture demos
      window.dragMoveListener = dragMoveListener;
</script>