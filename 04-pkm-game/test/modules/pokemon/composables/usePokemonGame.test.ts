import { usePokemonGame } from "@/modules/pokemon/composables/usePokemonGame"
import { withSetup } from "../../../utils/with-setup"
import { GameStatus } from "@/interfaces"
import { flushPromises } from "@vue/test-utils"
import MockAdapter from "axios-mock-adapter"
import { pokemonApi } from "@/modules/pokemon/api/pokemonApi"
import { pokemonListFake } from "../../../data/fake-pokemons"

const mockPokemonApi = new MockAdapter(pokemonApi)

mockPokemonApi.onGet('/?limit=151').reply(200, {
  results: pokemonListFake
})

describe('usePokemonGame', ()=> {
  test('should initialize with the correct default values', async()=> {
    const [results, app] = withSetup(usePokemonGame)

    expect(results.gameStatus.value).toBe(GameStatus.Playing)
    expect(results.isLoading.value).toBe(true)
    expect(results.pokemonOptions.value).toEqual([])
    expect(results.randomPokemon.value).toBe(undefined)

    await flushPromises()//se asegura de que todas las promesas pendientes terminen antes de continuar
    expect(results.isLoading.value).toBe(false)
    expect(results.pokemonOptions.value.length).toBe(4)
    expect(results.randomPokemon.value).toEqual({
      id: expect.any(Number),
      name: expect.any(String)
    })
  });

  test('should correctly handle getNextRound', async() => {
    const [results] = withSetup(usePokemonGame)
    await flushPromises()

    results.gameStatus.value = GameStatus.Won
    results.getNextRound(4)

    expect(results.gameStatus.value).toBe(GameStatus.Playing)
    expect(results.pokemonOptions.value).toHaveLength(4)
  })


})

