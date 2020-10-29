import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Dashboard from '../views/Dashboard.vue'
import Cases from '../views/Cases.vue'
import Users from '../views/Users.vue'
import UserAdd from '../views/UserAdd.vue'
import CaseAdd from '../views/CaseAdd.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    alias:['/admin','/home'],
    name: 'Home',
    component: Home,
    children:[
      { path: '', alias:['dashboard'], component: Dashboard },
      { path: 'users', component: Users },
      { path: 'cases', component: Cases },
      { path: 'users/add', component: UserAdd },
      { path: 'cases/add', component: CaseAdd },
    ]
  },
  {
    path: '/signin',
    alias:['/login'],
    name: 'Login',
    component: Login
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
