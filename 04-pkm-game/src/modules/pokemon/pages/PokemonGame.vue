<template>
  <section v-if="isLoading" class="flex flex-col justify-center items-center w-screen h-screen">
    <span class="loader"></span>
  </section>
  <section v-else class="flex flex-col justify-center items-center w-screen h-screen">
    <h1 class="m-5">Who is that?</h1>
    <h3>{{ randomPokemon }}</h3>

  <PokemonPicture :pokemon-id="randomPokemon.id"/>

  <PokemonOptions/>

  </section>

</template>

<script setup lang="ts">
import PokemonOptions from '@/components/PokemonOptions.vue';
import PokemonPicture from '@/components/PokemonPicture.vue';
import {usePokemonGame} from '../composables/usePokemonGame'

const {randomPokemon, isLoading} = usePokemonGame()

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
