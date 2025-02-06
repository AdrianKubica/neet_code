class Stack {
  constructor(arr = []) {
    this.arr = arr;
  }

  push(value) {
    this.arr.push(value);
  }

  pop() {
    return this.arr.pop();
  }

  peek() {
    return this.arr[this.arr.length - 1];
  }
}

l = [1, 2, 3, 4, 5];
stack = new Stack(l);

stack.push(101);
console.log(stack, l);
x = stack.pop();
console.log(stack, x);
console.log(stack.peek());
