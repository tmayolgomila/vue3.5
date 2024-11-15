<template>

  <div class="fixed inset-0 flex items-center justify-center w-full h-full ">
    <div class="bg-white dark:bg-neutral-800 w-96 h-64 p-4 rounded shadow-lg relative z-60">

      <button @click="closeModal" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700">
        <ion-icon size="large" name="close-outline"></ion-icon>
      </button>

      <div class="mb-4">
        <h2 v-if="!isEditingTitle" class="text-xl font-semibold mb-4 capitalize" @dblclick="enableEditing">
          {{ cardTitle }}</h2>
        <input v-else v-model="editableTitle" @blur="saveTitle" @keydown.enter="saveTitle"
          class="text-xl font-semibold rounded w-2/4 p-1 darl:bg-neutral-800 text-black" placeholder="Edit title"
          autofocus />

        <div class="mt-4">
          <label for="description" class="block font-medium">Description:</label>
          <textarea id="description" v-model="cardDescription" @blur="saveDescription" @keydown.enter="saveDescription"
            class="w-full rounded p-2 dark:bg-neutral-800" rows="3" placeholder="Enter description"></textarea>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps({
  cardId: {
    type: Number,
    required: true
  },
  cardTitle: {
    type: String,
    required: true
  },
  columnId: {
    type: Number,
    required: true
  },
  projectId: {
    type: Number,
    required: true
  },
  cardDescription: {
    type: String,
    required: false,
    default: '',
  }
})

const emit = defineEmits(['updateTitle', 'updateDescription', 'close'])
const isEditingTitle = ref(false)
const editableTitle = ref(props.cardTitle)
const cardDescription = ref(props.cardDescription)

const enableEditing = () => {
  isEditingTitle.value = true
}

const saveTitle = () => {
  isEditingTitle.value = false
  if (editableTitle.value !== props.cardTitle) {
    emit('updateTitle', editableTitle.value)
  }
}

const saveDescription = () => {
  emit('updateDescription', cardDescription.value)
}


const closeModal = () => {
  emit('close');
};

watch(() => props.cardDescription, (newDescription) => {
  cardDescription.value = newDescription
})

</script>
