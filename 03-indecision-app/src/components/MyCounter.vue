<template>
  <section>
    <h3>Counter: {{ counter }}</h3>
    <h3>Square: {{ squareCounter }}</h3>

    <div>
      <button class="btn" @click="counter++">+1</button>
      <button class="btn" @click="counter--">-1</button>
    </div>
  </section>
</template>

<script lang="ts" setup>
//We take the logic to the composable useCounter and then import it here:
import { useCounter } from '@/composables/useCounter';

const { counter, squareCounter } = useCounter(10)

//import { computed, ref } from 'vue';
//
//interface Props {
//  value: number;
//}
//
//const props = defineProps<Props>()

//const counter = ref(props.value)
//const squareCounter = computed(() => counter.value * counter.value)

</script>

<style scoped>
.btn {
  @apply p-3 bg-blue-500 rounded ml-1 w-10 hover:bg-blue-700
}
</style>
