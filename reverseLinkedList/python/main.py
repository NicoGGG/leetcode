from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def reverseList(head: Optional[ListNode]) -> Optional[ListNode]:
    prev = None
    while head:
        next = head.next
        head.next = prev
        prev = head
        head = next
    return prev


def buildLinkedList(a: list) -> Optional[ListNode]:
    if len(a) == 0:
        return None
    head = ListNode(a[0])
    node = head
    for i in range(1, len(a)):
        node.next = ListNode(a[i])
        node = node.next
    return head


def printLinkedList(head: Optional[ListNode]):
    if head is None:
        print("Empty list")
        return
    node = head
    while node:
        print(node.val)
        node = node.next


l1 = [1, 2, 3, 4, 5]
l2 = [1, 2]
l3 = []
ll1 = buildLinkedList(l1)
ll2 = buildLinkedList(l2)
ll3 = buildLinkedList(l3)

print("ReverseLinkedList 1")
printLinkedList(reverseList(ll1))
print("ReverseLinkedList 2")
printLinkedList(reverseList(ll2))
print("ReverseLinkedList 3")
printLinkedList(reverseList(ll3))
