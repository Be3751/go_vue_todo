<template>
  <v-app-bar absolute color="black" dark app>
    <v-toolbar-title>Task Management!</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn v-if="sessioned" to="/" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      View task list
    </v-btn>
    <v-btn v-else to="/" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-clipboard-account-outline</v-icon>
      Signup
    </v-btn>
    <v-btn v-if="sessioned" to="/create" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-pencil</v-icon>
      Create task
    </v-btn>
    <v-btn v-if="sessioned" to="/logout" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-logout-variant</v-icon>
      Logout
    </v-btn>
    <v-btn v-else to="/login" class="ma-2" color="teal accent-3">
      <v-icon dark>mdi-login-variant</v-icon>
      Login
    </v-btn>
  </v-app-bar>
</template>

<script>
  export default {
    data: () => ({
      sessioned: false
    }),
    mounted() {
      const userStatus = this.$cookies.get('user-status');
      if(userStatus == 'auth') {
        this.sessioned = true
      }
    }
  }
</script>