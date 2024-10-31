
interface Hero {
    name: string;
    age: number;
    codeName: string;
    power?: string
}

export const person:Hero = {
    name: 'John Doe',
    age: 30,
    codeName: 'JD',
    power: 'money'
}

//const {age, name, power = 'No tiene pwr'} = person
//console.log({age, name, power})

interface CreateHeroArgs {
    name: string;
    age: number;
    codeName: string;
    power?: string
}

const createHero = ({name, age, codeName, power }: CreateHeroArgs) => ({
    id:1234,
    name: name,
    age: age,
    codeName: codeName,
    power: power ?? 'no tiene'
})

console.log(createHero(person))













