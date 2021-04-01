### 解题思路
使用双端队列 队尾弹出放入对头

### 代码

```java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
        public ListNode rotateRight(ListNode head, int k) {
            if (head == null || head.next == null) return head;
            LinkedList<ListNode> deque = new LinkedList<>();
            //链表放入队列
            for (ListNode curr = head; curr != null; curr = curr.next) {
                deque.addLast(curr);
            }
            k = k % deque.size();//过滤掉整圈的
            ListNode h = head, pre, tail;//分别存头 倒数第一个 倒数第二个元素
            while (k-- > 0) {
                tail = deque.removeLast();//队列尾部的为要移动到最前面的元素
                pre = deque.peekLast();
                deque.addFirst(tail);//队尾元素放入对头
                //链表移动
                tail.next = h;
                pre.next = null;
                h = tail;
            }
            return h;
        }
    }
```