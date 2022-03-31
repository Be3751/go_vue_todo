<template>
  <v-app>
    <NavBar/>
    <v-main>
      <router-view/>
    </v-main>
    <PageFooter/>
  </v-app>
</template>

<script>
import NavBar from '@/components/NavBar.vue';
import PageFooter from '@/components/PageFooter.vue'

export default {
  name: 'App',
  components: {
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
