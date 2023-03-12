import mergeKLists from "../src/mergedSortedLists";
import { ListNode } from "../src/mergedSortedLists";

describe("Test", () => {
  it("Example 1", () => {
    const lists = [
      new ListNode(1, new ListNode(4, new ListNode(5))),
      new ListNode(1, new ListNode(3, new ListNode(4))),
      new ListNode(2, new ListNode(6)),
    ];
    const mergedList = mergeKLists(lists);
    expect(mergedList?.val).toBe(1);
    // Check that next is not null
    expect(mergedList?.next).not.toBeNull();
    if (mergedList?.next === null) return;
    expect(mergedList?.next.next).not.toBeNull();
    if (mergedList?.next.next === null) return;
    expect(mergedList?.next.next.next).not.toBeNull();
    if (mergedList?.next.next.next === null) return;
    expect(mergedList?.next.next.next.next).not.toBeNull();
    if (mergedList?.next.next.next.next === null) return;
    expect(mergedList?.next.next.next.next.next).not.toBeNull();
    if (mergedList?.next.next.next.next.next === null) return;
    expect(mergedList?.next.next.next.next.next.next).not.toBeNull();
    if (mergedList?.next.next.next.next.next.next === null) return;
    expect(mergedList?.next.next.next.next.next.next.next).not.toBeNull();
    if (mergedList?.next.next.next.next.next.next.next === null) return;
    expect(mergedList?.next.next.next.next.next.next.next.next).toBeNull();

    expect(mergedList?.val).toBe(1);
    expect(mergedList?.next.val).toBe(1);
    expect(mergedList?.next.next.val).toBe(2);
    expect(mergedList?.next.next.next.val).toBe(3);
    expect(mergedList?.next.next.next.next.val).toBe(4);
    expect(mergedList?.next.next.next.next.next.val).toBe(4);
    expect(mergedList?.next.next.next.next.next.next.val).toBe(5);
    expect(mergedList?.next.next.next.next.next.next.next.val).toBe(6);
  });

  it("Example 2", () => {
    const lists: ListNode[] = [];
    const mergedList = mergeKLists(lists);
    expect(mergedList).toEqual(null);
  });

  it("Example 3", () => {
    const lists: ListNode[] = [new ListNode()];
    const mergedList = mergeKLists(lists);
    expect(mergedList).toEqual(null);
  });
});
