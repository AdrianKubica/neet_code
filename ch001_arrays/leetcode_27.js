const removeValue = (l, val) => {
  idx = 0;
  for (let i of l) {
    if (i !== val) {
      l[idx] = i;
      idx++;
    }
  }
  return idx;
};

l = [0, 1, 1, 2, 2, 3, 0, 4, 2];
x = removeValue(l, 2);
console.log(l, l.slice(0, x), x);
