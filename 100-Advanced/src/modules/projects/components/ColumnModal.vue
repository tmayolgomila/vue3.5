<template>
  <div class="absolute bg-white dark:bg-slate-800 p-6 rounded shadow-md w-80 column"
    :style="{ top: `${position.top}px`, left: `${position.left}px`, position: 'absolute', zIndex: 50 }">


    <button @click="$emit('close')" class="absolute top-2 right-2"><ion-icon size="large"
        name="close-outline"></ion-icon></button>

    <div class=" flex flex-col items-start space-y-2">

      <button>Add Card</button>
      <button>Copy Column</button>
      <button>Move List</button>
      <button>Move all card list to another</button>
      <button>Sort by</button>
      <button>Follow</button>
      <button>Change column color</button>
      <button>Remove color</button>
      <button @click="handleDelete" class="text-red-500">Remove List</button>
    </div>

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
