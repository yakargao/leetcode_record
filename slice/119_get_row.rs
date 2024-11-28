impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let mut dp = vec![1; (row_index + 1) as usize];
        for i in 1..=row_index as usize {
            for j in (1..i).rev() {
                dp[j] += dp[j - 1];
            }
        }
        dp
    }
}
