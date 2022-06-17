<template>
    <v-container>
        <v-row class="mt-2 ml-2">
        <h1>Here is your task!</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-col cols="12" md="12">
                Task : {{ this.content }}<br/>
                Deadline : {{ this.deadline }}
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios"

export default {
    name: "readTask",
    data: () => ({
        id: "",
        content: "",
        deadline: "",
        required: value => !!value || "Must include any letter!",
        limit_length: value => value.length <= 50 || "Must include 50 letters or less!",
        info: null,
        errored: false,
        succeeded: false
    }),
    methods: {
        readTask(id) {
            axios.get("http://localhost:3000/v1/auth/tasks/"+id, {withCredentials: true})
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
        this.deadline = this.$route.params.deadline;
    }
}
</script>