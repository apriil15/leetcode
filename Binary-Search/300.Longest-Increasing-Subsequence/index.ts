// time: O(n log(n))
function lengthOfLIS(nums: number[]): number {
  const arr: number[] = []

  for (const num of nums) {
    if (arr.length === 0 || arr[arr.length - 1] < num) {
      arr.push(num)
      continue
    }

    const position = getReplacePosition(arr, num)
    arr[position] = num
  }

  return arr.length
}

// using binary search to find the replace position
function getReplacePosition(arr: number[], num: number): number {
  let left = 0
  let right = arr.length - 1

  while (left <= right) {
    const mid = Math.ceil((left + right) / 2)

    if (num < arr[mid]) {
      right = mid - 1
    } else if (num > arr[mid]) {
      left = mid + 1
    } else {
      return mid
    }
  }

  return left
}

console.log(lengthOfLIS([10, 9, 2, 5, 3, 4]))
