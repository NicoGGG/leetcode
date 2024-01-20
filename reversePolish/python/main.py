from typing import List


def evalRPN(tokens: List[str]) -> int:
    """
    Example 1
    tokens = ["2","1","+","3","*"]
    Pop into a stack until 2 nums consecutive. Pop them with the following operand
    2 + 1 = 3
    Push the result to the stack
    => ["*", "3", "3"]
    Two consecutive number, redo the same operation. End when stack is one number

    Example 2
    tokens = ["4","13","5","/","+"]
    Pop until two number => Stack: [+, /, 5, 13]
    Pop two numbers and a operand
    (13 / 5) = 2 => Stack :[+, 2]
    Pop last number => Stack :[+, 2, 4]
    => [6]
    """
    stack: List[str] = []
    while len(stack) != 1 or tokens:
        if (
            stack
            and len(stack) >= 2
            and stack[-1].lstrip("-").isnumeric()
            and stack[-2].lstrip("-").isnumeric()
        ):
            c = 2
        elif stack and len(stack) >= 1 and stack[-1].lstrip("-").isnumeric():
            c = 1
        else:
            c = 0
        while c != 2 and tokens:
            n = tokens.pop()
            if n.lstrip("-").isnumeric():
                c += 1
            else:
                c = 0
            stack.append(n)
        if len(stack) == 1:
            return int(stack.pop())
        n1 = int(stack.pop())
        n2 = int(stack.pop())
        o = stack.pop()
        if o == "+":
            stack.append(str(n1 + n2))
        elif o == "-":
            stack.append(str(n1 - n2))
        elif o == "*":
            stack.append(str(n1 * n2))
        elif o == "/":
            stack.append(str(int(n1 / n2)))
    return int(stack.pop())


t1 = ["2", "1", "+", "3", "*"]
t2 = ["4", "13", "5", "/", "+"]
t3 = ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
t4 = ["18"]

print(9, evalRPN(t1))
print(6, evalRPN(t2))
print(22, evalRPN(t3))
print(18, evalRPN(t4))
