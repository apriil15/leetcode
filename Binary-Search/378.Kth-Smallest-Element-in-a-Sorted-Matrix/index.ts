// time: O(n^2*log(n))
// space: O(n)
function kthSmallest(matrix: number[][], k: number): number {
  let result: number[] = []

  for (let i = 0; i < matrix.length; i++) {
    if (i === 0) {
      result.push(...matrix[0])
      continue
    }

    if (matrix[i][0] >= result[result.length - 1]) {
      result.push(...matrix[i])
      continue
    }
    if (matrix[i][matrix[i].length - 1] <= result[0]) {
      result = matrix[i].concat(result)
      continue
    }

    for (const num of matrix[i]) {
      if (num >= result[result.length - 1]) {
        result.push(num)
        continue
      }
      if (num <= result[0]) {
        result = [num].concat(result)
        continue
      }

      const index = getBiggerOrEqualIndex(result, num)
      result.splice(index, 0, num)
    }
  }
  return result[k - 1]
}

// binary search
function getBiggerOrEqualIndex(numbers: number[], target: number): number {
  let left = 0
  let right = numbers.length - 1

  while (left <= right) {
    const mid = Math.floor((left + right) / 2)

    if (numbers[mid] === target) {
      return mid
    }
    if (numbers[mid] < target) {
      if (numbers[mid + 1] >= target) {
        return mid + 1
      } else {
        left = mid + 2
      }
    }
    if (numbers[mid] > target) {
      right = mid - 1
    }
  }

  return left
}

console.log(
  kthSmallest(
    [
      [1, 5, 9],
      [10, 11, 13],
      [12, 13, 15]
    ],
    8
  )
)
