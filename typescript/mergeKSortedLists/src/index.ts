import mergeKLists, { ListNode } from "./mergedSortedLists.js";

console.log(
  mergeKLists([
    new ListNode(1, new ListNode(4, new ListNode(5))),
    new ListNode(1, new ListNode(3, new ListNode(4))),
    new ListNode(2, new ListNode(6)),
  ])
);
