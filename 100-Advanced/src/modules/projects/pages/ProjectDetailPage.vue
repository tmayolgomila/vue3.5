
<template>
  <div v-if="project">
    <h2 class="text-2xl font-semibold mb-4">Project: {{ project.name }}</h2>
    <div class="flex space-x-4">
      <ColumnDetail
        v-for="column in project.columns"
        :key="column.id"
        :column="column"
        @addCard="handleAddCard"
        @editCard="handleEditCard"
        @deleteCard="handleDeleteCard"
        @dragCard="handleDragCard"
        @editColumn="handleEditColumn"
        @deleteColumn="handleDeleteColumn"
      />
      <button @click="addColumn">Add New Column</button>
    </div>
  </div>
  <p v-else>No project found.</p>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import ColumnDetail from '../components/ColumnDetail.vue';
import { useProjectStore } from '../store/projectStore';
import { computed } from 'vue';

const route = useRoute()
const projectStore = useProjectStore()
const projectId = Number(route.params.id)

const project = computed(()=> projectStore.projects.find(p => p.id === projectId))

const addColumn = () => {
  const columnName = prompt('Enter column name')
  if(columnName){
    projectStore.addColumnToProject(projectId, columnName)
  }
}

const handleEditColumn = (columnId: number) => {
  const newName = prompt('Enter new column name')
  if(newName){
    projectStore.updateColumnName(projectId, columnId, newName)
  }
}

const handleDeleteColumn = (columnId: number) => {
  const confirmDelete = confirm('Are you sure you want to delete this column?')
  if(confirmDelete){
    projectStore.deleteColumn(projectId, columnId)
  }
}

const handleAddCard = (columnId: number) => {
  const cardTitle = prompt('Enter card title')
  const cardDescription = prompt('Enter card description')
  if(cardTitle && cardDescription){
    projectStore.addCardToColumn(projectId, columnId, cardTitle, cardDescription)
  }
}

const handleEditCard = ({columnId, cardId}: {columnId: number, cardId: number}) => {
  const cardTitle = prompt('Edit card title')
  const cardDescription = prompt('Edit card description')
  if(cardTitle && cardDescription){
    projectStore.editCardInColumn(projectId, columnId, cardId, cardTitle, cardDescription)
  }
}

const handleDeleteCard = ({columnId, cardId}: {columnId: number, cardId: number}) => {
  if(confirm('Are you sure you want to delete this card?')){
    projectStore.deleteCardFromColumn(projectId, columnId, cardId)
  }
}

interface DragCardParams {
  fromColumnId: number;
  toColumnId: number;
  cardId: number;
  newIndex: number;
}

const handleDragCard = ({ fromColumnId, toColumnId, cardId, newIndex } : DragCardParams) => {
  projectStore.moveCard(projectId, fromColumnId, toColumnId, cardId, newIndex);
};

</script>
