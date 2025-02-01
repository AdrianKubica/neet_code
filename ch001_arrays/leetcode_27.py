def removeValue(l: list[int], val: int) -> int:
  idx = 0
  for v in l:
    if v != val:
      l[idx] = v
      idx += 1
  return idx

l = [0,1,1,2,2,3,0,4,2]
x = removeValue(l, 2)
print(l, l[:x], x)
