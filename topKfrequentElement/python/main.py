from typing import List, Dict, Optional


def topKFrequent(nums: List[int], k: int) -> List[int]:
    f: Dict[int, int] = {}
    for n in nums:
        if n in f.keys():
            f[n] += 1
        else:
            f[n] = 1
    r1 = list(
        map(lambda x: x[0], sorted(f.items(), key=lambda x: x[1], reverse=True)[0:k])
    )
    return r1


l1 = [1, 1, 1, 2, 2, 2, 2, 3]
k1 = 2
l2 = [1]
k2 = 1
l3 = [4, 1, -1, 2, -1, 2, 3]
k3 = 2
print(topKFrequent(l1, k1))
print(topKFrequent(l2, k2))
print(topKFrequent(l3, k3))
