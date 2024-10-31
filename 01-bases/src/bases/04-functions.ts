
//function greetPerson(name: string){
//
//    return `Hola, ${ name }`
//}

//const greetPerson = (name:string) => {
//    return `Hola, ${ name }`
//}

//Si la funcion solo tiene un return:

const greetPerson = (name:string) => `Hola, ${ name }`;

const getUser = (uid:string) => ({
        uid,
        username: 'Tony001'
    })


const heroes = [
    {
        id: 1,
        name: 'heroe1',
    },
    {
        id: 2,
        name: 'heroe2',
        power: 'Powr1'
    },

]

const hero = heroes.find((h) => h.id === 2)
console.log(hero?.power?.toUpperCase())