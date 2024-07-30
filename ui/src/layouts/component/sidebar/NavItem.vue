<script setup>
const props = defineProps({ item: Object, level: Number });
</script>

<template>
  <!---Single Item-->
  <v-list-item
    :to="item.type === 'external' ? '' : item.to"
    :href="item.type === 'external' ? item.to : ''"
    rounded
    class="mb-1"
    color="primary"
    style="padding-left: 30px !important"
    :disabled="item.disabled"
    :target="item.type === 'external' ? '_blank' : ''"
  >
    <v-list-item-title>{{ $t(props.item.title) }}</v-list-item-title>
  </v-list-item>
</template>
