class Solution:
  def removeDuplicates(self, nums: list[int]) -> int:
    s = sorted(set(nums))
    nums[:len(s)] = s
    return len(nums)

  def removeDuplicatesInPlace(self, nums: list[int]) -> int:
    idx = 0
    for i in range(1, len(nums)):
      if nums[i] != nums[idx]:
        idx += 1
        nums[idx] = nums[i]
    return idx + 1

solution = Solution()
l = [0,0,1,1,1,2,2,3,3,4]
x = solution.removeDuplicatesInPlace(l)
print(l, l[:5], x)

