<template>
    <v-container>
        <v-row class="mt-2 ml-2 mb-5">
        <h1>Are you sure to log out?</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-btn v-on:click="cancel">Cancel</v-btn>
            <v-btn class="ml-3 mr-4" v-on:click="logout">Log out</v-btn>
        </v-row>
        <v-row class="mt-7 ml-2">
            <div>
                <h2 v-if="errored">Sorry. Couldn't log out!</h2>
            </div>
            <div>
                <h2 v-if="succeeded">Successful to log out!</h2>
            </div>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios"
export default {
    name: "LogOut",
    data: () => ({
        errored: false,
        succeeded: false
    }),
    methods: {
        logout() {
            axios.get("http://localhost:3000/auth/logout", {withCredentials: true})
            .then(response => {
                if(response.status == 200) {
                    this.$cookies.set('user-status', null);
                    this.$router.push({name: "signup"});
                    this.$router.go({path: this.$router.currentRoute.path, force: true}); // 遷移後にリロードを行うことでAPIにリクエスト
                    this.errored = true;
                }
            })
            .catch(error => {
                console.log(error)
                this.errored = true;
                this.succeeded = false;
            });
        },
        cancel() {
            console.log("canceled to jump to list page.");
            this.$router.push({name: "list"});
        }
    }
}
</script>

<style>

</style>