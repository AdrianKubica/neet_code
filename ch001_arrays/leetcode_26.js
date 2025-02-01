const removeDuplicates = (l) => {
  let idx = 0;
  for (let i = 1; i < l.length; i++) {
    if (l[i] !== l[idx]) {
      idx++;
      l[idx] = l[i];
    }
  }
  return idx + 1;
};
l = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4];
x = removeDuplicates(l);
console.log(l, l.slice(0, x), x);
