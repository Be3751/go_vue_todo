import Vue from 'vue'
import VueRouter from 'vue-router'
import VueCookies from 'vue-cookies'

import Signup from '@/views/SignUp.vue'
import Login from '@/views/LogIn.vue'
import Logout from '@/views/LogOut.vue'
import TaskList from '@/views/TaskList.vue'
import CreateTask from '@/views/CreateTask.vue'
import ReadTask from '@/views/ReadTask.vue'
import UpdateTask from '@/views/UpdateTask.vue'
import DeleteTask from '@/views/DeleteTask.vue'

Vue.use(VueRouter)
Vue.use(VueCookies)

const routes = []

// セッションクッキーの有無でアクセスを制御
const sessionCookie = Vue.$cookies.get('mysession');
if(sessionCookie != null) {
  console.log('Do you want to have a cookie?'); 
  routes.push(
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
    },
    {
      path: '/logout',
      name: 'logout',
      component: Logout
    }
  );
} else {
  console.log('No cookie!');
  routes.push({
    path: '/',
    name: 'signup',
    component: Signup
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  }
  );
}

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
