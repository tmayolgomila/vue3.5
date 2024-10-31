<template>
  <div class="bg-gray-100 h-screen flex flex-col max-w-lg mx-auto">
    <div class="bg-blue-500 p-4 text-white flex justify-between items-center">
      <span>User 2</span>
    </div>

    <ChatMessages :messages="messages" />

    <MessageBox @send-message="onMessage" />

  </div>
</template>


<script setup lang="ts">
import ChatMessages from '@/components/chat/ChatMessages.vue';
import MessageBox from '@/components/chat/MessageBox.vue';
//import type { ChatMessage } from '@/interfaces/chat-message-interface';
//import { ref } from 'vue';
import { useChat } from '../composables/useChat';

const { messages, onMessage } = useChat();



</script>
