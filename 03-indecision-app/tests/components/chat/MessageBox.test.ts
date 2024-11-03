import MessageBox from "@/components/chat/MessageBox.vue"
import { mount } from "@vue/test-utils"


describe('<MessageBox/>', ()=>{

  const wrapper = mount(MessageBox);// Hemos pasado el wrapper fuera de los test para usarlo en todos y no tener que llamarlo cada vez

  test('renders input and button elements correctly', ()=>{
    expect(wrapper.html()).toMatchSnapshot() // esto no evalua el render del input ni del button, solo 'toma la foto del código'
    expect(wrapper.find('input[type="text"]').exists()).toBe(true)//Podríamos buscar solo por el input pero así nos aseguramos de escoger el texto de dentro
    expect(wrapper.find('button').exists()).toBe(true)//Nos aseguramos de que el botón exista
    //Queremos asegurarnos de que el bot-on tiene el icono:
    expect(wrapper.find('button svg').exists()).toBe(true)
  })

  test('emits sendMessage event when button is clicked with message value', async()=>{
    const message = 'Hello World'

    await wrapper.find('input[type="text"]').setValue(message)//seteamos un mensaje en el input para enviarlo
    await wrapper.find('button').trigger('click')//simulación del click sobre el botón de enviar
    expect(wrapper.emitted('sendMessage')?.[0]).toEqual([message])// Aquí toBe no va a funcionar, por eso usamos toEqual, no es lo mismo = que ===
    //expect((wrapper.vm as any).message).toBe('')// vm es ViewModel (wrapper.vm as any).message accede a la propiedad message del componente montado.
  })

  test('emits sendMessage event when keypress.enter is triggered with message value', async()=>{
    const message = 'Hello World'
    const input = wrapper.find('input')
    await input.setValue(message)
    await input.trigger('keypress.enter')
    expect(wrapper.emitted('sendMessage')?.[0]).toEqual([message])// Aquí toBe no va a funcionar, por eso usamos toEqual, no es lo mismo = que ===
  })

  test('should not emit a message with an empty input', async()=>{
    const wrapper = mount(MessageBox);//Aquí si que declaramos el wrapper porque sino las pruebas anteriores nos lo rellenan con el setValue
    const input = wrapper.find('input')
    await input.trigger('keypress.enter')
    expect(wrapper.emitted('sendMessage')).toBeFalsy()
  })
})



