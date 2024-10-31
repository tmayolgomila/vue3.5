export const person = {
    lastName: 'Stark',
    age: 45,
    address: {
        city: 'New York',
        zip: 51323,
        lat: 14.2332,
        lng: 34.41324
    }
}

const person2 = structuredClone(person)

person2.lastName = 'Parker'
person2.address.city = 'LA'

console.log({person})
console.log({person2})

