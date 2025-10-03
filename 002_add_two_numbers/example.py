# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        def traverseListReversed(l):
            i = 0
            s = 0
            while l.next is not None:
                s += pow(10, i) * l.val
                i += 1
                l = l.next
            s += pow(10, i) * l.val
            return s
        
        def generateLinkedList(s):
            first = ListNode(val = s % 10)
            previous = first
            s //= 10
            while s != 0:
                previous.next = ListNode(val = s % 10)
                previous = previous.next
                s //= 10
            
            return first

        return generateLinkedList(traverseListReversed(l1) + traverseListReversed(l2))
