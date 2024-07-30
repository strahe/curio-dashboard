<script setup>
const props = defineProps({ item: Object, level: Number });
</script>

<template>
  <!---Single Item-->
  <router-link :to="`${item.to}`" class="navItemLink rounded-0" :disabled="item.disabled">
    <!---If icon-->
    <i class="navIcon">
      <component :is="props.item.icon" :level="props.level" :style="{ fontSize: '16px' }"></component>
    </i>
    <span>{{ item.title }}</span>
    <!---If Caption-->
    <small v-if="item.subCaption" class="text-caption mt-n1 hide-menu">
      {{ item.subCaption }}
    </small>
    <!---If any chip or label-->
    <template v-if="item.chip">
      <v-chip
        :color="item.chipColor"
        class="sidebarchip hide-menu ml-auto"
        :size="item.chipIcon ? 'small' : 'small'"
        :variant="item.chipVariant"
        :prepend-icon="item.chipIcon"
      >
        {{ item.chip }}
      </v-chip>
    </template>
  </router-link>
</template>
