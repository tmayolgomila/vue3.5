import {describe, test} from 'vitest'
import {mount} from '@vue/test-utils'
import MyCounter from '@/components/MyCounter.vue'

describe('<MyCounter/>', () => {

  test('should match snapshot', ()=>{
    const wrapper = mount(MyCounter, {
      props: {
        value: 5,
      },
    })
    expect(wrapper.html()).toMatchSnapshot()
  });

  test('renders the counter value correctly ', () => {
    const value = 5
    const wrapper = mount(MyCounter, {
      props:{
        value: value,
      }
    });
    expect(wrapper.find('h3').text()).toContain(`Counter: ${value}`)
    expect(wrapper.find('[data-testid="square-label"]').text()).toContain(`Square: ${value*value}`)

    const [counterLabel, squareLabel] = wrapper.findAll('h3');

    expect(counterLabel.text()).toContain(`Counter: ${value}`)
    expect(squareLabel.text()).toContain(`Square: ${value*value}`)

    //console.log(wrapper.html())
  });

  test('increments the counter when +1 button is pressed', async()=>{

    const value = 5
    const wrapper = mount(MyCounter, {
      props:{
        value: value,
      }
    });

    const [counterLabel, squareLabel] = wrapper.findAll('h3');

    const btnIncrement = wrapper.find('button');
    //simulamos un evento click,
    //se aconseja usar await por el cambio que le pueda realizar el botón al DOM
    //para ello usamos el aync en la función
    await btnIncrement.trigger('click');
    expect(counterLabel.text()).toContain(`Counter: ${value + 1}`)
    expect(squareLabel.text()).toContain(`Square: ${(value+1)*(value+1)}`)
  });

  test('decrement the counter when -1 button is clicked twice', async() => {

    const value = 5
    const wrapper = mount(MyCounter, {
      props:{
        value: value,
      }
    });

    const [counterLabel] = wrapper.findAll('h3');

    const btnDecrementTwice = wrapper.find('[data-testid="decrease-button"]');

    await btnDecrementTwice.trigger('click');
    await btnDecrementTwice.trigger('click');

    expect(counterLabel.text()).toContain(`Counter: ${value - 2}`)

  });

})









