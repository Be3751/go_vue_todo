import Vue from 'vue'
import VueRouter from 'vue-router'
import TaskList from '@/views/TaskList.vue'
import CreateTask from '@/views/CreateTask.vue'
import ReadTask from '@/views/ReadTask.vue'
import UpdateTask from '@/views/UpdateTask.vue'
import DeleteTask from '@/views/DeleteTask.vue'

Vue.use(VueRouter)

const routes = [
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
    path: '/update',
    name: 'update',
    component: UpdateTask 
  },
  {
    path: '/delete',
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
