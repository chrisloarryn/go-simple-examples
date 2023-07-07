"use strict";
// Dado un string A muy grande y un string B muy chico
// Imprimir los substrings de A que son una permutacion de B
// def substrings_permutations(a, b):
// ejemplo:
// A: cbbabcdbbb
// B: abbc
// respuesta: cbba babc
const a = 'cbbabcdbbb', b = 'abbc';
function substringsPermutations(A, B) {
    const result = [];
    const lenA = A.length;
    const lenB = B.length;
    const counterB = {};
    for (let i = 0; i < lenB; i++) {
        counterB[B[i]] = (counterB[B[i]] || 0) + 1;
    }
    for (let i = 0; i <= lenA - lenB; i++) {
        const substring = A.substring(i, i + lenB);
        const counterSub = {};
        for (let j = 0; j < lenB; j++) {
            counterSub[substring[j]] = (counterSub[substring[j]] || 0) + 1;
        }
        if (isEqualObjects(counterSub, counterB)) {
            result.push(substring);
        }
    }
    return result;
}
function isEqualObjects(obj1, obj2) {
    if (Object.keys(obj1).length !== Object.keys(obj2).length) {
        return false;
    }
    for (const key in obj1) {
        if (obj1[key] !== obj2[key]) {
            return false;
        }
    }
    return true;
}
// Example usage:
const A = 'cbbabcdbbb';
const B = 'abbc';
const result = substringsPermutations(A, B);
console.log(result);
