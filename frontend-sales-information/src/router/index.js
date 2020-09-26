import Vue from 'vue'
import VueRouter from 'vue-router'


Vue.use(VueRouter)

const routes = [{
        path: '/',
        name: 'Home',
        component: () =>
            import ( /* webpackChunkName: "Home" */ '../views/Home.vue')
    },
    {
        path: '/buyers',
        name: 'Buyers',
        component: () =>
            import ( /* webpackChunkName: "Buyers" */ '../views/Buyers.vue')
    },
    {
        path: '/upload',
        name: 'Upload',
        component: () =>
            import ( /* webpackChunkName: "Upload" */ '../views/Upload.vue')
    },
    {
        path: '/buyer/:buyerId',
        name: 'Buyer',
        component: () =>
            import ( /* webpackChunkName: "Buyer" */ '../views/Buyer.vue')
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router