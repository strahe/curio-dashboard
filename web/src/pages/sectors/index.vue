<script setup lang="ts">
import {GetSectorMeta, GetSectorsMeta} from "@/pages/sectors/query";
import {useQuery} from "@vue/apollo-composable";
import {Actor, Sector} from "@/typed-graph";
import {definePage} from "unplugin-vue-router/runtime";

definePage({
  meta: {
    title: 'Sectors',
  },
})

const itemsPerPage = ref(15)
const currentPage = ref(1)
const selectedMiner = ref();
const inputSector = ref();

const isSingleSearch = computed(() => selectedMiner.value && inputSector.value);

function buildVariables() {
  const variables = {
    actor: undefined,
    sectorNumber: undefined,
    offset: (currentPage.value-1) * itemsPerPage.value,
    limit: itemsPerPage.value
  }
  if (selectedMiner.value) {
    variables.actor = selectedMiner.value
  }
  if (isSingleSearch && inputSector.value) {
    variables.sectorNumber = inputSector.value
  }
  return variables
}

const { result, loading, fetchMore, error } = useQuery(GetSectorsMeta,
  buildVariables(), ()=>({
  fetchPolicy: 'cache-first',
}))

const items = computed(() => result.value?.sectors.map((sector: Sector) => sector.meta) || []);
const itemsCount = computed(() => result.value?.sectorsCount || 0);
const actors = computed(() => result.value?.actors.map((actor: Actor) => actor.address) || []);
const expanded = ref([]);

const headers = ref([
  { title: 'Miner', key: 'spId' },
  { title: 'Sector', key: 'sectorNum', sortable: false },
  { title: 'Seed Epoch', key: 'seedEpoch', sortable: false },
  { title: 'Unsealed', key: 'curUnsealedCid', sortable: false },
  { title: 'Sealed', key: 'curSealedCid', sortable: false },
  { title: 'Precommit', key: 'msgCidPrecommit', sortable: false },
  { title: 'Commit', key: 'msgCidCommit', sortable: false },
  { title: '  ', key: 'data-table-expand', sortable: false },
]);

function loadForPagination() {
  fetchMore({
    query: isSingleSearch ? GetSectorMeta : GetSectorsMeta,
    variables: buildVariables(),
    updateQuery(previousQueryResult, {fetchMoreResult}) {
      return {
        sectors: fetchMoreResult.sectors,
        sectorsCount: isSingleSearch ? fetchMoreResult.sectors?.length: fetchMoreResult.sectorsCount,
        actors: isSingleSearch? previousQueryResult.actors: fetchMoreResult.actors
      }
    },
  });
}
</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center"></v-row>
    <v-col cols="12">
      <Card title="Sectors" :error="error as Error">
        <template #titleAction>
          <v-btn icon="mdi-refresh" @click="loadForPagination" :disabled="loading" size="small"></v-btn>
        </template>
        <v-row>
          <v-col cols="3">
            <v-select
              v-model="selectedMiner"
              prepend-inner-icon="mdi-account-hard-hat"
              variant="outlined"
              :items="actors"
              label="Miner"
              @update:model-value="loadForPagination"
              clearable
            >
            </v-select>
          </v-col>
          <v-col cols="3">
            <v-text-field
              v-model="inputSector"
              type="number"
              label="Sector Number"
              prepend-inner-icon="mdi-magnify"
              variant="outlined"
              hint="Enter sector number"
              @keyup.enter="loadForPagination"
              :disabled="!selectedMiner"
              hide-details
              single-line
            ></v-text-field>
          </v-col>
        </v-row>

        <v-data-table-server
          v-model:expanded="expanded"
          v-model:items-per-page="itemsPerPage"
          v-model:page="currentPage"
          :headers="headers"
          :items="items"
          :items-length="itemsCount"
          :loading="loading"
          @update:page="loadForPagination"
          @update:itemsPerPage="loadForPagination"
          show-expand
        >
          <template v-slot:loading v-if="!items">
            <v-skeleton-loader type="table-row@15"></v-skeleton-loader>
          </template>
          <template v-slot:[`item.curUnsealedCid`]="{value}">
            <p
              class="d-inline-block text-truncate"
              style="max-width: 150px;"
            >{{value}}</p>
          </template>
          <template v-slot:[`item.curSealedCid`]="{value}">
            <p
              class="d-inline-block text-truncate"
              style="max-width: 150px;"
            >{{value}}</p>
          </template>
          <template v-slot:[`item.msgCidPrecommit`]="{value}">
            <p
              class="d-inline-block text-truncate"
              style="max-width: 150px;"
            >{{value}}</p>
          </template>
          <template v-slot:[`item.msgCidCommit`]="{value}">
            <p
              class="d-inline-block text-truncate"
              style="max-width: 150px;"
            >{{value}}</p>
          </template>
          <template v-slot:expanded-row="{ columns, item }">
            <tr>
              <td :colspan="columns.length">
                <v-row class="mt-2">
                  <template v-for="(value, key) in item" :key="key">
                    <v-col cols="6" v-if="String(key) !== '__typename'">
                      <v-text-field
                        :label="String(key)"
                        :model-value="value"
                        readonly
                        disabled
                      ></v-text-field>
                    </v-col>
                  </template>
                </v-row>
              </td>
            </tr>
          </template>
        </v-data-table-server>
      </Card>
    </v-col>
  </v-container>
</template>

<style scoped>

</style>
