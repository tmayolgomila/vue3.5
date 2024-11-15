<template>

  <div class="p-2 shadow w-80 h-full bg-white dark:bg-neutral-900" ref="columnContainer">

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
        <div
          class="p-2 bg-gray-50 hover:bg-gray-100 dark:bg-neutral-800 dark:hover:bg-neutral-700 rounded mb-2 shadow-md cursor-pointer"
          @click="openCard(card.id, card.title)">
          <h4 class="font-semibold capitalize">{{ card.title }}</h4>
        </div>
      </template>
    </draggableComponent>

    <button @click="addCard" class=" mt-2 flex items-center space-x-1">
      <ion-icon name="add-outline"></ion-icon><span>Add New Card</span>
    </button>

    <ColumnModal v-if="showModal" :columnId="column.id" :position="modalPosition" @deleteColumn="deleteColumn"
      @close="showModal = false" />

  </div>

  <CardDetail v-if="internalIsModalOpen" :key="`${selectedCardId}-${selectedCardTitle}`" :cardId="selectedCardId"
    :cardTitle="selectedCardTitle" @updateTitle="updateCardTitle" @close="closeCardDetail" />



</template>

<script lang="ts" setup>
import { useProjectStore } from '../store/projectStore';
import { computed, ref } from 'vue';

import draggableComponent from 'vuedraggable';
import ColumnModal from './ColumnModal.vue';
import CardDetail from './CardDetail.vue';


const props = defineProps({
  column: {
    type: Object,
    required: true,
  },
  columnId: {
    type: Number,
    required: true,
  },
  projectId: {
    type: Number,
    required: true,
  }
})

const projectStore = useProjectStore()
const showModal = ref(false)
const modalPosition = ref({ top: 0, left: 0 })
const isEditing = ref(false)
const editedColumnName = computed(() => props.column.name);


const selectedCardId = ref(0)
const selectedCardTitle = ref('')

const emits = defineEmits(['addCard', 'editCard', 'deleteCard', 'dragCard', 'editColumn', 'deleteColumn'])

const openModal = (event: MouseEvent) => {
  const columnContainer = (event.currentTarget as HTMLElement).closest('.p-2')
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

const onDragEnd = (event: { to: HTMLElement; item: HTMLElement; newIndex: number }) => {
  emits('dragCard', {
    fromColumnId: props.column.id,
    toColumnId: event.to.dataset.columnId,
    cardId: event.item.dataset.cardId,
    newIndex: event.newIndex,
  })
}

const internalIsModalOpen = ref(false);

const openCard = (cardId: number, cardTitle: string) => {
  selectedCardId.value = cardId;
  selectedCardTitle.value = cardTitle;
  internalIsModalOpen.value = true;
};

const closeCardDetail = () => {
  internalIsModalOpen.value = false;
  selectedCardId.value = 0;
  selectedCardTitle.value = '';
};

const updateCardTitle = (newTitle: string) => {
  if (selectedCardTitle.value) {
    projectStore.editCardInColumn(props.projectId, props.columnId, selectedCardId.value, newTitle)
    selectedCardTitle.value = newTitle
  }
}

</script>
