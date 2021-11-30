// dp solution
// time: O(m*n)
// space: O(m*n)
function findLength(nums1: number[], nums2: number[]): number {
  // number[][]
  const dp = Array.from(Array<number>(nums1.length + 1), () =>
    Array<number>(nums2.length + 1)
  )

  let max = 0

  for (let i = 0; i <= nums1.length; i++) {
    for (let j = 0; j <= nums2.length; j++) {
      if (i === 0 || j === 0) {
        dp[i][j] = 0
      } else if (nums1[i - 1] === nums2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1] + 1
        max = Math.max(max, dp[i][j])
      } else {
        dp[i][j] = 0
      }
    }
  }

  return max
}

console.log(findLength([1, 2, 3, 2, 1], [3, 2, 1, 4, 7]))
