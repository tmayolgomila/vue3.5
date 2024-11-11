<template>

  <div class="border p-4 rounded shadow w-64">

    <div class="flex items-center justify-center mb-2">
      <h3 class="text-lg font-semibold">{{ column.name }}</h3>
      <div class="space-x-2">
        <button @click="editColumn" class="text-blue-500">Edit</button>
        <button @click="deleteColumn" class="text-red-500">Delete</button>
      </div>
    </div>

    <draggableComponent v-model="column.cards" group="cards" itemKey="id" @end="onDragEnd" >
      <template #item="{ element: card }">
        <div class="p-2 bg-gray-100 rounded mb-2">
          <h4 class="font-semibold">{{ card.title }}</h4>
          <p>{{ card.description }}</p>
          <button @click="editCard(card.id)" class="text-blue-500">Edit</button>
          <button @click="deleteCard(card.id)" class="text-red-500">Delete</button>
        </div>
      </template>
    </draggableComponent>
    <button @click="addCard" class="text-blue-500 mt-2">Add New Card</button>

  </div>

</template>

<script lang="ts" setup>
import draggableComponent from 'vuedraggable';


const props = defineProps({
  column: {
    type: Object,
    required: true
  }
})

const emits = defineEmits(['addCard', 'editCard', 'deleteCard', 'dragCard', 'editColumn', 'deleteColumn'])

const editColumn = () => {
  emits('editColumn', props.column.id)
}

const deleteColumn = () => {
  emits('deleteColumn', props.column.id)
}

const addCard = () => {
  emits('addCard', props.column.id)
}

const editCard = (cardId: number) => {
  emits('editCard', { columnId: props.column.id, cardId })
}

const deleteCard = (cardId: number) => {
  emits('deleteCard', { columnId: props.column.id, cardId })
}

const onDragEnd = (event: any) => {
  emits('dragCard', {
    fromColumnId: props.column.id,
    toColumnId: event.to.dataset.columnId,
    cardId: event.item.dataset.cardId,
    newIndex: event.newIndex,

  })
}

</script>
