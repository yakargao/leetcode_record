// 注：为避免出现负数，使用左闭右开区间是最方便的
impl Solution {
    // lowerBound 返回最小的满足 nums[i] >= target 的 i
    // 如果数组为空，或者所有数都 < target，则返回 nums.length
    // 要求 nums 是非递减的，即 nums[i] <= nums[i + 1]

    // 左闭右开区间写法
    fn lower_bound(nums: &[i32], target: i32) -> usize {
        let mut left = 0;
        let mut right = nums.len(); // 左闭右开区间 [left, right)
        while left < right { // 区间不为空
            // 循环不变量：
            // nums[left-1] < target
            // nums[right] >= target
            let mid = left + (right - left) / 2;
            if nums[mid] < target {
                left = mid + 1; // 范围缩小到 [mid+1, right)
            } else {
                right = mid; // 范围缩小到 [left, mid)
            }
        }
        left // 或者 right
    }

    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        Self::lower_bound(&nums, target) as _
    }
}
