<script setup lang="ts">
import { PropType } from 'vue'
import { useI18n } from 'vue-i18n'
import { RouteLocationAsPathGeneric, RouteLocationAsRelativeGeneric } from 'vue-router'

const { t } = useI18n()

interface Breadcrumb {
  title: string;
  disabled: boolean;
  to: string | RouteLocationAsRelativeGeneric | RouteLocationAsPathGeneric;
}

const props = defineProps({
  title: {
    type: String,
    default: undefined,
  },
  breadcrumbs: {
    type: Array as PropType<Breadcrumb[]>,
    required: true,
  },
  icon: {
    type: String,
    default: undefined,
  },
})

</script>

<template>
  <v-row class="page-breadcrumb mb-0 mt-n2">
    <v-col
      cols="12"
      md="12"
    >
      <v-card
        elevation="0"
        variant="text"
      >
        <v-row
          class="align-center"
          no-gutters
        >
          <v-col sm="12">
            <v-breadcrumbs
              class="text-h6 pa-1 font-weight-medium mb-0"
              :items="props.breadcrumbs"
            >
              <template #divider>
                <div class="d-flex align-center">
                  /
                </div>
              </template>
              <template #prepend>
                <router-link
                  class="text-lightText text-h6 text-decoration-none"
                  to="/"
                >
                  {{ t('nav.Home') }}
                </router-link>
                <div class="d-flex align-center px-2">
                  /
                </div>
              </template>
            </v-breadcrumbs>
          </v-col>
        </v-row>
      </v-card>
    </v-col>
  </v-row>
</template>

<style lang="scss">
.page-breadcrumb {
  .v-toolbar {
    background: transparent;
  }
}
</style>
