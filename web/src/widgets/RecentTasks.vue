<script setup lang="ts">
import {format} from "timeago.js";
import { formatDurationFromStrings} from "@/utils/formatDuration";
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import { TaskHistory} from "@/typed-graph";

const { result, loading,refetch, error } = useQuery(gql`
  query GetTaskHistories {
    taskHistories(offset: 0, limit: 100) {
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
`, {
}, () => ({
  fetchPolicy: 'cache-first',
}));
const items: ComputedRef<[TaskHistory]> = computed(() => result.value?.taskHistories || []);

const headers = ref([
  { title: 'Task', key: 'name'},
  { title: 'ID', key: 'task_id'},
  { title: 'End', key: 'workEnd' },
  { title: 'Took', key: 'took' },
  { title: 'Error', key: 'err' },
]);

</script>

<template>
  <Card title="Recent Tasks" :error="error as Error">
    <template #titleAction>
        <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
    </template>
    <v-data-table
      :items="items"
      :headers="headers"
      :loading="loading"
      items-per-page="5"
    >
      <template v-slot:[`item.workEnd`]="{ item }">
        <v-chip>
          {{ format(item.workEnd, 'en_US') }}
        </v-chip>
      </template>
      <template v-slot:[`item.took`]="{ item }">
        <v-chip>
          {{ formatDurationFromStrings(item.workStart, item.workEnd) }}
        </v-chip>
      </template>
    </v-data-table>
  </Card>
</template>

<style scoped>

</style>
