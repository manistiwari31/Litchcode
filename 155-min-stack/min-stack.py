class MinStack:

    def __init__(self):
        self.item_stack = []
        self.min_stack = []
        

    def push(self, x: int) -> None:
        self.item_stack.append(x)
        if not self.min_stack or x <= self.min_stack[-1]:
            self.min_stack.append(x)
        

    def pop(self) -> None:
        if self.item_stack and self.item_stack.pop() == self.min_stack[-1]:
            self.min_stack.pop()
        

    def top(self) -> int:
        return self.item_stack[-1]
        

    def getMin(self) -> int:
        return self.min_stack[-1]
        