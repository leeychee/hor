<style scoped>
</style>
<template>
  <div style="margin-top: 20px;">
    <Row>
      <Col v-if="showError" span="4" offset="10" >
      <Alert type="error" closable>无法获取统计信息</Alert>
      </Col>
    </Row>
    <div style="width: 40%;margin: 0 auto;">
      <Table style="width: 100%;height: 100%;border-radius: 5px;" :columns="columns0" :data="data0">
      </Table>
      <Table style="width: 100%;height: 100%;border-radius: 5px;" :columns="columns1" :data="data1">
      </Table>
    </div>
    <div style="text-align: center;">
      <Button type="info" size="large" style="margin: 20px;">导出xml</Button>
    </div>
  </div>
</template>
<script>
  export default {
    data() {
      return {
        showError: false,
        all: 0,
        identified: 0,
        reviewed: 0,
        columns0: [
          {
            title: 'Statistics',
            key: 'signType'
          },
          {
            title: 'Num',
            key: 'num',
            align: 'right'
          }
        ],
        data0: [],
        columns1: [
          {
            title: 'User',
            key: 'user'   
          },
          {
            title: 'Identified',
            key: 'Identified1',
            align: 'center'
          },
          {
            title: 'Reviewed',
            key: 'Reviewed1',
            align: 'right'
          }
        ],
        data1: []
      }
    },
    mounted() {
      this.$http.get("/images/_stat").then(res => {
        console.log(res);
        let o = res.body;
        this.data0 = [
          {
            signType: 'All',
            num: o.all
          },
          {
            signType: 'Identified',
            num: o.identified
          },
          {
            signType: 'Reviewed',
            num: o.Reviewed
          }
        ];
        this.data1 = [
          {
            user: o.users[0].username,
            Identified1: o.users[0].demcount,
            Reviewed1: o.users[0].revcount
          },
          {
            user: o.users[1].username,
            Identified1: o.users[1].demcount,
            Reviewed1: o.users[1].revcount
          },
          {
            user: o.users[2].username,
            Identified1: o.users[2].demcount,
            Reviewed1: o.users[2].revcount
          }
        ];
      }, err => {
        this.showError = true;
      });
    },
    beforeDestroy() {
    },
    methods: {}
  }
</script>
