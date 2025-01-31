myArray = [1,3,5]

for i, v in enumerate(myArray):
  print(myArray[i])

i = 0
while i < len(myArray):
  print(myArray[i])
  i += 1

def remove_end(l):
  if len(l):
    l[len(l)-1] = 0

remove_end(myArray)
print(myArray)

def remove_middle(l, idx):
  for i in range(idx, len(l)-1):
    l[i] = l[i+1]

myArray = [1, 2, 3]
remove_middle(myArray, 1)
print(myArray)