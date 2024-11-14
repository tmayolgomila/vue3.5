<template>

  <div class="p-4 shadow w-64 h-full bg-white dark:bg-neutral-900" ref="columnContainer">

    <div class="flex items-center justify-between mb-2">

      <!--Edit title START-->
      <div @dblclick="editColumn" class="text-lg font-semibold">
        <template v-if="!isEditing">
          <button class="text-lg font-semibold">{{ column.name }}</button>
        </template>
        <template v-else>
          <input type="text" v-model="editedColumnName" @blur="saveColumnEdit" @keyup.enter="saveColumnEdit"
            class="text-lg font-semibold border-b text-black w-4/6" autofocus>
        </template>
      </div>
      <!--Edit title END-->

      <button @click="openModal">
        <ion-icon size="large" name="ellipsis-horizontal-outline"></ion-icon>
      </button>
    </div>

    <draggableComponent v-model="column.cards" group="cards" itemKey="id" @end="onDragEnd">
      <template #item="{ element: card }">
        <div class="p-2 bg-white dark:bg-neutral-800 rounded mb-2 shadow-md">
          <h4 class="font-semibold">{{ card.title }}</h4>
          <p>{{ card.description }}</p>
          <button @click="editCard(card.id)" class="text-blue-500"><ion-icon name="create-outline"></ion-icon></button>
          <button @click="deleteCard(card.id)" class="text-red-500">Delete</button>
        </div>
      </template>
    </draggableComponent>

    <button @click="addCard" class="text-blue-500 mt-2">Add New Card</button>

    <ColumnModal v-if="showModal" :columnId="column.id" :position="modalPosition" @deleteColumn="deleteColumn"
      @close="showModal = false" />

  </div>

</template>

<script lang="ts" setup>
import draggableComponent from 'vuedraggable';
import ColumnModal from './ColumnModal.vue';
import { ref, watch } from 'vue';


const props = defineProps({
  column: {
    type: Object,
    required: true
  }
})

const showModal = ref(false)
const modalPosition = ref({ top: 0, left: 0 })
const isEditing = ref(false)
const editedColumnName = ref(props.column.name)

const emits = defineEmits(['addCard', 'editCard', 'deleteCard', 'dragCard', 'editColumn', 'deleteColumn'])

const openModal = (event: MouseEvent) => {

  const columnContainer = (event.currentTarget as HTMLElement).closest('.p-4')
  const rect = columnContainer?.getBoundingClientRect()
  if (rect) {
    modalPosition.value = { top: rect.top, left: rect.left }
  }

  showModal.value = true
}
//start edition
const editColumn = () => {
  isEditing.value = true
};
//save new name for column
const saveColumnEdit = () => {
  if (editedColumnName.value.trim() !== props.column.name) {
    emits('editColumn', { columnId: props.column.id, newName: editedColumnName.value })
  }
  isEditing.value = false
}

const deleteColumn = () => {
  emits('deleteColumn', props.column.id);
};

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

watch(() => props.column.name, (newName) => {
  editedColumnName.value = newName
})

</script>
