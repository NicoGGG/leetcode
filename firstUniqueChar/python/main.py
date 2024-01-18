def firstUniqChar(s: str) -> int:
    m: dict[str, int] = {}
    for i in range(len(s)):
        c = s[i]
        if s[i] in m.keys():
            m[c] = -1
        else:
            m[c] = i
    r: int = -1
    for v in m.values():
        if (v < r and v != -1) or r == -1:
            r = v
    return r


s1 = "leetcode"
s2 = "loveleetcode"
s3 = "aabb"

print(firstUniqChar(s1))
print(firstUniqChar(s2))
print(firstUniqChar(s3))
