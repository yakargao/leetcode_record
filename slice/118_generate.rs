impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut v = vec![];
        for i in 0..num_rows as usize{
            v.push(vec![1;i+1]);
            for j in 1..i{
                v[i][j] = v[i-1][j] + v[i-1][j-1];
            }
        }
        v
    }
}