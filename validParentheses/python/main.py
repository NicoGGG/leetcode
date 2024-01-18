def isValid(s: str) -> bool:
    stack = []
    if len(s) % 2 != 0:
        return False
    for c in s:
        if c in "([{":
            stack.append(c)
        if c == "]":
            if not stack or stack.pop() != "[":
                return False
        if c == "}":
            if not stack or stack.pop() != "{":
                return False
        if c == ")":
            if not stack or stack.pop() != "(":
                return False
    return not stack


s1 = "()"
s2 = "()[]{}"
s3 = "(]"
s4 = "{[]}"
s5 = "["
s6 = "]"
s7 = "){"
s8 = "(("

print(True, isValid(s1))
print(True, isValid(s2))
print(False, isValid(s3))
print(True, isValid(s4))
print(False, isValid(s5))
print(False, isValid(s6))
print(False, isValid(s7))
print(False, isValid(s8))
