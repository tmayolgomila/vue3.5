import { defineStore } from 'pinia'
import { ref } from 'vue'

interface Project {
  id: number
  name: string
  description: string
}

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>([])
  const nextId = ref(1) //Generate ID's for new projects

  const addProject = (name: string, description: string) => {
    projects.value.push({
      id: nextId.value++,
      name,
      description,
    })
  }

  const updateProject = (id: number, updatedData: Partial<Project>) => {
    const project = projects.value.find((p) => p.id === id)
    if (project) {
      Object.assign(project, updatedData)
    }
  }

  const deleteProject = (id: number) => {
    projects.value = projects.value.filter((p) => p.id !== id)
  }

  return { addProject, updateProject, deleteProject }
})
