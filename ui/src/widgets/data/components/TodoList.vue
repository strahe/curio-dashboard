<script setup lang="ts">
import { shallowRef } from 'vue';

const todos = shallowRef([
  {
    item: 'Check your Email',
    done: true
  },
  {
    item: 'Make YouTube Video',
    done: true
  },
  {
    item: 'Create Banner',
    done: true
  },
  {
    item: 'Upload Project',
    done: false
  },
  {
    item: 'Update Task',
    done: false
  },
  {
    item: 'Task Issue',
    done: false
  },
  {
    item: 'Deploy Project',
    done: false
  }
]);
</script>

<template>
  <div class="pa-5">
    <div class="pb-0 todo-list" v-for="(todo, i) in todos" :key="i" :value="todo.item">
      <v-checkbox :model-value="todo.done" color="primary" density="compact" hide-details>
        <template v-slot:label>
          <div class="ml-2">
            {{ todo.item }}
          </div>
        </template>
      </v-checkbox>
    </div>
  </div>
</template>
