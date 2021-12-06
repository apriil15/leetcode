// space: O(1)
function kthSmallest(matrix: number[][], k: number): number {
  let left = matrix[0][0]
  let right = matrix[matrix.length - 1][matrix.length - 1]

  while (left < right) {
    const mid = Math.floor((left + right) / 2)

    const count = getSmallerOrEqualCount(matrix, mid)

    if (count < k) {
      left = mid + 1
    } else {
      right = mid
    }
  }
  return left
}

function getSmallerOrEqualCount(matrix: number[][], mid: number) {
  let count = 0
  let j = matrix.length - 1

  for (let i = 0; i < matrix.length; i++) {
    while (j >= 0 && matrix[i][j] > mid) {
      j--
    }
    count += j + 1
  }
  return count
}
