// normal solution
function mySqrt(x: number): number {
  let result = 0

  while (true) {
    if (result * result > x) {
      return result - 1
    } else if (result * result === x) {
      return result
    }

    result++
  }
}

function mySqrt_binary_search(x: number): number {
  if (x === 0) {
    return 0
  }

  let left = 0
  let right = x

  while (true) {
    const mid = Math.ceil((left + right) / 2)

    if (mid < x / mid) {
      if (mid + 1 > x / (mid + 1)) {
        return mid
      }
      left = mid + 1
    } else if (mid > x / mid) {
      right = mid - 1
    } else {
      return mid
    }
  }
}
