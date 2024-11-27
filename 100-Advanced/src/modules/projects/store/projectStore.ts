import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface Card {
  id: number
  title: string
  description: string
  position: number
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
  const projects = ref<Project[]>([])

  // Projects start
  const getProjects = async () => {
    try {
      const response = await fetch('http://localhost:8080/projects')
      if (!response.ok) throw new Error('Failed to fetch projects')
      projects.value = await response.json()
    } catch (error) {
      console.error('Error fetching projects:', error)
    }
  }

  const addProject = async (name: string, description: string) => {
    try {
      const response = await fetch('http://localhost:8080/projects', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ name, description }),
      })
      if (!response.ok) throw new Error('Failed to add project')
      const newProject = await response.json()
      projects.value.push(newProject)
    } catch (error) {
      console.error('Error adding project:', error)
    }
  }

  const updateProject = async (id: number, updatedData: Partial<Project>) => {
    try {
      const response = await fetch(`http://localhost:8080/projects?id=${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(updatedData),
      })
      if (!response.ok) throw new Error('Failed to update project')
      const projectIndex = projects.value.findIndex((p) => p.id === id)
      if (projectIndex !== -1) {
        Object.assign(projects.value[projectIndex], updatedData)
      }
    } catch (error) {
      console.error('Error updating project:', error)
    }
  }

  const deleteProject = async (id: number) => {
    try {
      const response = await fetch(`http://localhost:8080/projects?id=${id}`, { method: 'DELETE' })
      if (!response.ok) throw new Error('Failed to delete project')
      projects.value = projects.value.filter((p) => p.id !== id)
    } catch (error) {
      console.error('Error deleting project:', error)
    }
  }
  //Projects end

  //Columns start
  const getColumns = async (projectId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/columns?project_id=${projectId}`)
      if (!response.ok) throw new Error('Failed to fetch columns')
      const columns = await response.json()
      const project = projects.value.find((p) => p.id === projectId)
      if (project) project.columns = columns
    } catch (error) {
      console.error('Error fetching columns:', error)
    }
  }

  const addColumnToProject = async (projectId: number, columnName: string) => {
    try {
      const response = await fetch('http://localhost:8080/columns', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ project_id: projectId, name: columnName }),
      })
      if (!response.ok) throw new Error('Failed to add column')
      const newColumn = await response.json()
      const project = projects.value.find((p) => p.id === projectId)
      if (project) project.columns.push({ ...newColumn, cards: [] })
    } catch (error) {
      console.error('Error adding column to project:', error)
    }
  }

  const updateColumnName = async (columnId: number, newName: string) => {
    try {
      const response = await fetch(`http://localhost:8080/columns?id=${columnId}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ id: columnId, name: newName }),
      })
      if (!response.ok) throw new Error('Failed to update column name')
      projects.value.forEach((project) => {
        const column = project.columns.find((c) => c.id === columnId)
        if (column) column.name = newName
      })
    } catch (error) {
      console.error('Error updating column name:', error)
    }
  }

  const deleteColumn = async (columnId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/columns?id=${columnId}`, {
        method: 'DELETE',
      })
      if (!response.ok) throw new Error('Failed to delete column')
      projects.value.forEach((project) => {
        project.columns = project.columns.filter((c) => c.id !== columnId)
      })
    } catch (error) {
      console.error('Error deleting column:', error)
    }
  }
  //Columns end

  //Cards start
  const getCards = async (columnId: number): Promise<Card[]> => {
    try {
      const response = await fetch(`http://localhost:8080/cards?column_id=${columnId}`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
      })

      if (!response.ok) {
        throw new Error(`Failed to fetch cards for column ${columnId}: ${response.statusText}`)
      }

      const cards: Card[] = await response.json()

      projects.value.forEach((project) => {
        const column = project.columns.find((c) => c.id === columnId)
        if (column) {
          column.cards = Array.isArray(cards) ? cards : []
        }
      })

      return cards
    } catch (error) {
      console.error('Error fetching cards:', error)
      return []
    }
  }

  const addCardToColumn = async (
    columnId: number,
    cardTitle: string,
    projectId: number,
    description = '',
  ) => {
    try {
      const response = await fetch('http://localhost:8080/cards', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ column_id: columnId, title: cardTitle, description }),
      })
      if (!response.ok) throw new Error('Failed to add card')

      const newCard = await response.json()

      const project = projects.value.find((project) => project.id === projectId)
      if (!project) {
        console.error(`Project with ID ${projectId} not found`)
        return
      }

      const column = project.columns.find((c) => c.id == columnId)

      if (column) {
        if (!Array.isArray(column.cards)) {
          console.warn(`Initializing cards for column ${column.id}`)
          column.cards = []
        }
        column.cards.push(newCard)
      } else {
        console.error(`Column with ID ${columnId} not found in project`, project)
      }
    } catch (error) {
      console.error('Error adding card to column:', error)
    }
  }

  const moveCard = async (
    fromColumnId: number,
    cardId: number,
    toColumnId: number,
    newPosition: number,
    oldPosition: number,
  ) => {
    console.log('Sending data to backend:', {
      fromColumnId,
      cardId,
      toColumnId,
      newPosition,
      oldPosition,
    })

    try {
      const response = await fetch('http://localhost:8080/cards/move', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          fromColumnId,
          cardId,
          toColumnId,
          newPosition,
          oldPosition,
        }),
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.message || 'Failed to move card')
      }

      console.log('Card moved successfully')
    } catch (error) {
      console.error('Error moving card:', error)
    }
  }

  const editCardInColumn = async (cardId: number, title?: string, description?: string) => {
    try {
      const requestBody: { title?: string; description?: string } = {}
      if (title !== undefined) requestBody.title = title
      if (description !== undefined) requestBody.description = description

      const response = await fetch(`http://localhost:8080/cards?id=${cardId}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(requestBody),
      })

      if (!response.ok) throw new Error('Failed to update card')
    } catch (error) {
      console.error('Error updating card:', error)
    }
  }

  const deleteCardFromColumn = async (cardId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/cards?id=${cardId}`, {
        method: 'DELETE',
      })
      if (!response.ok) throw new Error('Failed to delete card')
    } catch (error) {
      console.error(error)
    }
  }

  const sortCardsInColumn = async (
    columnId: number,
    sortBy: 'title' | 'id',
    order: 'asc' | 'desc',
  ) => {
    try {
      const response = await fetch(
        `http://localhost:8080/cards?column_id=${columnId}&sortBy=${sortBy}&order=${order}`,
      )
      if (!response.ok) {
        throw new Error('Failed to sort cards')
      }
      const sortedCards = await response.json()
      console.log(sortedCards)
    } catch (error) {
      console.error(error)
    }
  }

  const syncCardPositions = async (projectId: number, columnId: number) => {
    
    const project = projects.value.find((p) => p.id === projectId);
    if (!project) {
      console.error(`Project with ID ${projectId} not found`);
      return;
    }
  
    const column = project.columns.find((c) => c.id === columnId);
    if (!column) {
      console.error(`Column with ID ${columnId} not found in project ${projectId}`);
      return;
    }
  
    const updatedPositions = column.cards.map((card, index) => ({
      id: card.id,
      position: index,
    }));
  
    try {
      const response = await fetch('http://localhost:8080/updateCardPositions', {
        method: 'POST',
        body: JSON.stringify(updatedPositions),
        headers: { 'Content-Type': 'application/json' },
      });
  
      if (!response.ok) {
        throw new Error(`Failed to update card positions: ${response.statusText}`);
      }
  
      console.log('Card positions updated successfully');
    } catch (error) {
      console.error('Error updating card positions:', error);
    }
  };
  

  return {
    projects,

    getProjects,
    addProject,
    updateProject,
    deleteProject,

    getColumns,
    addColumnToProject,
    updateColumnName,
    deleteColumn,

    getCards,
    addCardToColumn,
    editCardInColumn,
    deleteCardFromColumn,
    moveCard,

    sortCardsInColumn,
    syncCardPositions,
  }
})
