<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {Task} from "@/typed-graph";
import {format} from "timeago.js";
import TaskHistory from "@/widgets/TaskHistory.vue";

const props = defineProps({
  id: {
    type: Number,
    required: true
  }
});

const { result, loading,refetch,error } = useQuery(gql`
  query GetTask($id: Int!) {
    task(id: $id) {
      id
      name
      addedBy {
        hostAndPort
      }
      initiatedBy {
        hostAndPort
      }
      owner{
        hostAndPort
        detail {
          id
          machineName
          startupTime
        }
      }
      postedTime
      updateTime
      previousTask {
        id
        name
      }
    }
  }
`, {
  id: props.id,
}, () => ({
  fetchPolicy: 'cache-first',
}));
const data: ComputedRef<Task> = computed(() => result.value?.task || {});
</script>

<template>
<v-container fluid>
  <v-row>
    <v-col cols="6">
      <v-card :loading="loading">
        <v-alert v-if="error" type="error">
          {{ error.message }}
        </v-alert>
        <v-card-title class="d-flex justify-space-between">
          <div>Details</div>
          <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
        </v-card-title>
        <v-card-text>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>ID: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="data.id"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>Name: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="data.name"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>Owner: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="data.owner?.hostAndPort || 'N/A'"
                  prepend-inner-icon="mdi-server"
                  :label="'Startup: ' + (format(data.owner?.detail?.startupTime) || 'N/A')"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>Previous Task: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="data.previousTask?.name"
                  :label="String(data.previousTask?.id)"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>Posted: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="format(data.postedTime)"
                  :label="data.postedTime"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="3">
                <v-list-subheader>Updated: </v-list-subheader>
              </v-col>

              <v-col cols="9">
                <v-text-field
                  :model-value="format(data.updateTime)"
                  :label="data.updateTime"
                  disabled
                ></v-text-field>
              </v-col>
            </v-row>
        </v-card-text>
      </v-card>
    </v-col>
    <v-col cols="6">
      <TaskHistory :id="props.id"></TaskHistory>
    </v-col>
  </v-row>
</v-container>
</template>

<style scoped>

</style>
