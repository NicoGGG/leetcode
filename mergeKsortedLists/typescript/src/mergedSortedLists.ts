// Definition for singly-linked list.
export class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

// Return the index of the min in the array
const min = (list: Array<number>): number => {
  let min: number = 1e4;
  let ret: number = 0;
  for (const i in list) {
    if (list[i] < min) {
      min = list[i];
      ret = +i;
    }
  }
  return ret;
};

const mergeKLists = (lists: Array<ListNode | null>): ListNode | null => {
  lists = lists.filter((i) => i !== null);
  const retRoot: ListNode = new ListNode();
  let ret: ListNode | null = retRoot;
  let currentArray: number[] = [];
  let currentMinNode: ListNode | null;
  while (lists.length) {
    currentArray = [];
    for (const i in lists) {
      currentArray.push(lists[i]?.val!);
    }
    currentMinNode = lists.reduce((prev, next) =>
      prev?.val! < next?.val! ? prev : next
    );
    let currentMinNodeIndex = lists.indexOf(currentMinNode);
    let nextVal = currentMinNode?.val!;
    ret.next = new ListNode(nextVal);
    ret = ret.next;
    lists[currentMinNodeIndex] = currentMinNode?.next!;
    lists = lists.filter((i) => i !== null);
  }
  return retRoot.next;
};

export default mergeKLists;
