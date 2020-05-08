package main

func jump(nums []int) int {
	cur, size := 0, len(nums)
	count := 0
	for cur < size-1 {
		step := nums[cur]
		max := 0
		temp := 0
		if cur+step >= size-1 {
			return count + 1
		}
		for i := cur + 1; i < cur+step; i++ {
			curval := nums[i] + i // 表示当前能够到达的长度，选取最大的
			if curval > max {
				max = curval
				temp = i
			}
		}
		cur = temp
		count++
	}
	return count
}

func main() {

}
