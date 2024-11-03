<template>
  <section>
    <h3>Counter: {{ counter }}</h3>
    <h3 data-testid="square-label">Square: {{ squareCounter }}</h3>

    <div>
      <button class="btn" @click="counter++">+1</button>
      <button data-testid="decrease-button" class="btn" @click="counter--">-1</button>
    </div>
  </section>
</template>

<script lang="ts" setup>
//We take the logic to the composable useCounter and then import it here:
import { useCounter } from '@/composables/useCounter';
interface Props {
  value: number;
}
const props = defineProps<Props>()

const { counter, squareCounter } = useCounter(props.value)

</script>

<style scoped>
.btn {
  @apply p-3 bg-blue-500 rounded ml-1 w-10 hover:bg-blue-700
}
</style>
