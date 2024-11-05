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
  test('gets the response when message ends with "?"', async()=>{
    const text = 'Do you want a coffe?'
    const {messages, onMessage} = useChat()
    await onMessage(text)
    await new Promise((r) => setTimeout(r, 2000))

    const [myMessage, herMessage] = messages.value

    expect(messages.value.length).toBe(2)
    expect(herMessage).toEqual({
      id:expect.any(Number),
      image:expect.any(String),
      message: expect.any(String),
      itsMine: false
    })
    expect(myMessage).toEqual({
      id: expect.any(Number),
      itsMine: true,
      message: text
    })
  })
  test('mock response - fetch api', async () => {
    const mockResponse = { answer: 'yes', image: 'example.gif' };

    // Mock del fetch global para devolver la respuesta simulada
    (window as any).fetch = vi.fn(async () => ({
      json: async () => mockResponse,
    }));

    const text = 'Do you want a coffee?';
    const { messages, onMessage } = useChat();

    // Llama a onMessage con una pregunta para forzar la respuesta
    await onMessage(text);
    await new Promise((r) => setTimeout(r, 1600));

    const [, herMessage] = messages.value;

    // Verifica que el mensaje de la respuesta sea correcto
    expect(herMessage).toEqual({
      id: expect.any(Number),
      itsMine: false,
      message: mockResponse.answer, // Aquí es la respuesta simulada
      image: mockResponse.image, // Aquí es la imagen simulada
    });
  });


})
