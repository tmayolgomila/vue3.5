<template>
  <div class="border p-4 rounded shadow">
    <h3 class="text-xl font-semibold">{{ project.name }}</h3>
    <p class="text-gray-600">{{ project.description }}</p>
    <RouterLink :to="{ name: 'projectDetail', params: { id: project.id } }" class="text-blue-500 mt-2 inline-block">View
      Details</RouterLink>
    <RouterLink :to="{ name: 'projectEdit', params: { id: project.id } }" class="text-blue-500 mt-2 inline-block">Edit
      Project</RouterLink>
  </div>
</template>


<script setup lang="ts">

defineProps({
  project: {
    type: Object,
    required: true
  }
})

</script>
