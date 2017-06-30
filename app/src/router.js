const routers = [{
    path: '/',
    meta: {
        title: ''
    },
    component: (resolve) => require(['./views/index.vue'], resolve),
    children: [
        {
            path: 'demarcate',
            component: (resolve) => require(['./views/demarcate.vue'], resolve)
        },
        {
            path: 'review',
            component: (resolve) => require(['./views/review.vue'], resolve)
        },
        {
            path: 'export',
            component: (resolve) => require(['./views/export.vue'], resolve)
        },
        {
            path: '',
            component: (resolve) => require(['./views/demarcate.vue'], resolve)
        }
    ]

},{
    path: '*',
    component: (resolve) => require(['./views/index.vue'], resolve)
}];
export default routers;