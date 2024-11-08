<template>
  <div class="max-w-md mx-auto mt-6">
    <h2 class="text-2xl font-semibold mb-4">
      {{ isEditing ? 'Edit Project' : 'Create New Project' }}
    </h2>
    <form @submit.prevent="handleSubmit">
      <div class="mb-4">
        <label for="name" class="block text-gray-600">Project Name</label>
        <input
          v-model="projectData.name"
          type="text"
          id="name"
          class="w-full border rounded-md py-2 px-3 focus:outline-none focus:border-blue-500"
          required
        />
      </div>
      <div class="mb-4">
        <label for="description" class="block text-gray-600">Description</label>
        <textarea
          v-model="projectData.description"
          id="description"
          class="w-full border rounded-md py-2 px-3 focus:outline-none focus:border-blue-500"
          rows="3"
          required
        ></textarea>
      </div>
      <button
        type="submit"
        class="bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-md py-2 px-4 w-full"
      >
        {{ isEditing ? 'Update Project' : 'Create Project' }}
      </button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useProjectStore } from '../store/projectStore'
import { useRouter, useRoute } from 'vue-router'

const projectStore = useProjectStore()
const router = useRouter()
const route = useRoute()

const isEditing = computed(() => !!route.params.id)
const projectData = ref({
  name: '',
  description: '',
})

if (isEditing.value) {
  const projectId = parseInt(route.params.id as string, 10)
  const project = projectStore.projects.find((p) => p.id === projectId)
  if (project) {
    projectData.value = { name: project.name, description: project.description }
  }
}

const handleSubmit = () => {
  if (isEditing.value) {
    const projectId = parseInt(route.params.id as string, 10)
    projectStore.updateProject(projectId, projectData.value)
  } else {
    projectStore.addProject(projectData.value.name, projectData.value.description)
  }
  router.push({ name: 'projectList' })
}
</script>
