// Given integer n, return true if n is a power of 3. Otherwise, return false.
//

const n1 = 27
const n2 = 0
const n3 = -1

/**
  * @param {number} n
  * @returns {boolean}
  */
function isPowerOfThree1(n) {
  if (n == 1) {
    return true
  } else if (n < 1) {
    return false
  }
  return isPowerOfThree1(n / 3)
}

console.log(isPowerOfThree1(n1))
console.log(isPowerOfThree1(n2))
console.log(isPowerOfThree1(n3))

/**
  * @param {number} n
  * @returns {boolean}
  */
function isPowerOfThree2(n) {
  return n > 0 && 3**19 % n === 0
}

console.log(isPowerOfThree2(n1))
console.log(isPowerOfThree2(n2))
console.log(isPowerOfThree2(n3))

