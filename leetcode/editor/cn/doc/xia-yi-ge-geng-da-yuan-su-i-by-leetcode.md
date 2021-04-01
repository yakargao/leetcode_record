#### 方法一：暴力解法

题目中的关键信息：两个数组各自 **没有重复元素**。模拟题目的意思：对于每一个 `nums1[i]` 中的元素，先在 `nums2` 中找到它，然后向右遍历找到第 1 个大于 `nums1[i]` 的元素。

**参考代码 1**：

```Java []
import java.util.Arrays;

public class Solution {

    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int len1 = nums1.length;
        int len2 = nums2.length;

        int[] res = new int[len1];
        if (len1 < 1) {
            return res;
        }

        for (int i = 0; i < len1; i++) {
            int curVal = nums1[i];
            int j = 0;
            while (j < len2 && nums2[j] != curVal) {
                j++;
            }

            // 此时 nums[j] = nums[i]
            j++;
            while (j < len2 && nums2[j] < curVal) {
                j++;
            }

            if (j == len2) {
                res[i] = -1;
                continue;
            }
            res[i] = nums2[j];
        }
        return res;
    }
}
```

**复杂度分析**：

+ 时间复杂度：*O(NM)*，这里 *N* 是数组 `nums1` 的长度， *M* 是数组 `nums2` 的长度，对于 *N* 个 `nums1` 中的元素，最差情况下需要遍历完 `nums2` 中的每个元素，因此时间复杂度为 *O(NM)*；
+ 空间复杂度：
  + 如果不计算保存结果的空间，空间复杂度为 *O(1)*；
  + 如果计算保存结果的空间，空间复杂度为 *O(N)*。

「暴力解法」时间复杂度高 ，空间复杂度低，优化「暴力解法」的思路是「空间换时间」。事实上，找右边第 1 个大于比自己大的数，这是一个典型的「栈」的应用（也叫「单调栈」），将每一个元素进栈一次、出栈一次，这样的过程中就可以找到找右边第 1 个大于比自己大的数。


---

#### 方法二：栈（单调栈）

根据题意，数组 `nums1` 视为询问。我们可以：

+ 先对 `nums2` 中的每一个元素，求出它的右边第一个更大的元素；
+ 将上一步的对应关系放入哈希表（HashMap）中；
+ 再遍历数组 `nums1`，根据哈希表找出答案。

下面我们解释如何得到 `nums2` 的每个元素右边第 1 个比它大的元素，这里以 `nums2 = [2, 3, 5, 1, 0, 7, 3]` 为例进行说明，我们从左到右遍历数组 `nums2` 中的元素。

友情提示：只需要看文字和图例的前几句话，找到解决问题的一帮规律，这样可以节约阅读时间。

 ![0496.001.jpeg](https://pic.leetcode-cn.com/1616403527-TfSHQh-0496.001.jpeg) ![0496.002.jpeg](https://pic.leetcode-cn.com/1616403527-wkRXsj-0496.002.jpeg) ![0496.003.jpeg](https://pic.leetcode-cn.com/1616403527-lxOUsI-0496.003.jpeg) ![0496.004.jpeg](https://pic.leetcode-cn.com/1616403527-vQnAid-0496.004.jpeg) ![0496.005.jpeg](https://pic.leetcode-cn.com/1616403527-lSfwYs-0496.005.jpeg) ![0496.006.jpeg](https://pic.leetcode-cn.com/1616403527-EwkSGP-0496.006.jpeg) ![0496.007.jpeg](https://pic.leetcode-cn.com/1616403527-AORSDA-0496.007.jpeg) ![0496.008.jpeg](https://pic.leetcode-cn.com/1616403527-lQdwMk-0496.008.jpeg) ![0496.009.jpeg](https://pic.leetcode-cn.com/1616403527-vUWIDY-0496.009.jpeg) 


**归纳重点**：

+ 可以发现，我们维护的栈恰好保证了单调性：**栈中的元素从栈顶到栈底是单调不降的**。当我们遇到一个新的元素 `nums2[i]` 时，我们判断栈顶元素是否小于 `nums2[i]`，如果是，那么栈顶元素的下一个更大元素即为 `nums2[i]`，我们将栈顶元素出栈。重复这一操作，直到栈为空或者栈顶元素大于 `nums2[i]`。此时我们将 `nums2[i]` 入栈，保持栈的单调性，并对接下来的 `nums2[i + 1], nums2[i + 2] ...` 执行同样的操作。

**参考代码 2**：

```Java []
import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;
import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

public class Solution {

    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int len1 = nums1.length;
        int len2 = nums2.length;

        Deque<Integer> stack = new ArrayDeque<>();
        Map<Integer, Integer> map = new HashMap<>();
        // 先处理 nums2，把对应关系存入哈希表
        for (int i = 0; i < len2; i++) {
            while (!stack.isEmpty() && stack.peekLast() < nums2[i]) {
                map.put(stack.removeLast(), nums2[i]);
            }
            stack.addLast(nums2[i]);
        }

        // 遍历 nums1 得到结果集
        int[] res = new int[len1];
        for (int i = 0; i < len1; i++) {
            res[i] = map.getOrDefault(nums1[i], -1);
        }
        return res;
    }
}
```
**复杂度分析**：
+ 时间复杂度：*O(N + M)*，分别遍历数组 `nums1` 和数组 `nums2` 各一次即可，对于 `nums2` 中的每个元素，进栈一次，出栈一次；
+ 空间复杂度：*O(N)*。我们在遍历 `nums2` 时，需要使用栈，以及哈希映射用来临时存储答案。