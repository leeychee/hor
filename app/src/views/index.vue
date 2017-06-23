<style scoped>
    .layout {
        width: 100%;
        position: absolute;
        border: 1px solid #d7dde4;
        background: #f5f7f9;
        /*position: relative;*/
        height: 100%;
        overflow: hidden;
    }

    .layout-breadcrumb {
        padding: 10px 15px 0;
    }

    .layout-content {
        min-height: 200px;
        overflow: hidden;
        background: #eeeeee;
        height: 100%;
    }

    .layout-content-main {
        padding: 10px;
    }

    .layout-copy {
        text-align: center;
        padding: 10px 0 20px;
        color: #9ea7b4;
    }

    .layout-menu-left {
        background: #464c5b;
    }

    .layout-header {
        height: 60px;
        background: #fff;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .1);
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

    .layout-ceiling-main a {
        color: #9ba7b5;
    }

    .layout-text {
        font-size: 1.4em;
    }

    .layout-hide-text .layout-text {
        display: none;
    }

    .ivu-col {
        transition: width .2s ease-in-out;
    }

    .layout .ivu-row-flex {
        height: 100%;
    }

    /*.ivu-table td{*/
    /*text-align: center;*/
    /*width: 10px;*/
    /*}*/
    .ivu-table td.table-info-column {
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

</style>
<template>
    <div class="layout" :class="{'layout-hide-text': spanLeft < 4}">
        <Row type="flex">
            <i-col :span="spanLeft" class="layout-menu-left">
                <Menu active-name="1" theme="dark" width="auto" @on-select="onSelect">
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
                    <h1 style="display: inline;position: absolute;margin-top: 6px;">{{title}}</h1>

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

                    <router-view></router-view>
                    <!--<div class="resize-container">-->
                    <!--<div class="resize-drag">-->
                    <!--Resize from any edge or corner-->
                    <!--</div>-->
                    <!--</div>-->
                </div>

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
    export default {
        data () {
            return {
                title: "HOR",
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
            },
            onSelect (d){
                console.log(d);
                switch (d) {
                    case "1":
                        this.$router.replace('/demarcate');
                        this.title = "Demarcate";
                        break;
                    case "2":
                        this.$router.replace('/review');
                        this.title = "Review";
                        break;
                    case "3":
                        this.$router.replace('/export');
                        this.title = "Export";
                        break;
                }
            }
        },
        mounted: function () {
            this.$router.replace('/demarcate');
            this.title = "Demarcate";
        }
    };

</script>