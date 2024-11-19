<template>

  <div class="fixed inset-0 flex items-center justify-center w-full h-full ">
    <div class="bg-white dark:bg-neutral-800 w-96 h-64 p-4 rounded shadow-lg relative z-60">

      <h2 class="text-lg font-semibold mb-4">Move cards to another column</h2>

      <ul class="space-y-2">
        <li v-for="column in otherColumns" :key="column.id">
          <button @click="moveCards(column.id)"
            class="w-full text-left p-2 bg-gray-100 dark:bg-neutral-700 hover:bg-gray-200 dark:hover:bg-neutral-600 rounded">
            {{ column.name }}
          </button>
        </li>
      </ul>
      <button @click="closeModal" class="mt-4 text-red-500 hover:underline">
        Cancel
      </button>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Column } from '../store/projectStore';


const props = defineProps({
  columnId: { type: Number, required: true },
  allColumns: { type: Array as () => Column[], required: true },
})

const emits = defineEmits(['moveCardsToColumn', 'close'])

const otherColumns = ref(
  props.allColumns ? props.allColumns.filter(column => column.id !== props.columnId) : []
);

const moveCards = (targetColumnId: number) => {

  emits('close')
  emits('moveCardsToColumn', targetColumnId)

}

const closeModal = () => {
  emits('close')
}

</script>
