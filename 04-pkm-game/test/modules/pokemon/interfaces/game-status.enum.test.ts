import { GameStatus } from "@/interfaces"

describe('GameStatus enum', ()=>{
  test('should have the value of playing', ()=>{
    expect(GameStatus.Playing).toBe('playing')
  })
  test('should have the value of won', ()=>{
    expect(GameStatus.Won).toBe('won')
  })
  test('should have the value of lost', ()=>{
    expect(GameStatus.Lost).toBe('lost')
  })
})

