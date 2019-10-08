package main

// 统计一个数字在排序数组中出现的次数。
func main() {
	numberOfK := GetNumberOfK([]int{1, 2, 3, 4, 4, 4, 5}, 4)
	println(numberOfK)
}

/*
class Solution:
    def GetNumberOfK(self, data, k):
        # write code here
        return self.binSearch(data, k+0.5) - self.binSearch(data, k-0.5)

    def binSearch(self, data, find):
        s, end = 0, len(data)-1
        while(s <= end):
            mid = (end-s)//2+s
            if data[mid] < find:
                s = mid + 1
            elif data[mid] > find:
                end = mid - 1
        return s
*/

func GetNumberOfK(data []int, k int) int {
	return findKIdx(data, float64(k)+0.5) - findKIdx(data, float64(k)-0.5)
}

// 利用二分搜索来找到索引
func findKIdx(data []int, find float64) int {
	start, end := 0, len(data)-1


	for start <= end {
		mid := (end-start)/2 + start
		if float64(data[mid]) < find {
			start = mid + 1
		} else if float64(data[mid]) > find {
			end = mid - 1
		}
	}
	return start
}
