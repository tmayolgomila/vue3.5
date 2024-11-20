<template>

  <div v-if="isVisible" class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
    <div class="bg-white dark:bg-slate-800 p-6 rounded-lg shadow-xl w-80">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-lg font-semibold">Sort by</h2>
        <button @click="$emit('close')" class="text-gray-500 hover:text-gray-700">
          <ion-icon name="close-outline" size="large"></ion-icon>
        </button>
      </div>

      <div class="space-y-4">
        <button @click="selectSort('title:asc')"
          class="w-full text-left p-3 hover:bg-gray-100 dark:hover:bg-slate-700 rounded">
          Title (A-Z)
        </button>
        <button @click="selectSort('title:desc')"
          class="w-full text-left p-3 hover:bg-gray-100 dark:hover:bg-slate-700 rounded">
          Title (Z-A)
        </button>
        <button @click="selectSort('id:asc')"
          class="w-full text-left p-3 hover:bg-gray-100 dark:hover:bg-slate-700 rounded">
          Date (Ascending)
        </button>
        <button @click="selectSort('id:desc')"
          class="w-full text-left p-3 hover:bg-gray-100 dark:hover:bg-slate-700 rounded">
          Date (Descending)
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

defineProps({
  isVisible: {
    type: Boolean,
    required: true
  }
})

const emits = defineEmits(['close', 'selectSort'])

const selectSort = (sortOption: string) => {
  emits('selectSort', sortOption)
  emits('close')
}
</script>
