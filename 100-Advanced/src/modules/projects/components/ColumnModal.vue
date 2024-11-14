<template>
  <div class="bg-white dark:bg-slate-600 p-6 rounded shadow-md w-64 absolute"
    :style="{ top: `${position.top}px`, left: `${position.left}px`, position: 'absolute', zIndex: 50 }">

    <button @click="handleDelete" class="text-red-500">Remove List<ion-icon
        name="close-circle-outline"></ion-icon></button>
    <button @click="$emit('close')">Cancel <ion-icon name="close-outline"></ion-icon></button>
  </div>
</template>
<script setup lang="ts">

const props = defineProps({
  columnId: {
    type: Number,
    required: true,
  },
  position: {
    type: Object,
    required: true,
    validator: (val: any) => 'top' in val && 'left' in val,
  }
})

const emits = defineEmits(['deleteColumn', 'close'])

const handleDelete = () => {
  emits('deleteColumn', props.columnId)
}

</script>
