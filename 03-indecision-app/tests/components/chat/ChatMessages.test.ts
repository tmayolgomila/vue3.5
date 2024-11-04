

import ChatMessages from "@/components/chat/ChatMessages.vue"
import type { ChatMessage } from "@/interfaces/chat-message.interface"
import { mount } from "@vue/test-utils"

const messages: ChatMessage[] = [
  {id: 1, message: 'Hello', itsMine: true},
  {id:2, message: 'World', itsMine: false, image: 'quedisecabesa.jpg'}
]

describe('<ChatMessages/>', ()=>{
  const wrapper = mount(ChatMessages, {
    props:{
      messages //declaramos arriba un array de mensajes importando la interface de ChatMessages para el test
    }
  })
  test('renders chat messages correctly', ()=>{
    const chatBubbles = wrapper.findAllComponents({name: 'ChatBubble'}) //Quiero que busque mi componente que tiene ese nombre
    expect(chatBubbles.length).toBe(messages.length)//deben de ser el mismo numero que los mensajes que creamos l principio
  })

  test('scrolls down to the bottom after messages update', async() => {
    const scrollToMock = vi.fn() //apunta al .scrollToHeight del watch
    const chatRef = wrapper.vm.$refs.chatRef as HTMLDivElement
    chatRef.scrollTo = scrollToMock

    await wrapper.setProps({
      messages: [...messages, {id: 3, message: 'Hey', itsMine: true}]
    });

    await new Promise((r) => setTimeout(r, 150))

    expect(scrollToMock).toHaveBeenCalledTimes(1)
    expect(scrollToMock).toHaveBeenCalledWith({
      behavior: 'smooth',
      top: expect.any(Number),
    })

  })

})
