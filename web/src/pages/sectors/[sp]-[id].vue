<script setup lang="ts">

const route = useRoute()

onMounted(() => {
  const params = route.params as { sp: string; id: string };
  console.log(params.sp)
  console.log(params.id)
})

</script>

<template>
<v-container fluid>
</v-container>
</template>

<style scoped>

</style>
