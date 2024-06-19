'use strict'

import {should, assert} from 'chai';

let numbers = [1, 2, 3, 4, 5];

console.log('numbers', numbers)
assert.isArray(numbers, 'is array of numbers');
assert.include(numbers, 2, 'array contains 2');
assert.lengthOf(numbers, 5, 'array contains 5 numbers');