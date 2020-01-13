import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
    linkExactActiveClass: 'active',
    routes: [
      {
        /* path: '/',
        redirect: '',      
        children: [
          { */
            path: '/',
            name: 'search',
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            component: () => import(/* webpackChunkName: "search" */ './views/Search.vue')
          },
          {
            path: '/history',
            name: 'history',
            component: () => import(/* webpackChunkName: "history" */ './views/History.vue')
          }    
          ]
      //},    
    //] 
  })
  