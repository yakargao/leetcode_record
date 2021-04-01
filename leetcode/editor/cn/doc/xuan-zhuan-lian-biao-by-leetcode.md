
#### 方法 1：

**直觉**

链表中的点已经相连，一次旋转操作意味着：

* 先将链表闭合成环
* 找到相应的位置断开这个环，确定新的链表头和链表尾

![61.png](https://pic.leetcode-cn.com/e3371c6b03e3c8d3758dcf0b35a45d0a6b39c111373cf7b5bde53e14b6271a04-61.png)


> 新的链表头在哪里？

在位置 `n-k` 处，其中 `n` 是链表中点的个数，新的链表尾就在头的前面，位于位置 `n-k-1`。

> 我们这里假设 `k < n`
>
> 如果  `k >= n`  怎么处理？

`k` 可以被写成 `k = (k // n) * n + k % n` 两者加和的形式，其中前面的部分不影响最终的结果，因此只需要考虑 `k%n` 的部分，这个值一定比 `n` 小。

**算法**

算法实现很直接：

* 找到旧的尾部并将其与链表头相连 `old_tail.next = head`，整个链表闭合成环，同时计算出链表的长度 `n`。
* 找到新的尾部，第 `(n - k % n - 1)` 个节点 ，新的链表头是第 `(n - k % n)` 个节点。
* 断开环 `new_tail.next = None`，并返回新的链表头 `new_head`。

**实现**

 ![image.png](https://pic.leetcode-cn.com/a61c4ed3b26e5ee52dc03ff96f99c2c1162a37b38baf1318c0cbf288d97798e0-image.png) ![image.png](https://pic.leetcode-cn.com/7378ad1943fbb08eefdade59e1c154bcde7456f0beed5462b80e7a5872849cf1-image.png) ![image.png](https://pic.leetcode-cn.com/925962244f761bfae3f34ff0f4639e4dc5f3158b114e76e39af214eca059fb80-image.png) ![image.png](https://pic.leetcode-cn.com/c9b4c8a3ab7d7d14e4e6887d9efe74708d3f87ca8fd8a834fdc5d081351bff3e-image.png) ![image.png](https://pic.leetcode-cn.com/72fd989626c2a50e05456f115141923f55e9b6f98b9015ed02da51e861fdfddb-image.png) ![image.png](https://pic.leetcode-cn.com/328899e3498be9374b275e7600c80c03f0a10e68c252f41bdfe1ea3fe0dc2435-image.png) ![image.png](https://pic.leetcode-cn.com/498b6d70f52728da68996d05d14f30e9340739d552a1c6abfb24cdb27ddb153c-image.png) ![image.png](https://pic.leetcode-cn.com/ced79be85150072424f6d3a091ce0810ccc18bd09935bd9744dce5281d2b5e3f-image.png) ![image.png](https://pic.leetcode-cn.com/8f3fee4e0c90deaf28861d3fda347b7e8342b6729f31d38611a812dc4d4f7f07-image.png) ![image.png](https://pic.leetcode-cn.com/7a0bf66e90e30adcf7f4cc69855cd87a235d39254e81ef8f86f6f3928d1d402e-image.png) ![image.png](https://pic.leetcode-cn.com/f8dccee606e78854b993fe6ddf9eb9c52234924cbc11716ea2bb6f415b59e938-image.png) ![image.png](https://pic.leetcode-cn.com/7f9f494e18e1a5c7da4d718ff78d477192bb99665d3f016a216709d4d376a0b1-image.png) ![image.png](https://pic.leetcode-cn.com/3ae5cff7d6acc47909ff74289a27a16701d1b975610940d961139675455a71b6-image.png) ![image.png](https://pic.leetcode-cn.com/c6d1447ece5***2b52956102be2aee5962d3bdfa59240937127ed8cb674c3a68-image.png) 


```Java []
class Solution {
  public ListNode rotateRight(ListNode head, int k) {
    // base cases
    if (head == null) return null;
    if (head.next == null) return head;

    // close the linked list into the ring
    ListNode old_tail = head;
    int n;
    for(n = 1; old_tail.next != null; n++)
      old_tail = old_tail.next;
    old_tail.next = head;

    // find new tail : (n - k % n - 1)th node
    // and new head : (n - k % n)th node
    ListNode new_tail = head;
    for (int i = 0; i < n - k % n - 1; i++)
      new_tail = new_tail.next;
    ListNode new_head = new_tail.next;

    // break the ring
    new_tail.next = null;

    return new_head;
  }
}
```

```Python []
class Solution:
    def rotateRight(self, head: 'ListNode', k: 'int') -> 'ListNode':
        # base cases
        if not head:
            return None
        if not head.next:
            return head
        
        # close the linked list into the ring
        old_tail = head
        n = 1
        while old_tail.next:
            old_tail = old_tail.next
            n += 1
        old_tail.next = head
        
        # find new tail : (n - k % n - 1)th node
        # and new head : (n - k % n)th node
        new_tail = head
        for i in range(n - k % n - 1):
            new_tail = new_tail.next
        new_head = new_tail.next
        
        # break the ring
        new_tail.next = None
        
        return new_head
```

**复杂度分析**

* 时间复杂度：*O(N)*，其中 *N* 是链表中的元素个数
* 空间复杂度：*O(1)*，因为只需要常数的空间
