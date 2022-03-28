<template>
  <v-app>
    <NavBar/>
    <v-main>
      <TaskList/>
    </v-main>
    <PageFooter/>
  </v-app>
</template>

<script>
import TaskList from './view/TaskList.vue'
import NavBar from './components/NavBar.vue';
import PageFooter from './components/PageFooter.vue'

export default {
  name: 'App',
  components: {
    TaskList,
    NavBar,
    PageFooter
  },
  data: () => ({
    title: "",
  }),
  methods: {
  }
};
</script>
