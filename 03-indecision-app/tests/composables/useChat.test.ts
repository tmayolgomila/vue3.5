import { useChat } from "@/composables/useChat"

describe('useChat', ()=>{
  test('add message correctly when onMessage is called', async()=>{
    const text = 'Hello World from useChat test'
    const {messages, onMessage} = useChat()
    await onMessage(text)
    expect(messages.value.length).toBe(1)
    expect(messages.value[0].itsMine).toBe(true)
    expect(messages.value[0].message).toBe(text)
  })
  test('return nothing is text is empty', async()=>{
    const text = ''
    const {messages, onMessage} = useChat()
    await onMessage(text)
    expect(messages.value.length).toBe(0)
  })
  test('gets the response when message ends with "?"', ()=>{
    const text = 'Do you want a coffe?'
    const {messages, onMessage} = useChat()
    await onMessage(text)
    await new Promise((r) => setTimeout(r, 2000))
  })
})
