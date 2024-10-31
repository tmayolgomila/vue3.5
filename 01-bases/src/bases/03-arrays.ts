

const numberArray = [1,2,3,4,5] //as const ==> si le ponemos as const lo hacemos de solo lectura y no se podría modificar

numberArray.push(6)

const numberArray2 : (number|string)[] = [...numberArray]

numberArray2.push('7')

console.log({numberArray, numberArray2})

export {}












