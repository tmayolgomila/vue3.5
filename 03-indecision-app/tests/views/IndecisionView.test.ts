import ChatMessages from "@/components/chat/ChatMessages.vue"
import MessageBox from "@/components/chat/MessageBox.vue"
import IndecisionView from "@/views/IndecisionView.vue"
import { mount } from "@vue/test-utils"

describe('<IndecisionView />', ()=>{
  test('renders chat messages and messagebox correctly', ()=>{
    const wrapper = mount(IndecisionView)

    expect(wrapper.html()).toMatchSnapshot()//este el día de mañana va a fallar porque es la foto fija
    //es el que ha generado el archivo y carpeta snapshots(la foto)

    expect(wrapper.findComponent(ChatMessages).exists()).toBe(true)
    expect(wrapper.findComponent(MessageBox).exists()).toBe(true)
  })

  test('calls onMessage when sending a message', async()=>{
    const wrapper = mount(IndecisionView)
    //Simular el evento:
    const messageBoxComponent = wrapper.findComponent(MessageBox)
    messageBoxComponent.vm.$emit('sendMessage', 'Hello World from indecisionView test')
    await new Promise ((r)=> setTimeout(r, 150))
  })

})

