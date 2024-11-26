<template>

  <div class="fixed inset-0 flex items-center justify-center w-full h-full ">
    <div class="bg-white dark:bg-neutral-800 w-96 h-64 p-4 rounded shadow-lg relative z-60">

      <button @click="closeModal" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700">
        <ion-icon size="large" name="close-outline"></ion-icon>
      </button>

      <div class="mb-4">
        <h2 v-if="!isEditingTitle" class="text-xl font-semibold mb-4 capitalize" @dblclick="enableEditing">
          {{ cardTitle }}</h2>
        <input v-else v-model="editableTitle" @keydown.enter="saveTitle"
          class="text-xl font-semibold rounded w-2/4 p-1 darl:bg-neutral-800 text-black new-title-input"
          placeholder="Edit title" autofocus />

        <div class="mt-4">
          <label for="description" class="block font-medium">Description:</label>
          <textarea id="description" v-model="cardDescription" @blur="saveDescription" @keydown.enter="saveDescription"
            class="w-full rounded p-2 dark:bg-neutral-800" rows="3" placeholder="Enter description"></textarea>
        </div>
      </div>
      <button class="text-red-700 text-lg font-bold" @click="deleteCard">
        Delete Card
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { nextTick, ref, watch } from 'vue';
import { useProjectStore } from '../store/projectStore';

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

const projectStore = useProjectStore()
const emit = defineEmits(['updateTitle', 'updateDescription', 'close', 'deleteCard'])
const isEditingTitle = ref(false)
const editableTitle = ref(props.cardTitle)
const cardDescription = ref(props.cardDescription)

const enableEditing = () => {
  isEditingTitle.value = true

  nextTick(() => {
    const inputElement = document.querySelector('.new-title-input') as HTMLInputElement;
    if (inputElement) {
      inputElement.focus()
    }
  })
}

const saveTitle = async () => {
  isEditingTitle.value = false;
  if (editableTitle.value !== props.cardTitle) {
    try {
      await projectStore.editCardInColumn(props.cardId, editableTitle.value, undefined);
      emit('updateTitle', editableTitle.value);
    } catch (error) {
      console.error('Error updating title:', error);
    }
  }
};

const saveDescription = async () => {
  if (cardDescription.value !== props.cardDescription) {
    try {
      await projectStore.editCardInColumn(props.cardId, undefined, cardDescription.value);
      emit('updateDescription', cardDescription.value);
    } catch (error) {
      console.error('Error saving description:', error);
    }
  }
};


const closeModal = () => {
  emit('close');
};

const deleteCard = () => {
  emit('deleteCard', props.projectId, props.columnId, props.cardId);
}

watch(() => props.cardDescription, (newDescription) => {
  cardDescription.value = newDescription
})

</script>
