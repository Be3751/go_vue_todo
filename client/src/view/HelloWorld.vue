<template>
  <v-container>
    <div v-for="(task, i) in this.tasks" :key="i">
      {{ task }}
    </div>
  </v-container>
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data: () => ({
    tasks: []
  }),
  async mounted() {
    const response = await axios.get('http://localhost:3000/list');
    this.tasks = response.data;
    console.log(response.data);
    console.log(this.tasks);
  }
}
</script>
