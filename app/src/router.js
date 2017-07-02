import index from './views/index'

const routers = [{
    path: '/',
    meta: {
        title: ''
    },
    component: index,
    children: [
        {
            name: 'demarcate',
            path: 'demarcate/:type',
            component: (resolve) => require(['./views/demarcate.vue'], resolve)
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
        },
        {
            path: '',
            component: (resolve) => require(['./views/demarcate.vue'], resolve)
        }
    ]

},{
    path: '*',
    component: index
}];
export default routers;