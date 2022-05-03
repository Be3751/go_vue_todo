<template>
    <v-container>
        <v-row class="mt-2 ml-2">
        <h1>Are you sure to delete this task?</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-form ref="delete_form">
                <v-container>
                    <v-row>
                        <v-col cols="12" md="12">
                           Task : {{ this.content }}
                        </v-col>
                        <v-btn class="mr-4" v-on:click="submit">submit</v-btn>
                        <v-btn v-on:click="clear">clear</v-btn>
                    </v-row>
                </v-container>
            </v-form>
        </v-row>
        <v-row class="mt-7 ml-2">
            <div>
                <h2 v-if="errored">Sorry. Couldn't delete it!</h2>
            </div>
            <div>
                <h2 v-if="succeeded">Successful to delete it!</h2>
            </div>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios"

export default {
    name: "deleteTask",
    data: () => ({
        id: "",
        content: "",
        required: value => !!value || "Must include any letter!",
        limit_length: value => value.length <= 50 || "Must include 50 letters or less!",
        info: null,
        errored: false,
        succeeded: false
    }),
    methods: {
        submit() {
            if(this.$refs.delete_form.validate()) {
                this.deleteTask(this.id);
                this.$refs.delete_form.reset();
            }
        },
        clear() {
            this.$refs.delete_form.reset();
        },
        deleteTask(id) {
            axios.delete("http://localhost:3000/auth/tasks/"+id, {withCredentials: true})
            .then(response => {
                this.info = response.bpi;
                this.succeeded = true;
            })
            .catch(error => {
                console.log(error)
                this.errored = true;
                this.succeeded = false;
            });
        }
    },
    mounted() {
        this.id = this.$route.params.id;
        this.content = this.$route.params.content;
    }
}
</script>