package hot100

import (
	"fmt"
	"math"
)

/*
53. 最大子序和: https://leetcode-cn.com/problems/maximum-subarray
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：
输入：nums = [1]
输出：1

示例 3：
输入：nums = [0]
输出：0

示例 4：
输入：nums = [-1]
输出：-1

示例 5：
输入：nums = [-100000]
输出：-100000


提示：
1 <= nums.length <= 3 * 104
-105 <= nums[i] <= 105

进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/

// 考虑最大值为负数的时候，maxNum初始化的值不能是0
// 需要考虑单个值最大的时候
func MaxSubArray(nums []int) int {
	maxNum := math.MaxInt64 * -1
	var tmp int
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums); i++ {
		tmp = nums[i]
		if tmp > maxNum {
			maxNum = tmp
		}
		for k := i + 1; k < len(nums); k++ {
			tmp += nums[k]
			if tmp > maxNum {
				maxNum = tmp
			}
		}

		if tmp > maxNum {
			maxNum = tmp
		}

		fmt.Print(maxNum, "\r\n")
	}

	return maxNum
}

func MaxSubArr(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
