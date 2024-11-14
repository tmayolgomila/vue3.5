<template>
  <div class="absolute bg-white dark:bg-slate-800 p-6 rounded shadow-md w-80"
    :style="{ top: `${position.top}px`, left: `${position.left}px`, position: 'absolute', zIndex: 50 }">

    <button @click="handleDelete" class="text-red-500">Remove List</button>
    <button @click="$emit('close')">Cancel</button>
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
