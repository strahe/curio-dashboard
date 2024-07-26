<script setup lang="ts">
import {useMutation, useQuery} from "@vue/apollo-composable";
import {CreateConfig, GetConfigs, UpdateConfig} from "@/pages/configs/query";
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
const editTitle = ref('');
const editConfig = ref('');
const editCreate = ref(false);
const saveLoading = ref(false);

const extensions = [StreamLanguage.define(toml), oneDark]

function openDialog(title: string, config: string, isCreate: boolean) {
  editTitle.value = title;
  editConfig.value = config;
  editCreate.value = isCreate;
  dialog.value = true;
  saveLoading.value = false;
}

function saveEdit() {
  saveLoading.value = true;
  if (editCreate.value) {
    createConfig({
      title: editTitle.value,
      config: editConfig.value,
    }).then(() => {
      dialog.value = false;
      refetch()
    });
    return;
  }else {
    updateConfig({
      title: editTitle.value,
      config: editConfig.value,
    }).then(() => {
      dialog.value = false;
      refetch()
    });
  }
}

const { mutate: updateConfig  } = useMutation(UpdateConfig)
const { mutate: createConfig } = useMutation(CreateConfig)

</script>

<template>
  <v-container fill-height fluid grid-list-xl>
    <v-row justify="center">
      <v-col cols="12">
        <Card title="Configs" :error="error as Error">
          <template #titleAction>
            <v-btn icon="mdi-refresh" @click="refetch" :disabled="loading" size="small"></v-btn>
          </template>
          <v-btn color="primary" @click="openDialog('', '', true)">Create</v-btn>

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
                @click="openDialog(item.title, item.config, false)"
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

                <v-toolbar-title>
                  <v-text-field
                    v-model="editTitle"
                    :label="editCreate ? 'Title' : ''"
                    variant="underlined"
                    :disabled="!editCreate"
                  ></v-text-field>
                </v-toolbar-title>
                <v-spacer></v-spacer>

                <v-toolbar-items>
                  <v-btn
                    :loading="saveLoading"
                    color="primary"
                    @click="saveEdit"
                  >
                    Save
                  </v-btn>
                </v-toolbar-items>
              </v-toolbar>

              <v-card-text>
                <Codemirror
                  v-model="editConfig"
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
