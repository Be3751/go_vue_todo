<template>
  <v-app>
    <NavBar/>
    <v-main>
      <HelloWorld/>
    </v-main>
    <PageFooter/>
  </v-app>
</template>

<script>
import HelloWorld from './view/HelloWorld';
import NavBar from './components/NavBar.vue';
import PageFooter from './components/PageFooter.vue'

export default {
  name: 'App',

  components: {
    HelloWorld,
    NavBar,
    PageFooter
  },

  data: () => ({
    //
  }),
};
</script>
