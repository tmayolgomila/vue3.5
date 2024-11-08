import { defineStore } from 'pinia'
import { ref, watch} from 'vue'

interface Project {
  id: number
  name: string
  description: string
}

export const useProjectStore = defineStore('project', () => {
  const projects = ref<Project[]>(JSON.parse(localStorage.getItem('projects') || '[]'))
  const nextId = ref(projects.value.length ? Math.max(...projects.value.map(p=>p.id)) + 1 : 1)

  const saveProjectsToLocalStorage = ()=>{
    localStorage.setItem('projects', JSON.stringify(projects.value))
  }

  const addProject = (name: string, description: string) => {
    projects.value.push({
      id: nextId.value++,
      name,
      description,
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

  watch(projects, saveProjectsToLocalStorage, { deep: true })

  return {projects, addProject, updateProject, deleteProject }
})
