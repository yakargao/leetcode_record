思路:不同的两数相互抵消，最后剩下的肯定是多于一半的那个数。
```
class Solution {
public:
    int moreThanHalfNum_Solution(vector<int>& nums) {
        int val, cnt = 0;
        
        for (auto x : nums)
        {
            if (!cnt) val = x, cnt ++ ;     //目标值与其他值刚好配对抵消时，重置计数
            else
            {
                if (x == val) cnt ++ ;
                else cnt -- ;
            }
        }
        
        return val;                         //最后剩下的一定是多于半数的目标值
    }
};
```