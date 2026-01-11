package main

// ：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	p := &ListNode{}
	q := head
	for q != nil {
		p.Next = q
		q = q.Next
	}

	return head
}

public:
    int cur=0;
    ListNode* removeNthFromEnd(ListNode* head, int n) {
       if(!head) return NULL;
       head->next = removeNthFromEnd(head->next,n);
       cur++;
       if(n==cur) return head->next;
       return head;
    }
};