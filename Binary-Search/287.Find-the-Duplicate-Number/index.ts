// time: O(n)
// space: O(n)
function findDuplicate_map(nums: number[]): number {
  const map = new Map<number, number>()

  for (const num of nums) {
    if (map.has(num)) {
      return num
    } else {
      map.set(num, 1)
    }
  }

  throw new Error('')
}

// time: O(n)
// space: O(1)
function findDuplicate_tortoise_hare(nums: number[]): number {
  let slow = 0
  let fast = 0

  // find the point that slow and fast meet in the circle
  do {
    slow = nums[slow]
    fast = nums[nums[fast]]
  } while (slow != fast)

  // reset slow to the beginning
  slow = 0

  // find the intersection
  do {
    slow = nums[slow]
    fast = nums[fast]
  } while (slow != fast)

  return slow
}
