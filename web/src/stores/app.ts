// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    drawer: true,
    color: 'success',
  }),
  actions: {
    setDrawer(value: boolean) {
      this.drawer = value;
    },
  },
})
