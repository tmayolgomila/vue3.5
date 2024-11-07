import PokemonPicture from '@/components/PokemonPicture.vue';
import { mount } from '@vue/test-utils';

describe('<PokemonPicture />', () => {
  const imageSource =
    'https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/dream-world/26.svg';

  test('should render the hidden image when showPokemon prop is false', () => {
    const wrapper = mount(PokemonPicture, {
      props: {
        pokemonId: 26,
        showPokemon: false,
      },
    });
    console.log(wrapper.html());

    const image = wrapper.find('img');
    const attributes = image.attributes();

    expect(attributes).toEqual(
      expect.objectContaining({
        class: 'brightness-0 h-[200px]',
        src: imageSource,
      }),
    );
  });

  test('should render the image when showPokemon is true', () => {
    const wrapper = mount(PokemonPicture, {
      props: {
        pokemonId: 26,
        showPokemon: true,
      },
    });
    console.log(wrapper.html());

    const image = wrapper.find('img');
    const attributes = image.attributes();

    expect(attributes).toEqual(
      expect.objectContaining({
        class: 'fade-in h-[200px]',
        src: imageSource,
      }),
    );
  });
});
