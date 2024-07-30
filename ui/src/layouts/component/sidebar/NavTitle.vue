<script setup>
const props = defineProps({ item: Object });
</script>

<template>
  <h6 class="text-subtitle-1 text-darkText pl-6 my-5">{{ $t(props.item.header) }}</h6>
</template>
