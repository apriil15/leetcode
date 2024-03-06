package main

func main() {

}

func reachNumber(target int) int {
	// 因為是對稱的，所以負數轉成正數，只看單邊
	if target < 0 {
		target *= -1
	}

	step := 0
	sum := 0

	for sum < target {
		step++
		sum += step
	}
	for (sum-target)%2 == 1 { // 如果差是奇數的話，就要繼續找
		step++
		sum += step
	}
	return step
}
