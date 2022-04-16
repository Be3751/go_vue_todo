import Vue from 'vue'
import VueRouter from 'vue-router'
import Signup from '@/views/SignUp.vue'
import Login from '@/views/LogIn.vue'
import Logout from '@/views/LogOut.vue'
import TaskList from '@/views/TaskList.vue'
import CreateTask from '@/views/CreateTask.vue'
import ReadTask from '@/views/ReadTask.vue'
import UpdateTask from '@/views/UpdateTask.vue'
import DeleteTask from '@/views/DeleteTask.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/signup',
    name: 'signup',
    component: Signup
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/logout',
    name: 'logout',
    component: Logout
  },
  {
    path: '/',
    name: 'list',
    component: TaskList
  },
  {
    path: '/create',
    name: 'create',
    component: CreateTask
  },
  {
    path: '/read',
    name: 'read',
    component: ReadTask 
  },
  {
    path: '/update/:id',
    name: 'update',
    component: UpdateTask 
  },
  {
    path: '/delete/:id',
    name: 'delete',
    component: DeleteTask 
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
