from typing import Optional, Any

class Stack:
  def __init__(self, arr: Optional[list[Any]]=None):
    if arr is None:
      self.arr = []
    else:
      self.arr = arr
  
  def append(self, val: Any):
    self.arr.append(val)

  def pop(self) -> Any:
    return self.arr.pop()
  
  def peek(self)-> Any:
    return self.arr[-1]
  
  def __str__(self) -> str:
    return f'[{",".join(str(x) for x in self.arr)}]'
  
s = Stack(list(range(10)))
s.append(101)
print(s)
s.pop()
print(s)
print(s.peek())
