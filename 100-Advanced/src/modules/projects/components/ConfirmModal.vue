<template>

  <div class="fixed inset-0 flex items-center justify-center w-full h-full ">
    <div class="bg-white dark:bg-neutral-800 w-96 h-64 p-4 rounded shadow-lg relative z-60">

      <button @click="$emit('cancel')" class="absolute top-2 right-2 text-gray-500 hover:text-gray-700">
        <ion-icon size="large" name="close-outline"></ion-icon>
      </button>

      <div class="flex flex-col justify-between h-full">

        <div>
          <h2 class="text-lg font-bold mb-2">{{ title }}</h2>
          <p>{{ message }}</p>
        </div>

        <div class="flex justify-between mt-6">

          <button class="text-red-700 text-lg font-bold" @click="$emit('confirm')">
            {{ confirmAction }}
          </button>
          <button class="text-gray-300 text-lg font-bold" @click="$emit('cancel')">
            Cancel
          </button>

        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">

defineProps({
  title: {
    type: String,
    required: true
  },
  message: {
    type: String,
    required: true
  },
  confirmAction: {
    type: String,
    required: true
  }
})

</script>
