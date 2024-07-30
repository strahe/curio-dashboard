import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';
import type { UserProfile } from '@/types/user/userProfile';

interface contactType {
  contact: UserProfile[];
  selectedContact: number;
}

export const useContactStore = defineStore({
  id: 'contact',
  state: (): contactType => ({
    contact: [],
    selectedContact: 0
  }),
  getters: {},
  actions: {
    // Fetch contacts
    async fetchContacts() {
      try {
        const data = await axios.get('/api/contact/list');
        this.contact = data.data;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch contacts
    async editContacts(contact: UserProfile) {
      try {
        const response = await axios.post('/api/contact/modify', contact);
        this.contact = response.data;
      } catch (error) {
        alert(error);
      }
    },
    // Fetch contacts
    async addContacts(contact: UserProfile) {
      try {
        const response = await axios.post('/api/contact/add', contact);
        this.contact = response.data;
      } catch (error) {
        alert(error);
      }
    },

    //select chat
    SelectContact(itemID: number) {
      this.selectedContact = itemID - 1;
    }
  }
});
