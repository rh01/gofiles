/*
 * Copyright 2019 @rh01 https://github.com/rh01
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
two sum
*/

package leetcode001


/** use hash map */
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for k, v := range nums {
		// 判断，如果不存在则将map(k, v)放入到hashmap中
		if _, exist := hashmap[v]; !exist {
			hashmap[v] = k
		}

		// 然后检查
		want := target - v
		if wantK, exist := hashmap[want]; exist && wantK != k {
			return []int{k, wantK}
		}
	}
	return []int{}
}

/** brute solution*/
func twoSumBrute(nums []int, target int) []int {
	for k1, v1 := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if v1+nums[k2] == target {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}