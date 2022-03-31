<template>
  <v-app-bar absolute color="black" dark app>
    <v-toolbar-title>Task Management!</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn to="/" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      View task list
    </v-btn>
    <v-btn to="/create" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-pencil</v-icon>
      Create task
    </v-btn>
  </v-app-bar>
</template>

<script>
  export default {
    data: () => ({

    }),
  }
</script>