def isAnagram1(s: str, t: str) -> bool:
    """
    The python list sort() method is O(n*log(n))
    """
    s1 = list(s)
    s1.sort()
    t1 = list(t)
    t1.sort()
    return s1 == t1


def isAnagram2(s: str, t: str) -> bool:
    """
    O(n) on average. Converting to a set() gains a lot of time on very long strings
    """
    for c in set(s + t):
        if s.count(c) != t.count(c):
            return False
    return True


s = "anagram"
t = "nagaram"

print(isAnagram1(s, t))
print(isAnagram2(s, t))

s = "rat"
t = "car"

print(isAnagram1(s, t))
print(isAnagram2(s, t))
