def double(l: list[int]) -> int:
  return l.extend(l) or l

def double2(l: list[int]) -> int:
  return l * 2

l1 = [1,2,3,4,5]
print(double(l1))
print(double2(l1))
