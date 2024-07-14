<template>
  <v-footer height="40" app>
    <a
      v-for="item in items"
      :key="item.title"
      :href="item.href"
      :title="item.title"
      class="d-inline-block mx-2 social-link"
      rel="noopener noreferrer"
      target="_blank"
    >
      <v-icon
        :icon="item.icon"
        size="20"
      />
    </a>

    <div
      class="text-caption text-disabled"
      style="position: absolute; right: 16px;"
    >
      &copy; 2024-{{ (new Date()).getFullYear() }} <span class="d-none d-sm-inline-block">Curio Dashboard</span>
      —
      <a
        class="text-decoration-none on-surface"
        href="https://github.com/strahe/curio-dashboard/blob/main/LICENSE"
        rel="noopener noreferrer"
        target="_blank"
      >
        Apache License
      </a>
    </div>
  </v-footer>
</template>

<script setup lang="ts">
  const items = [
    {
      title: 'GitHub',
      icon: `mdi-github`,
      href: 'https://github.com/strahe/curio-dashboard',
    }
  ]
</script>

<style scoped lang="sass">
</style>
