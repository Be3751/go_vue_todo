<template>
    <v-container>
        <v-row class="mt-2 ml-2">
        <h1>Make your task!</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-form ref="create_form">
                <v-container>
                    <v-row>
                        <v-col cols="12" md="12">
                            <v-text-field v-model="content" :rules="[required, limit_length]" label="something to do" counter=50></v-text-field>
                            <v-date-picker
                                v-model="date"
                                :allowed-dates="allowedDates"
                                class="mt-4"
                                min="2022-04-01"
                                max="2024-03-31"
                            ></v-date-picker>
                        </v-col>
                        <v-btn class="mr-4" v-on:click="submit">submit</v-btn>
                        <v-btn v-on:click="clear">clear</v-btn>
                    </v-row>
                </v-container>
            </v-form>
        </v-row>
        <v-row class="mt-7 ml-2">
            <div>
                <h2 v-if="errored">Sorry. Couldn't create the task!</h2>
            </div>
            <div>
                <h2 v-if="succeeded">Successful to create the task!</h2>
            </div>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios"

export default {
    name: "CreateTask",
    data: () => ({
        content: "",
        required: value => !!value || "Must include any letter!",
        limit_length: value => value.length <= 50 || "Must include 50 letters or less!",
        info: null,
        errored: false,
        succeeded: false,
        date: '2022-04-01',
    }),
    methods: {
        submit() {
            if(this.$refs.create_form.validate()) {
                this.createTask(this.content, this.date);
                this.$refs.create_form.reset();
            }
        },
        clear() {
            this.$refs.create_form.reset();
        },
        createTask(content, date) {
            axios.post("http://localhost:3000/v1/auth/tasks", {content: content, deadline: date}, {withCredentials: true})
            .then(response => {
                this.info = response.bpi;
                this.succeeded = true;
                console.log(date);
            })
            .catch(error => {
                console.log(date);
                console.log(error);
                this.errored = true;
                this.succeeded = false;
            });
        },
        allowedDates: val => parseInt(val.split('-')[2], 10) % 1 === 0,
    }
}
</script>