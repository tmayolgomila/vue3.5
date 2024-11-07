import type { Pokemon } from "@/interfaces"

describe('Pkm interface', ()=>{

  const pkm:Pokemon = {
    id: 1,
    name: 'bulbasaur',
  }

  test('Should have an id property of type number', ()=> {
    expect(pkm.id).toEqual(expect.any(Number))
  })
  test('Should have a name property of type string', ()=> {
    expect(pkm.name).toEqual(expect.any(String))
    })
})
