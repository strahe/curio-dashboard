<script setup lang="ts">
import {useQuery} from "@vue/apollo-composable";
import {GetConfigs} from "@/pages/configs/query";
import {Config} from "@/typed-graph";
import {definePage} from "unplugin-vue-router/runtime";
import {Codemirror} from "vue-codemirror";
import {oneDark} from "@codemirror/theme-one-dark";
import {StreamLanguage} from "@codemirror/language";
import {toml} from "@codemirror/legacy-modes/mode/toml"
import {ComputedRef} from "vue";

definePage({
  meta: {
    title: 'Configs',
  },
})

const { result, loading, refetch, error } = useQuery(GetConfigs, null, () => ({
  fetchPolicy: 'cache-first',
}));
const items: ComputedRef<[Config]> = computed(() => result.value?.configs || []);
const headers = ref([
  { title: 'ID', key: 'id' },
  { title: 'Name', key: 'title' },
  { title: 'Used By', key: 'usedBy' },
  { title: '    ', key: 'edit', sortable: false },
]);

const dialog = ref(false);
const editItem = ref()

const extensions = [StreamLanguage.define(toml), oneDark]
function openDialog(data: Config) {
  dialog.value = true;
  editItem.value = data;
}

function saveEdit() {
  dialog.value = false;
}

</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="12">
        <Card title="Configs" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading"></v-btn>
          </template>
          <v-data-table
            :headers="headers"
            :items="items"
            :loading="loading"
          >
            <template v-slot:[`item.usedBy`]="{ item }">
              <v-chip
                v-for="usedBy in item.usedBy"
                :key="usedBy?.id"
                :title="usedBy?.machineName"
                color="primary"
                label
              >
                {{usedBy?.machineId }}
              </v-chip>
            </template>
            <template v-slot:[`item.edit`]="{ item }">
              <v-icon
                class="me-2"
                size="small"
                @click="openDialog(item)"
              >
                mdi-pencil
              </v-icon>
            </template>
          </v-data-table>

          <v-dialog
            v-model="dialog"
            transition="dialog-bottom-transition"
            fullscreen
          >
            <v-card>
              <v-toolbar>
                <v-btn
                  icon="mdi-close"
                  @click="dialog = false"
                ></v-btn>

                <v-toolbar-title>{{editItem.title}}</v-toolbar-title>
                <v-spacer></v-spacer>

                <v-toolbar-items>
                  <v-btn
                    @click="saveEdit"
                    disabled
                  >
                    Save
                  </v-btn>
                </v-toolbar-items>
              </v-toolbar>

              <v-card-text>
                <Codemirror
                  v-model="editItem.config"
                  :extensions="extensions"
                  :autofocus="true"
                  :indent-with-tab="true"
                  :tab-size="2"
                />
              </v-card-text>
            </v-card>
          </v-dialog>
        </Card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
