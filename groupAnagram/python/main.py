from typing import List, Dict


def groupAnagrams(strs: List[str]) -> List[List[str]]:
    r: List[List[str]] = []
    m: Dict[str, List[str]] = {}
    for s in strs:
        k = "".join(sorted(list(s)))
        if k in m.keys():
            m[k].append(s)
        else:
            m[k] = [s]
    for v in m.values():
        r.append(v)

    return r


def isAnagram(s: str, t: str) -> bool:
    for c in set(s + t):
        if s.count(c) != t.count(c):
            return False
    return True


strs1 = ["eat", "tea", "tan", "ate", "nat", "bat"]
strs2 = [""]
strs3 = ["a"]

print(groupAnagrams(strs1))
print(groupAnagrams(strs2))
print(groupAnagrams(strs3))
