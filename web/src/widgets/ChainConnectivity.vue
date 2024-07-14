<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import gql from "graphql-tag";
import {NodeInfo} from "@/typed-graph";
import {ComputedRef} from "vue";

defineProps({
  title: {
    type: String,
    default: undefined,
  },
  color: {
    type: String,
    default: undefined,
  }
})

const { result, loading, refetch, error } = useQuery(gql`
  query GetNodesInfo {
    nodesInfo {
      id
      address
      reachability
      epoch
      behind
      peersToPublishBlocks
      peersToPublishMsgs
      version
    }
  }
`, null, () => ({
  fetchPolicy: 'cache-first',
}));

const items: ComputedRef<[NodeInfo]> = computed(() => result.value?.nodesInfo || []);

const headers = ref([
  {title: "Address", key: "address"},
  {title: "Reachability", key: "reachability"},
  {title: "Epoch", key: "epoch"},
  {title: "Status", key: "status"},
  {title: "Version", key: "version"},
]);

</script>

<template>
  <Card :title="title" :color="color" :error="error as Error">
    <template #titleAction>
      <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
    </template>
    <v-card-text>
      <v-data-table-virtual
        :headers="headers"
        :items="items"
        :loading="loading"
        item-key="address"
      >
        <template v-slot:[`item.reachability`]="{ value }">
          <v-icon
            v-if="value"
            color="success"
            icon="mdi-check-circle-outline"
          ></v-icon>
          <v-icon
            v-else
            color="error"
            icon="mdi-alpha-x-circle-outline"
          ></v-icon>
        </template>
        <template v-slot:[`item.status`]="{ item }">
          <v-chip
            v-if="item.behind > 0"
            color="error"
          >Behind {{ item.behind }} Epoch</v-chip>
          <v-icon
            v-else
            color="success"
            icon="mdi-check-circle-outline"
          ></v-icon>
        </template>
      </v-data-table-virtual>
    </v-card-text>
  </Card>
</template>

<style scoped>

</style>
