/**
 * @link https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
 */

// time: O(n)
// space: O(n)
function twoSum_map(numbers: number[], target: number): number[] {
  const map = new Map<number, number>()

  for (let i = 0; i < numbers.length; i++) {
    const theOtherNumber = target - numbers[i]

    if (!map.has(theOtherNumber)) {
      map.set(numbers[i], i + 1)
    } else {
      return [map.get(theOtherNumber)!, i + 1]
    }
  }
  throw new Error('')
}

// time: O(n)
// space: O(1)
function twoSum_twoPointers(numbers: number[], target: number): number[] {
  let left = 0
  let right = numbers.length - 1

  while (left < right) {
    if (numbers[left] + numbers[right] === target) {
      return [left + 1, right + 1]
    }
    if (numbers[left] + numbers[right] > target) {
      right--
    }
    if (numbers[left] + numbers[right] < target) {
      left++
    }
  }
  throw new Error('')
}

// time: O(nlogn)
// space: O(1)
function twoSum_binarySearch(numbers: number[], target: number): number[] {
  for (let index = 0; index < numbers.length; index++) {
    const wanted = target - numbers[index]

    let left = index + 1
    let right = numbers.length - 1

    while (left <= right) {
      const mid = Math.ceil((left + right) / 2)

      if (wanted === numbers[mid]) {
        return [index + 1, mid + 1]
      }
      if (wanted > numbers[mid]) {
        left = mid + 1
      }
      if (wanted < numbers[mid]) {
        right = mid - 1
      }
    }
  }

  throw new Error('')
}
