impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        // 二分查找
        if nums.is_empty() {
            return 0;
        }

        let (mut left, mut right) = (0, nums.len() - 1);

        while left <= right {
            let mid = left + (right - left) / 2;
            match nums[mid].cmp(&target) {
                std::cmp::Ordering::Less => left = mid + 1,
                std::cmp::Ordering::Greater => right = mid - 1,
                std::cmp::Ordering::Equal => return mid as i32,
            }
        }
        return left as i32;
    }
}