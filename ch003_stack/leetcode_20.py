def isValid(s: str) -> bool:
        stack = []
        pairs = {')': '(', ']': '[', '}': '{'}
        for value in s:
            if value in pairs.values():
                stack.append(value)
                continue
            if len(stack) == 0:
                return False
            if pairs.get(value, None) != stack.pop():
                return False
        return len(stack) == 0

print(isValid("[]{}()[{}]"))