// sum.test.js
import { expect, test } from 'vitest'
import {addArray, sum} from '../../src/helpers/sum'
import { describe } from 'node:test';

describe('sum function', ()=>{
  test('adds 1 + 2 to equal 3', () => {
    //Prepare
    const a = 1;
    const b = 4;
    //Incentive
    const result = sum(a, b);
    //Expected behavior
    expect(result).toBe(a + b);

    //expect(sum(1, 2)).toBe(3)
    //if(sum(1, 2) !== 3 ){
    //  throw new Error('That sum is not correct')
    //}
  })
});

describe('addArray function', ()=>{

  test('should return 0 if the array is empty', ()=> {
    const arr = []
    const result = addArray(arr)
    expect(result).toBe(0)
  })

  test('add numbers to an array and sum them', () => {
    //Prepare
    const arr = [1,2,3];
    //Incentive
    const result = addArray(arr);
    //Expected behavior
    expect(result).toBe(6)
  })


});


