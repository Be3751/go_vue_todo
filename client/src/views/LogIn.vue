<template>
    <v-container>
        <v-row class="mt-2 ml-2">
        <h1>Log in!</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-form ref="create_form">
                <v-container>
                    <v-row>
                        <v-col cols="12" md="12">
                            <v-text-field v-model="id" :rules="[required]" label="id"></v-text-field>
                        </v-col>
                        <v-col cols="12" md="12">
                            <v-text-field type="password" v-model="pwd" :rules="[required, limit_length]" label="password"></v-text-field>
                        </v-col>
                        <!-- <v-col cols="12" md="12">
                            <v-text-field type="password" v-model="confPwd" :rules="[required, if_matching]" label="confirm password"></v-text-field>
                        </v-col> -->
                        <v-btn class="mr-4" v-on:click="submit">submit</v-btn>
                        <v-btn v-on:click="clear">clear</v-btn>
                    </v-row>
                </v-container>
            </v-form>
        </v-row>
        <v-row class="mt-7 ml-2">
            <div>
                <h2 v-if="errored">Sorry. Couldn't log in!</h2>
            </div>
            <div>
                <h2 v-if="succeeded">Successful to log in!</h2>
            </div>
        </v-row>
    </v-container>
</template>

<script>
import axios from "axios"

export default {
    name: "SignUp",
    data: () => ({
        id: "",
        pwd: "",
        confPwd: "",
        required: value => !!value || "Must include any letter!",
        limit_length: value => value.length >= 8 || "Must include 8 letters or more!",
        // if_matching: value => value !=  || "Not same as password!",
        info: null,
        errored: false,
        succeeded: false
    }),
    methods: {
        submit() {
            if(this.$refs.create_form.validate()) {
                this.requestLogIn(this.id, this.id);
                this.$refs.create_form.reset();
            }
        },
        clear() {
            this.$refs.create_form.reset();
        },
        requestLogIn(id, pwd) {
            const params = new URLSearchParams();
            params.append("id", id);
            params.append("pwd", pwd);
            axios.post("http://localhost:3000/login", params, {withCredentials:true}) // 異なるオリジンにアクセスする場合はwithCredentialsをtrueにする
            .then(response => {
                this.info = response.bpi;
                console.log(response.headers["set-cookie"]);
                this.succeeded = true;
            })
            .catch(error => {
                console.log(error)
                this.errored = true;
                this.succeeded = false;
            });
        }
    }
}
</script>