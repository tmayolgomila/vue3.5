import { defineStore } from 'pinia'
import { ref, watch } from 'vue'

interface Card {
  id: number
  title: string
  description: string
}
export interface Column {
  id: number
  name: string
  cards: Card[]
  color?: string
}
interface Project {
  id: number
  name: string
  description: string
  columns: Column[]
}

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>(JSON.parse(localStorage.getItem('projects') || '[]'))
  const nextId = ref(projects.value.length ? Math.max(...projects.value.map((p) => p.id)) + 1 : 1)

  const saveProjectsToLocalStorage = () => {
    localStorage.setItem('projects', JSON.stringify(projects.value))
  }

  const addProject = (name: string, description: string) => {
    projects.value.push({
      id: nextId.value++,
      name,
      description,
      columns: [],
    })
    saveProjectsToLocalStorage()
  }

  const updateProject = (id: number, updatedData: Partial<Project>) => {
    const project = projects.value.find((p) => p.id === id)
    if (project) {
      Object.assign(project, updatedData)
      saveProjectsToLocalStorage()
    }
  }

  const deleteProject = (id: number) => {
    projects.value = projects.value.filter((p) => p.id !== id)
    saveProjectsToLocalStorage()
  }

  const addColumnToProject = (projectId: number, columnName: string, cards: Card[] = []) => {
    const project = projects.value.find((p) => p.id === projectId)
    if (project) {
      const newColumn: Column = { id: Date.now(), name: columnName, cards }
      project.columns.push(newColumn)
      saveProjectsToLocalStorage()
    }
  }

  const addCardToColumn = (projectId: number, columnId: number, cardTitle: string) => {
    const project = projects.value.find((p) => p.id === projectId)
    const column = project?.columns.find((c) => c.id === columnId)
    if (column) {
      const newCard: Card = { id: Date.now(), title: cardTitle, description: '' }
      column.cards.push(newCard)
      saveProjectsToLocalStorage()
    }
  }

  const editCardInColumn = (projectId: number, columnId: number, cardId: number, title: string) => {
    const project = projects.value.find((p) => p.id === projectId)
    const column = project?.columns.find((c) => c.id === columnId)
    const card = column?.cards.find((card) => card.id === cardId)
    if (card) {
      card.title = title
      saveProjectsToLocalStorage()
    }
  }

  const setCardDescription = (
    projectId: number,
    columnId: number,
    cardId: number,
    description: string,
  ) => {
    const project = projects.value.find((p) => p.id === projectId)
    const column = project?.columns.find((c) => c.id === columnId)
    const card = column?.cards.find((card) => card.id === cardId)
    if (card) {
      card.description = description
      saveProjectsToLocalStorage()
    }
  }

  const deleteCardFromColumn = (projectId: number, columnId: number, cardId: number) => {
    const project = projects.value.find((p) => p.id === projectId)
    const column = project?.columns.find((c) => c.id === columnId)
    if (column) {
      column.cards = column.cards.filter((card) => card.id !== cardId)
      saveProjectsToLocalStorage()
    }
  }

  const moveCard = (
    projectId: number,
    fromColumnId: number,
    toColumnId: number,
    cardId: number,
    mewIndex: number,
  ) => {
    const project = projects.value.find((p) => p.id === projectId)
    const fromColumn = project?.columns.find((c) => c.id === fromColumnId)
    const toColumn = project?.columns.find((c) => c.id === toColumnId)

    if (fromColumn && toColumn) {
      const cardIndex = fromColumn.cards.findIndex((card) => card.id === cardId)
      if (cardIndex !== -1) {
        const [movedCard] = fromColumn.cards.splice(cardIndex, 1)
        toColumn.cards.splice(mewIndex, 0, movedCard)
        saveProjectsToLocalStorage()
      }
    }
  }

  const updateColumnName = (projectId: number, columnId: number, newName: string) => {
    const project = projects.value.find((p) => p.id === projectId)
    const column = project?.columns.find((c) => c.id === columnId)
    if (column) {
      column.name = newName
      saveProjectsToLocalStorage()
    }
  }

  const deleteColumn = (projectId: number, columnId: number) => {
    const project = projects.value.find((p) => p.id === projectId)
    if (project) {
      project.columns = project.columns.filter((c) => c.id !== columnId)
      saveProjectsToLocalStorage()
    }
  }

  watch(projects, saveProjectsToLocalStorage, { deep: true })

  return {
    projects,
    addProject,
    updateProject,
    deleteProject,
    addColumnToProject,
    addCardToColumn,
    editCardInColumn,
    deleteCardFromColumn,
    moveCard,
    updateColumnName,
    deleteColumn,
    setCardDescription,
    saveProjectsToLocalStorage,
  }
})
