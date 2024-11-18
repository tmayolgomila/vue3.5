<template>

  <div v-if="project" class="h-full">
    <h2 class="text-2xl font-semibold mb-4">Project: {{ project.name }}</h2>
    <div class="flex space-x-4">
      <!-- Contenedor de columnas y botón -->
      <div class="flex space-x-4 items-start">
        <!-- Columnas -->
        <div class="flex space-x-4">

          <ColumnDetail v-for="column in project.columns" :key="column.id" :column="column" :columnId="column.id"
            :projectId="projectId" @editCard="handleEditCard" @deleteCard="handleDeleteCard" @dragCard="handleDragCard"
            @editColumn="handleEditColumn" @deleteColumn="handleDeleteColumn" />

        </div>

        <!-- Botón Add New Column alineado en la parte superior derecha -->
        <div class="flex items-center">
          <div v-if="isAddingColumn" class="relative flex items-center">
            <input type="text" v-model="newColumnName" @blur="handleBlur" @keyup.enter="saveNewColumn"
              placeholder="Enter column title" class="text-black border rounded px-2 py-1 new-column-input" autofocus />
            <button @click="saveNewColumn" class="ml-2 bg-blue-500 text-white font-bold py-1 px-3 rounded">
              Save
            </button>

          </div>
          <button v-else @click="startNewColumn"
            class="flex items-center space-x-1 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full h-12">
            <ion-icon name="add-outline"></ion-icon>
            <span>Add New Column</span>
          </button>
        </div>

      </div>
    </div>
  </div>

  <p v-else>No project found.</p>

</template>

<script setup lang="ts">

import { useRoute } from 'vue-router';
import ColumnDetail from '../components/ColumnDetail.vue';
import { useProjectStore } from '../store/projectStore';
import { computed, nextTick, ref } from 'vue';

const route = useRoute()
const projectStore = useProjectStore()
const projectId = Number(route.params.id)

const project = computed(() => projectStore.projects.find(p => p.id === projectId))

const isAddingColumn = ref(false)
const newColumnName = ref('')

const startNewColumn = () => {
  isAddingColumn.value = true
  newColumnName.value = ''

  nextTick(() => {
    const inputElement = document.querySelector('.new-column-input') as HTMLInputElement;
    if (inputElement) {
      inputElement.focus()
    }
  })
}

const saveNewColumn = () => {
  if (newColumnName.value.trim() !== '') {
    projectStore.addColumnToProject(projectId, newColumnName.value.trim())
  }
  resetNewColumnState()
}

const handleBlur = () => {
  if (!newColumnName.value.trim()) {
    resetNewColumnState();
  }
};

const resetNewColumnState = () => {
  isAddingColumn.value = false;
  newColumnName.value = '';
};

const handleEditColumn = ({ columnId, newName }: { columnId: number; newName: string }) => {
  projectStore.updateColumnName(projectId, columnId, newName);
};


const handleDeleteColumn = (columnId: number) => {
  const confirmDelete = confirm('Are you sure you want to delete this column?')
  if (confirmDelete) {
    projectStore.deleteColumn(projectId, columnId)
  }
}

const handleEditCard = ({ columnId, cardId }: { columnId: number, cardId: number }) => {
  const cardTitle = prompt('Edit card title')
  if (cardTitle) {
    projectStore.editCardInColumn(projectId, columnId, cardId, cardTitle)
  }
}

const handleDeleteCard = ({ columnId, cardId }: { columnId: number, cardId: number }) => {
  if (confirm('Are you sure you want to delete this card?')) {
    projectStore.deleteCardFromColumn(projectId, columnId, cardId)
  }
}

interface DragCardParams {
  fromColumnId: number;
  toColumnId: number;
  cardId: number;
  newIndex: number;
}

const handleDragCard = ({ fromColumnId, toColumnId, cardId, newIndex }: DragCardParams) => {
  projectStore.moveCard(projectId, fromColumnId, toColumnId, cardId, newIndex);
};


</script>
