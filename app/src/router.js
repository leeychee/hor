const routers = [{
  path: '/',
  meta: {
    title: ''
  },
  component: (resolve) => require(['./views/index.vue'], resolve),
  children: [
    {
      name: 'demarcate',
      path: 'demarcate/:type',
      component: (resolve) => require(['./views/demarcate.vue'], resolve),
      props: true
    },
    {
      name: 'review',
      path: 'review',
      component: (resolve) => require(['./views/review.vue'], resolve)
    },
    {
      name: 'export',
      path: 'export',
      component: (resolve) => require(['./views/export.vue'], resolve)
    }
  ]
}, {
  path: '',
  redirect: { name: 'demarcate', params: { type: 'd' } }
}];
export default routers;