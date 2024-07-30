import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

export const useGalleryStore = defineStore({
  id: 'Gallery',
  state: () => ({
    gallery: []
  }),
  getters: {},
  actions: {
    // Fetch followers from action
    async fetchGallery() {
      try {
        const response = await axios.get('/api/gallery/list');
        this.gallery = response.data.gallery;
      } catch (error) {
        alert(error);
      }
    }
  }
});
