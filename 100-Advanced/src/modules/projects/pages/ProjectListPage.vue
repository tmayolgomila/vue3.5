<template>
  <div>
    <h2 class="text-2xl mb-4">Project List</h2>

    <div v-if="projectStore.projects.length > 0">
      <ProjectCard v-for="project in projectStore.projects" :key="project.id" :project="project" />
    </div>
    <div v-else class="text-center mt-6">
      <p>No projects available. Add a new project to get started!</p>
      <RouterLink :to="{ name: 'projectNew' }">
        <button class="bg-blue-500 hover:bg-blue-600 text-white font-semibold rounded-md py-2 px-4 mt-4">
          Create New Project
        </button>
      </RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useProjectStore } from '../store/projectStore';
import ProjectCard from '../components/ProjectCard.vue';
import { onMounted } from 'vue';

const projectStore = useProjectStore();

onMounted(() => {
  projectStore.getProjects()
})

</script>
