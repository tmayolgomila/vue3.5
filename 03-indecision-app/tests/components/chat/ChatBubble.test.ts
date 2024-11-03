import ChatBubble from '@/components/chat/ChatBubble.vue'
import { mount } from '@vue/test-utils'

describe('<ChatBubble/>', () => {
  test('renders own message correctly', () => {
    const message = 'Hello World!'
    const wrapper = mount(ChatBubble, {
      props: {
        message,
        itsMine: true,
      },
    })
    //esperamos que el mío sea el mensaje azul
    expect(wrapper.find('.bg-blue-200').exists()).toBe(true)
    //también se puede poner así:
    expect(wrapper.find('.bg-blue-200').exists()).toBeTruthy()
    expect(wrapper.find('.bg-blue-200').text()).toContain(message)
    //esperamos que el mensaje gris no sea mío, así que esperamos que sea FALSE
    expect(wrapper.find('.bg-gray-300').exists()).toBeFalsy()
  })

  test('renders received message correctly', () => {
    const message = 'Hello World!'
    const wrapper = mount(ChatBubble, {
      props: {
        message,
        itsMine: false,
      },
    })
    //esperamos que el mensaje sea la respuesta al mío, por tanto que sea el GRIS
    expect(wrapper.find('.bg-gray-300').exists()).toBe(true)
    //por tanto no puede ser el azul
    expect(wrapper.find('.bg-blue-200').exists()).toBe(false)

    expect(wrapper.find('.bg-gray-300').text()).toContain(message)
    expect(wrapper.find('img').exists()).toBe(false) // porque en esta prueba no mandamos una imagen
  })

  //Vamos a hacer el test en el que sí mandemos una imagen:
  test('renders received message correctly with image', () => {
    const message = 'Hello World!'
    const image = 'example.jpg';
    const wrapper = mount(ChatBubble, {
      props: {
        message,
        itsMine: false,
        image, //No hace falta mandarle una img de verdad
      },
    })
    //esperamos que el mensaje sea la respuesta al mío, por tanto que sea el GRIS
    expect(wrapper.find('.bg-gray-300').exists()).toBe(true)
    //por tanto no puede ser el azul
    expect(wrapper.find('.bg-blue-200').exists()).toBe(false)

    expect(wrapper.find('.bg-gray-300').text()).toContain(message)
    expect(wrapper.find('img').exists()).toBe(true)
    expect(wrapper.find('img').attributes('src')).toBe(image)

  })
})
