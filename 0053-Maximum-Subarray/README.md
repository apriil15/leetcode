# 53. Maximum Subarray

## 題目

[Maximum Subarray - LeetCode](https://leetcode.com/problems/maximum-subarray/)

## 思路

其實是一個領土擴張的概念，我吃掉你有變大的話，就吃掉

沒變大的話，我就變成你就好

(-2) + 1 跟 1 做比較，後者大，所以 current 變成 1

1 + (-3) 跟 -3 做比較，前者大，所以 current 變成 1 + (-3)

這樣一直迭代，同時拿 current 跟 max 做比較，較大的成為新的 max
