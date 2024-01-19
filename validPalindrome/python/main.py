import re


def isPalindrome(s: str) -> bool:
    s = re.sub("[^0-9a-zA-Z]", "", s)
    s = s.lower()
    l = len(s)
    for i in range(l):
        l -= 1
        if s[i] != s[l]:
            return False

    return True


s1 = "A man, a plan, a canal: Panama"
s2 = "race a car"
s3 = " "


print(True, isPalindrome(s1))
print(False, isPalindrome(s2))
print(True, isPalindrome(s3))
