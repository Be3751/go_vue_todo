<template>
    <v-card class="mt-5 mb-5 ml-2 mr-2" style="width: 30%;">
        <v-card-title>{{ task.content }}</v-card-title>
        <v-card-actions>
            <v-btn outlined rounded text color="green">Update</v-btn>
            <v-btn outlined rounded text color="red">Delete</v-btn>
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
    }
}
</script>