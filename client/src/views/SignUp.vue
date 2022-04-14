<template>
    <v-container>
        <v-row class="mt-2 ml-2">
        <h1>Sign up!</h1>
        </v-row>
        <v-row class="mt-2 ml-2">
            <v-form ref="create_form">
                <v-container>
                    <v-row>
                        <v-col cols="12" md="12">
                            <v-text-field v-model="content" :rules="[required, limit_length]" label="id" counter=50></v-text-field>
                        </v-col>
                        <v-col cols="12" md="12">
                            <v-text-field v-model="content" :rules="[required, limit_length]" label="password" counter=50></v-text-field>
                        </v-col>
                        <v-col cols="12" md="12">
                            <v-text-field v-model="content" :rules="[required, limit_length]" label="confirm password" counter=50></v-text-field>
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
export default {

}
</script>

<style>

</style>