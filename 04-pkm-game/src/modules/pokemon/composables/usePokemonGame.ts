import { GameStatus,type Pokemon, type PokemonListResponse } from "@/interfaces"
import { computed, onMounted, ref } from "vue"
import {pokemonApi} from "../api/pokemonApi"

export const usePokemonGame = () => {
//El juego va a tener 3 estados, cuando se gana, cuando se pierde y cuando estamos jugando
  const gameStatus = ref<GameStatus>(GameStatus.Playing)
  //Queremos saber si tenemos o no pokemons para saber si disparar el loader:
  const pokemons = ref<Pokemon[]>([])
  const pokemonOptions= ref<Pokemon[]>([])

  //vamos a escoger un pokemon del array de pokemonOptions y este será la respuesta:
  const randomPokemon = computed(()=>{
    const randomIndex = Math.floor(Math.random()*pokemonOptions.value.length)
    return pokemonOptions.value[randomIndex]
  })
  const isLoading = computed(()=> pokemons.value.length === 0)

  const getPokemons = async():Promise<Pokemon[]> => {
    const response = await pokemonApi.get<PokemonListResponse>('/?limit=151')

    const pokemonArray = response.data.results.map(pokemon =>{
      const urlParts = pokemon.url.split('/')
      const id = urlParts[urlParts.length - 2] ?? 0
      return{
        name: pokemon.name,
        id: +id, //con este + lo convierto de string a numero

      }
    })
    return pokemonArray.sort(() => Math.random() - 0.5)//Los devolvemos desordenados siempre, sino siempre saldran los 4 primeros
  }
  //Esta va a ser la función que empiece un nuevo juego
  const getNextOptions = (howMany:number=4) => {
    gameStatus.value = GameStatus.Playing
    pokemonOptions.value = pokemons.value.slice(0, howMany)//almaceno los primeros 4 (ya estan ordenados aleatoriamente)
    pokemons.value = pokemons.value.slice(howMany) //separo el resto que no he almacenado
  }



  onMounted(async()=>{
    await new Promise(r => setTimeout(r,300)) //esto lo usamos simplemente para simular la carga de pokemons, porque es tan rápida qu eno la tiene(Se puede borrar)

    pokemons.value = await getPokemons()
    getNextOptions()
    console.log(pokemonOptions.value)
  })


  return{
    gameStatus,
    isLoading,
    pokemonOptions,
    randomPokemon,

    //Methods
    getNextOptions,

  }
}
