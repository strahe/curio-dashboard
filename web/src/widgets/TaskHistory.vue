<script setup lang="ts">
import {format} from "timeago.js";
import { formatDurationFromStrings} from "@/utils/formatDuration";
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import { TaskHistory} from "@/typed-graph";

const props = defineProps({
  id: {
    type: Number,
    required: true
  }
});

const { result, loading,refetch, error } = useQuery(gql`
  query GetTaskHistories($id: Int!) {
    task(id: $id) {
      histories {
        id
        taskId
        name
        posted
        workStart
        workEnd
        result
        err
        completedByHostAndPort
      }
    }
  }
`, {
  id: props.id,
}, () => ({
  fetchPolicy: 'cache-first',
}));
const items: ComputedRef<[TaskHistory]> = computed(() => result.value?.task.histories || []);

const headers = ref([
  { title: 'Task', key: 'name'},
  { title: 'ID', key: 'task_id'},
  { title: 'Started', key: 'workStart' },
  { title: 'Took', key: 'took' },
  { title: 'Error', key: 'err' },
]);

</script>

<template>
<v-card>
  <v-alert v-if="error" type="error">
    {{ error.message }}
  </v-alert>
  <v-card-title class="d-flex justify-space-between">
    <div>Task History</div>
    <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
  </v-card-title>
  <v-data-table
    :items="items"
    :headers="headers"
    :loading="loading"
  >
    <template v-slot:[`item.workStart`]="{ item }">
      <v-chip>
        {{ format(item.workStart, 'en_US') }}
      </v-chip>
    </template>
    <template v-slot:[`item.took`]="{ item }">
      <v-chip>
        {{ formatDurationFromStrings(item.workStart, item.workEnd) }}
      </v-chip>
    </template>
  </v-data-table>
</v-card>
</template>

<style scoped>

</style>
