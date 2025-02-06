class MinStack:

    def __init__(self):
        self.stack = []

    def push(self, val: int) -> None:
        self.stack.append((val, val if not self.stack else min(val, self.stack[-1][1])))
        
    def pop(self) -> None:
        self.stack.pop()[0]

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]
    
minStack = MinStack()
minStack.push(-10)
minStack.push(14)
minStack.push(-20)
minStack.pop()
minStack.push(10)
minStack.push(-7)
minStack.getMin()
print(minStack.stack)
