<style scoped>
</style>
<template>
  <div style="margin-top: 20px;">
    <Row>
      <Col v-if="showError" span="4" offset="10" >
      <Alert type="error" closable>无法获取统计信息</Alert>
      </Col>
      <Col span="8" offset="8">
      <Card dis-hover>
        <p slot="title">Statistics</p>
        <span style="float: right;">{{all}}</span><h4>All</h4>
        <span style="float: right;">{{identified}}</span><h4>Identified</h4>
        <span style="float: right;">{{reviewed}}</span><h4>Reviewed</h4>
      </Card>
      </Col>
    </Row>
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
        reviewed: 0
      };
    },
    mounted() {
      this.$http.get("/images/_stat").then(res => {
        console.log(res);
        let o = res.body;
        this.all = o.all;
        this.identified = o.identified;
        this.reviewed = o.Reviewed;
      }, err => {
        this.showError = true;
      });
    },
    beforeDestroy() {

    },
    methods: {}
  };
</script>
