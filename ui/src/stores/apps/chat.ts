import { defineStore } from 'pinia';
// project imports
import axios from '@/utils/axios';

import type { ChatTypes, ChatHistory } from '@/types/chats/ChatTypes';

interface chatType {
  chats: ChatTypes[];
  chatContent: number;
}

export const useChatStore = defineStore({
  id: 'chat',
  state: (): chatType => ({
    chats: [],
    chatContent: 1
  }),
  getters: {
    // Get Chats from Getters
    getChats(state) {
      return state.chats;
    }
  },
  actions: {
    // Fetch Chat from action
    async fetchChats() {
      try {
        const data = await axios.get('/api/data/chat/ChatData');
        this.chats = data.data;
      } catch (error) {
        alert(error);
      }
    },
    //select chat
    SelectChat(itemID: number) {
      this.chatContent = itemID - 1;
    },

    sendMsg(this, itemID: number, item: string) {
      this.chats = this.chats.map((chat) => {
        if (chat.id === itemID) {
          const updatedChatHistory: ChatHistory[] = chat.chatHistory.map((historyItem) => ({
            from: { ...historyItem.from, id: historyItem.from.id || 0 },
            to: { ...historyItem.to, id: historyItem.to.id || 0, to: historyItem.to.to ? [...historyItem.to.to, item] : [item] }
          }));
          return {
            ...chat,
            chatHistory: updatedChatHistory
          };
        }
        return chat;
      });
    }
  }
});
