<template>
  <section v-if="isLoading || randomPokemon.id === null"
    class="flex flex-col justify-center items-center w-screen h-screen">
    <span class="loader"></span>
  </section>
  <section v-else class="flex flex-col justify-center items-center w-screen h-screen">
    <h1 class="m-5">Who is that?</h1>

    <div class="h-20">
      <button
      v-if="gameStatus!== GameStatus.Playing"
      @click="() => getNextRound()"
      class="bg-blue-500 text-white p-2 rounded-md hover:bg-blue-700 transition-all"
      >Next game</button>
    </div>


    <PokemonPicture :pokemon-id="randomPokemon.id" :show-pokemon="gameStatus !== GameStatus.Playing" />

    <PokemonOptions
    :options="pokemonOptions"
    :block-selection="gameStatus !== GameStatus.Playing"
    :correct-answer="randomPokemon.id"
    @selected-option="checkAnswer" />

  </section>

</template>

<script setup lang="ts">
import PokemonOptions from '@/components/PokemonOptions.vue';
import PokemonPicture from '@/components/PokemonPicture.vue';
import { usePokemonGame } from '../composables/usePokemonGame'
import { GameStatus } from '@/interfaces';

const { randomPokemon, isLoading, gameStatus, pokemonOptions, checkAnswer, getNextRound} = usePokemonGame()


</script>

<style scoped>
.loader {
  width: 48px;
  height: 48px;
  border: 5px solid #7ad4f0;
  border-bottom-color: transparent;
  border-radius: 50%;
  display: inline-block;
  box-sizing: border-box;
  animation: rotation 1s linear infinite;
}

@keyframes rotation {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}
</style>
