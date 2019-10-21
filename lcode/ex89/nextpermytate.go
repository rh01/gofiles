// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

// 必须原地修改，只允许使用额外常数空间。

// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/next-permutation
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package ex89

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}
	length := len(nums)
	// 标记是否是完全逆序
	i, j := 0, length-1

	for i = length - 1; i>0 && nums[i-1] >= nums[i]; i-- {
	}

	if i !=0 {
		for j = length - 1; nums[i-1] >= nums[j] && i < j; j-- {
		}
		swap(&nums, j, i-1)

		for k, j := i, length-1; k < j; k, j = k+1, j-1 {
			swap(&nums, k, j)
		}
		
	} else {
		for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
			swap(&nums, i, j)
		}
	}
	//nextPermutation(nums[idx:])
	return
}

func swap(nums *[]int, a, b int) {
	(*nums)[a], (*nums)[b] = (*nums)[b], (*nums)[a]
}
