def isPalindrome1(x: int) -> bool:
    """
    Generic way
    """
    if x < 0:
        return False
    if x == 0:
        return True
    originalX = x
    reversedX = 0
    while x != 0:
        reversedX = reversedX * 10 + x % 10
        x //= 10
    return originalX == reversedX


def isPalindrome2(x: int) -> bool:
    """
    Pythonic way
    """
    return str(x) == str(x)[::-1]


x1 = 121
x2 = -121
x3 = 10

print(True, isPalindrome1(x1))
print(False, isPalindrome1(x2))
print(False, isPalindrome1(x3))


print(True, isPalindrome2(x1))
print(False, isPalindrome2(x2))
print(False, isPalindrome2(x3))
