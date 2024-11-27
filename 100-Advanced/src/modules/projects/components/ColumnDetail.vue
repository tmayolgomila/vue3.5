<template>

  <div :class="`${column.color || 'bg-white'} dark: ${column.color || 'dark:bg-neutral-900'} p-2 shadow w-80 h-full`"
    ref="columnContainer">

    <div class="flex items-center justify-between mb-2">

      <!--Edit title START-->
      <div @dblclick="editColumn" class="text-lg font-semibold">
        <template v-if="!isEditing">
          <button class="text-lg font-semibold">{{ column.name }}</button>
        </template>
        <template v-else>
          <input type="text" v-model="editedColumnName" @keyup.enter="saveColumnEdit"
            class="text-lg font-semibold border-b text-black w-4/6" autofocus>
        </template>
      </div>

      <!--Edit title END-->
      <button @click="openModal">
        <ion-icon size="large" name="ellipsis-horizontal-outline"></ion-icon>
      </button>
    </div>

    <draggableComponent v-model="cards" group="cards" itemKey="id" @end="onDragEnd" :animation="200"
      :ghost-class="'ghost'" :swap-threshold="0.5" :data-column-id="column.id">
      <template #item="{ element: card }">
        <div :data-card-id="card.id" :data-column-id="props.column.id"
          class="p-2 bg-gray-50 hover:bg-gray-100 dark:bg-neutral-800 dark:hover:bg-neutral-700 rounded mb-2 shadow-md cursor-pointer"
          @click="openCard(card.id, card.title, card.description)">
          <h4 class="font-semibold capitalize">{{ card.title }}</h4>
        </div>
      </template>
    </draggableComponent>

    <!-- Input for new card -->
    <div v-if="isAddingCard" class="mb-2">
      <input type="text" v-model="newCardTitle" @blur="cancelAddCard" @keydown.enter="saveNewCard"
        class="w-full p-2 border rounded text-black new-card-input" placeholder="Enter card title" autofocus />
    </div>

    <button @click="startAddCard" class=" mt-2 flex items-center space-x-1">
      <ion-icon name="add-outline"></ion-icon><span>Add New Card</span>
    </button>

    <ColumnModal v-if="showModal" :columnId="column.id" :position="modalPosition" :projectId="props.projectId"
      @addCard="startAddCard" @deleteColumn="deleteColumn" @close="showModal = false" @copyList="copyList"
      @changeColor="updateColumnColor" />

  </div>

  <CardDetail v-if="internalIsModalOpen" :key="`${selectedCardId}-${selectedCardTitle}`" :cardId="selectedCardId"
    :cardTitle="selectedCardTitle" :columnId="props.columnId" :projectId="props.projectId"
    :cardDescription="selectedCardDescription" @updateTitle="updateCardTitle" @updateDescription="updateCardDescription"
    @deleteCard="handleDeleteCard" @close="closeCardDetail" />

</template>

<script lang="ts" setup>
import { useProjectStore, type Card } from '../store/projectStore';
import { nextTick, onMounted, ref, watch } from 'vue';

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
const isAddingCard = ref(false);
const newCardTitle = ref('')
const editedColumnName = ref(props.column.name)
const selectedCardId = ref(0)
const selectedCardTitle = ref('')
const selectedCardDescription = ref('')
const internalIsModalOpen = ref(false);
const cards = ref<Card[]>([]);

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

const onDragEnd = (event: { to: HTMLElement; item: HTMLElement; newIndex: number }) => {
  const fromColumnId = props.column.id;
  const toColumnId = Number(event.to.dataset.columnId);
  const cardId = Number(event.item.dataset.cardId);
  const newIndex = event.newIndex;

  emits('dragCard', { fromColumnId, toColumnId, cardId, newIndex });
};

// Add a new card
const startAddCard = () => {
  isAddingCard.value = true
  newCardTitle.value = ''

  nextTick(() => {
    const inputElement = document.querySelector('.new-card-input') as HTMLInputElement;
    if (inputElement) {
      inputElement.focus()
    }
  })
}

const saveNewCard = async () => {
  if (newCardTitle.value.trim() === '') {
    console.error('Card title cannot be empty');
    return;
  }

  try {
    await projectStore.addCardToColumn(props.columnId, newCardTitle.value, props.projectId);
    cards.value = await projectStore.getCards(props.columnId); // Sincroniza la lista
  } catch (error) {
    console.error('Error saving new card:', error);
  } finally {
    isAddingCard.value = false;
  }
};

const cancelAddCard = () => {
  if (newCardTitle.value.trim() === '') {
    isAddingCard.value = false
  }
}

const openCard = (cardId: number, cardTitle: string, cardDescription: string) => {
  selectedCardId.value = cardId;
  selectedCardTitle.value = cardTitle;
  internalIsModalOpen.value = true;
  selectedCardDescription.value = cardDescription || ''
};

const closeCardDetail = () => {
  internalIsModalOpen.value = false;
  selectedCardId.value = 0;
  selectedCardTitle.value = '';
};

const updateCardTitle = async (newTitle: string) => {
  if (selectedCardId.value) {
    try {

      await projectStore.editCardInColumn(selectedCardId.value, newTitle);

      const refreshedCards = await projectStore.getCards(props.columnId);

      cards.value = refreshedCards;

      selectedCardTitle.value = newTitle;
    } catch (error) {
      console.error('Error updating card title:', error);
    }
  }
};

const updateCardDescription = async (newDescription: string) => {
  if (selectedCardId.value) {
    try {
      await projectStore.editCardInColumn(selectedCardId.value, selectedCardTitle.value, newDescription);
      selectedCardDescription.value = newDescription;

      // Sincronizar las tarjetas con el servidor
      const refreshedCards = await projectStore.getCards(props.columnId);
      cards.value = refreshedCards;
    } catch (error) {
      console.error('Error updating card description:', error);
    }
  }
};

const handleDeleteCard = async (projectId: number, columnId: number, cardId: number) => {
  try {
    await projectStore.deleteCardFromColumn(cardId)
    cards.value = cards.value.filter((card) => card.id !== cardId)
    internalIsModalOpen.value = false
  } catch (error) {
    console.error('Error deleting card:', error);
  }
}

const copyList = () => {
  const clonedColumn = {
    ...props.column,
    id: Date.now(),
    name: `${props.column.name} (Copy)`,
    cards: props.column.cards.map(card => ({ ...card, id: Date.now() + Math.random() }))
  }

  projectStore.addColumnToProject(props.projectId, clonedColumn.name)
  const newColumn = projectStore.projects.find(project => project.id === props.projectId)?.columns.find(column => column.name === clonedColumn.name)
  if (newColumn) {
    newColumn.cards = clonedColumn.cards

  }
}

const updateColumnColor = ({ columnId, color }: { columnId: number; color: string }) => {
  const column = projectStore.projects.find(project => project.id === props.projectId)?.columns.find(column => column.id === columnId)

  if (column) {
    column.color = color
  }
}

watch(() => props.column.name, (newName) => {
  editedColumnName.value = newName
})

watch(cards, (newValue) => {
  console.log('Cards updated in the frontend:', newValue);
});


onMounted(async () => {
  try {
    const loadedCards = await projectStore.getCards(props.columnId);
    console.log('Cards loaded from store:', loadedCards);
    cards.value = loadedCards;
  } catch (error) {
    console.error('Error loading cards:', error);
  }
});


</script>
