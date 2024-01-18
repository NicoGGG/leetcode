class MinStack:
    def __init__(self):
        self.stack = []
        self.min = []

    def push(self, val: int) -> None:
        self.stack.append(val)
        if not self.min or val <= self.min[-1]:
            self.min.append(val)

    def pop(self) -> None:
        if self.stack[-1] == self.min[-1]:
            del self.min[-1]
        del self.stack[-1]

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min[-1]


minStack = MinStack()
minStack.push(2)
minStack.push(0)
minStack.push(-3)
print(-3, minStack.getMin())
minStack.pop()
print(0, minStack.top())
print(-2, minStack.getMin())
