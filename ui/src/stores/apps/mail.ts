import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

export const useMailStore = defineStore({
  id: 'mail',
  state: () => ({
    mails: [],
    activeFilter: 'all',
    unreadCount: undefined
  }),
  getters: {},
  actions: {
    // Fetch mail from action
    async fetchMails() {
      try {
        const response = await axios.get('/api/mails/list');
        this.mails = response.data.mails;
        this.unreadCount = response.data.unreadCount;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch filtered mail from action
    async filterMails(filter: string) {
      try {
        const response = await axios.post('/api/mails/filter', { filter });
        this.mails = response.data;
      } catch (error) {
        alert(error);
      }
    },
    // setRead
    async setRead(id: string) {
      try {
        await axios.post('/api/mails/setRead', { id });
      } catch (error) {
        alert(error);
      }
    },
    // Set Starred
    async setStarred(id: string) {
      try {
        await axios.post('/api/mails/setStarred', { id });
      } catch (error) {
        alert(error);
      }
    },
    // Set Imporatnt
    async setImportant(id: string) {
      try {
        await axios.post('/api/mails/setImportant', { id });
      } catch (error) {
        alert(error);
      }
    },
    //select filter
    SelectFilter(item: string) {
      this.activeFilter = item;
    }
  }
});
