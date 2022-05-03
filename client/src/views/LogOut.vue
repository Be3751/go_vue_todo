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
export default {
    name: "LogOut",
    data: () => ({
        errored: false,
        succeeded: false
    }),
    methods: {
        logout() {
            const sessionKeyName = "mysession"
            if(this.$cookies.isKey(sessionKeyName)) {
                this.$cookies.remove(sessionKeyName);
                console.log("removed session cookie.");
                this.succeeded = true;
            } else {
                console.log("no cookie.");
                this.errored = true;
            }
        }
    }
}
</script>

<style>

</style>