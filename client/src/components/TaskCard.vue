<template>
    <v-card class="mt-5 mb-5 ml-2 mr-2" style="width: 30%;">
        <v-card-title>{{ task.content }}</v-card-title>
        <v-card-actions>
            <v-btn v-on:click="toRead" outlined rounded text color="orange">Detail</v-btn>
            <v-btn v-on:click="toUpdate" outlined rounded text color="green">Update</v-btn>
            <v-btn v-on:click="toDelete" outlined rounded text color="red">Delete</v-btn>
        </v-card-actions>
    </v-card>
</template>

<script>
export default {
    name: "TaskCard",
    props: ["task"],
    data: function() {
        return {
            
        }
    },
    methods: {
        toRead() {
            this.$router.push({name: "read", params: {id: this.task.id, content: this.task.content, deadline: this.task.deadline}}) // パラメータidを指定してupdate/:idに遷移
        },
        toUpdate() {
            this.$router.push({name: "update", params: {id: this.task.id}}) // パラメータidを指定してupdate/:idに遷移
        },
        toDelete() {
            this.$router.push({name: "delete", params: {id: this.task.id, content: this.task.content}}) // パラメータidを指定してdelete/:idに遷移
        }
    }
}
</script>