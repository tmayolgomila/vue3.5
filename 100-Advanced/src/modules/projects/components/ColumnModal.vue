<template>
  <div class="absolute bg-white dark:bg-slate-800 p-6 rounded shadow-md w-80 column"
    :style="{ top: `${position.top}px`, left: `${position.left}px`, position: 'absolute', zIndex: 50 }">


    <button @click="$emit('close')" class="absolute top-2 right-2"><ion-icon size="large"
        name="close-outline"></ion-icon></button>

    <div class=" flex flex-col items-start space-y-2">

      <button class="text-lg" @click="addCard">Add Card</button>
      <button class="text-lg" @click="copyList">Copy List</button>
      <button class="text-lg" @click="moveAllCardsToOtherList">Move all card list to another</button>
      <button class="text-lg">Sort by</button>
      <button class="text-lg">Follow</button>

      <br>

      <h2 class="text-lg ">Change column color:</h2>
      <div class="flex flex-items gap-1">
        <button @click="changeColor('bg-red-500')" class="bg-red-500 w-20 h-10"></button>
        <button @click="changeColor('bg-yellow-500')" class="bg-yellow-500 w-20 h-10"></button>
        <button @click="changeColor('bg-orange-500')" class="bg-orange-500 w-20 h-10"></button>

      </div>
      <div class="flex flex-items gap-1">
        <button @click="changeColor('bg-blue-500')" class="bg-blue-500 w-20 h-10"></button>
        <button @click="changeColor('bg-green-500')" class="bg-green-500 w-20 h-10"></button>
        <button @click="changeColor('bg-fuchsia-500')" class="bg-fuchsia-500 w-20 h-10"></button>
      </div>

      <button @click="changeColor('')" class="flex items-center justify-center space-x-1">
        <span class="text-center">Remove color </span>
        <ion-icon name="close-outline" class="ion-icon-med"></ion-icon>
      </button>

      <br>

      <button @click="showDeleteModal = true" class="text-red-500 text-lg font-semibold">Remove List</button>
    </div>

  </div>

  <ConfirmModal v-if="showDeleteModal" title="Delete List"
    message="Are you sure you want to delete this list? This action cannot be undone." confirmAction="Delete List"
    @confirm="confirmDelete" @cancel="cancelDelete" />

</template>

<script setup lang="ts">
import { ref } from 'vue';
import ConfirmModal from './ConfirmModal.vue';

const props = defineProps({
  columnId: {
    type: Number,
    required: true,
  },
  position: {
    type: Object,
    required: true,
    validator: (val: any) => 'top' in val && 'left' in val,
  }
})

const emits = defineEmits(['addCard', 'deleteColumn', 'close', 'copyList', 'changeColor'])

const showDeleteModal = ref(false)

const addCard = () => {
  emits('close')
  emits('addCard')
}

const copyList = () => {
  emits('close')
  emits('copyList')
}

const confirmDelete = () => {
  emits('deleteColumn', props.columnId)
  showDeleteModal.value = false
}

const cancelDelete = () => {
  showDeleteModal.value = false
}

const changeColor = (color: string) => {
  emits('close')
  emits('changeColor', { columnId: props.columnId, color })
}

const moveAllCardsToOtherList = () => {
  emits('close')
}

</script>


<style>
.ion-icon-med {
  font-size: 24px;
}
</style>
