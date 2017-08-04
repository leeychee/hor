<style scoped>
  .layout {
    width: 100%;
    position: absolute;
    border: 1px solid #d7dde4;
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
    background: #f7f7f7;
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
    height: 40px;
    background: #e4e5e7;
    box-shadow: 0 1px 1px rgba(0, 0, 0, .1);
  }

  .layout-logo-left {
    width: 100%;
    height: 40px;
    /*background: #112129;*/
    border-radius: 3px;
    margin: 0 auto 10px;
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

  .operate-tip {
    z-index: 2;
    position: fixed;
    bottom: 20px;
    left: 20px;
  }

</style>
<template>
  <div class="layout" :class="{'layout-hide-text': spanLeft < 4}">
    <Row type="flex">
      <i-col :span="spanLeft" class="layout-menu-left">
        <Menu active-name="1" theme="dark" width="auto" @on-select="onSelect">
          <div class="layout-logo-left">
            <!--<embed :src="logoSVG" width="95" height="30" type="image/svg+xml"  />-->
            <img :src="logoSVG" width="80%" height="30" style="margin: 5% 10%;">
          </div>
          <Menu-item name="1" title="标定">
            <Icon type="ios-navigate" :size="iconSize"></Icon>
            <span class="layout-text">Demarcate</span>
          </Menu-item>
          <Menu-item name="2" title="审阅">
            <Icon type="ios-keypad" :size="iconSize"></Icon>
            <span class="layout-text">Review</span>
          </Menu-item>
          <Menu-item name="3" title="导出">
            <Icon type="ios-analytics" :size="iconSize"></Icon>
            <span class="layout-text">Export</span>
          </Menu-item>
        </Menu>
      </i-col>
      <i-col :span="spanMiddle">
        <div v-if="showHeader" class="layout-header">
          <h1 style="display: inline;margin-left: 10px;">{{title}}</h1>
          <Tag v-if="showCurrent" style="margin-left: 30px;" type="dot" color="blue"><b>Current Image: {{imgName}}</b>
          </Tag>
          <!--<h4 style="display: inline;margin-left: 30px;">{{imgName}}</h4>-->
          <Tag v-if="showCurrent" style="margin-left: 30px;" type="dot" color="blue"><b>Current
            Counts：{{currentCounts}}</b>
          </Tag>
          <!--<h4 style="display: inline;margin-left: 30px;">Current Counts：{{currentCounts}}</h4>-->
          <div v-if="showRectSet" style="float: right;margin: 5px;">
            <h3 style="display: inline;">Min Rect Size(px)：</h3>
            <Select v-model="minSize" size="small" style="width:80px;" @on-change="changeMinSize">
              <Option v-for="item in minList" :value="item.value" :key="item.value">{{ item.label }}</Option>
            </Select>
          </div>
          <div v-if="showSetType" style="float: right;margin: 5px;">
            <Radio-group v-model="preType" size="large" style="width: 220px;" @on-change="changePreType">
              <Radio label="car">
                <Icon value="car" type="social-car" style="color:red"></Icon>
                <span>汽车</span>
              </Radio>
              <Radio label="bike">
                <Icon value="bike" type="social-bike" style="color:green"></Icon>
                <span>自行车</span>
              </Radio>
              <Radio label="people">
                <Icon value="people" type="social-people" style="color:blue"></Icon>
                <span>行人</span>
              </Radio>
            </Radio-group>
          </div>
        </div>
        <div class="layout-content">
          <router-view :minSize.sync="minSize" :preType.sync="preType" :v-on:updateImgName="showImgName" :key="key"></router-view>
        </div>
      </i-col>
    </Row>
    <div class="operate-tip">
      <Poptip trigger="hover" placement="right-end">
        <Icon type="compass" size="24" color="#aaa"></Icon>
        <div slot="title"><h3>操作快捷键</h3></div>
        <div slot="content">
          <Table :columns="columnsTip" :data="dataTip"></Table>
        </div>
      </Poptip>
    </div>
  </div>
</template>
<script>
  import logoUrl from '../static/img/hor.svg'
  import demarcate from './demarcate.vue'
  import review from './review.vue'
  import Bus from '../libs/bus';
  export default {
    data () {
      return {
        title: "HOR",
        logoSVG: logoUrl,
        showHeader: true,
        spanLeft: 2,
        spanMiddle: 22,
        spanRight: 6,
        minList: [
          {
            value: '10',
            label: '10'
          },
          {
            value: '20',
            label: '20'
          },
          {
            value: '30',
            label: '30'
          },
          {
            value: '40',
            label: '40'
          },
          {
            value: '50',
            label: '50'
          },
          {
            value: '60',
            label: '60'
          }
        ],
        preType: '',
        minSize: '60',
        preType: 'car',
        showRectSet: true,
        showSetType: true,
        showCurrent: true,
        columnsTip: [
          {
            title: '按键',
            key: 'key',
            width: 60
          },
          {
            title: '操作',
            key: 'value',
            width: 110

          }
        ],
        dataTip: [
          {
            key: 'a',
            value: '上一张'
          },
          {
            key: 'd',
            value: '下一张'
          },
          {
            key: 'w',
            value: '放大'
          },
          {
            key: 's',
            value: '缩小'
          },
          {
            key: 'q',
            value: '切换选框'
          },
          {
            key: 'x',
            value: '删除选中选框'
          },
          {
            key: '↑→←↓',
            value: '移动选中选框'
          }
        ],
        imgName: '',
        currentCounts: 0
      };
    },
    computed: {
      iconSize () {
        return this.spanLeft === 4 ? 24 : 32;
      },
      key() {
        return this.$route.name !== undefined ? this.$route.name + new Date() : this.$route + new Date()
      }
    },
    created: function () {
      Bus.$on("updateImgName", imgName => {
        this.imgName = imgName;
      });
      Bus.$on("updateCounts", count => {
        this.currentCounts = count;
      });
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
        switch (d) {
          case "1":
//                        this.$router.replace('demarcate');
            this.$router.replace({name: 'demarcate', params: {type: 'd'}});
            this.title = "Demarcate";
            this.imgName = "";
            this.currentCounts = 0;
            this.showRectSet = true;
            this.showSetType = true;
            this.showCurrent = true;
            break;
          case "2":
//                        this.$router.replace('review');
            this.$router.replace({name: 'demarcate', params: {type: 'r'}});
            this.title = "Review";
            this.imgName = "";
            this.currentCounts = 0;
            this.showRectSet = true;
            this.showSetType = true;
            this.showCurrent = true;
            break;
          case "3":
            this.$router.replace({name: 'export'});
            this.title = "Export";
            this.showRectSet = false;
            this.showSetType = false;
            this.showCurrent = false;
            break;
        }
      },
      changeMinSize (val) {
      },
      changePreType (val){
      },
      showImgName: function (imgName) {
        this.imgName = imgName;
      }
    },
    mounted: function () {
//            this.$router.replace('demarcate');
      this.$router.replace({name: 'demarcate', params: {type: 'd'}});
      this.title = "Demarcate";
      this.showRectSet = true;
      this.showSetType = true;
    }
  };
</script>
