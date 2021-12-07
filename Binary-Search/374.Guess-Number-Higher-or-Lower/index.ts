/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	            -1 if num is lower than the guess number
 *			             1 if num is higher than the guess number
 *                       otherwise return 0
 * var guess = function(num) {}
 */

const target = 6

function guessNumber(n: number): number {
  let left = 1
  let right = n

  while (left <= right) {
    const mid = Math.floor((left + right) / 2)

    if (guess(mid) === -1) {
      right = mid - 1
    } else if (guess(mid) === 1) {
      left = mid + 1
    } else {
      return mid
    }
  }
  throw new Error('')
}

function guess(num: number): number {
  if (target < num) {
    return -1
  } else if (target > num) {
    return 1
  } else {
    return 0
  }
}

console.log(guessNumber(10))
