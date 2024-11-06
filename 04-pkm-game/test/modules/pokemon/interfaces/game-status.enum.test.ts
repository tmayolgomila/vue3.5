import { GameStatus } from "@/interfaces"

describe('GameStatus enum', ()=>{
  test('should have the value of playing', ()=>{
    expect(GameStatus.Playing).toBe('playing')
  })
})

