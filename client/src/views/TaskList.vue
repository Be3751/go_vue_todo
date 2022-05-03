<template>
  <v-container>
    <v-row class="mt-2 ml-2">
      <h1>Here are your tasks!</h1>
    </v-row>
    <v-row>
      <!-- <v-col> -->
        <TaskCard :task="task" v-for="(task, i) in this.tasks" :key="i"/>
      <!-- </v-col> -->
    </v-row>
  </v-container>
</template>

<script>
import TaskCard from '@/components/TaskCard.vue'
import axios from 'axios'

export default {
  name: 'TaskList',
  components: {
    TaskCard
  },
  data: () => ({
    tasks: []
  }),
  async mounted() {
    await axios.get('http://localhost:3000/auth/tasks', {withCredentials: true})
    .then(response => {
        this.tasks = response.data;
    })
    .catch(error => {
        console.log(error);
    });
  }
}
</script>
