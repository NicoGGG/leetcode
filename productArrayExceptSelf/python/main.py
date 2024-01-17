from typing import List


def productExceptSelf1(nums: List[int]) -> List[int]:
    r = []
    for i, n in enumerate(nums):
        left = product_array(nums[:i])
        right = product_array(nums[i + 1 :])
        r.append(left * right)

    return r


def product_array(nums):
    r = 1
    for n in nums:
        r *= n
    return r


def productExceptSelf2(nums: List[int]) -> List[int]:
    """
    Time complexity: O(n)
    First solution is O(n^2) because of the product_array function
    We can actually do this in O(n) time and O(1) space by using prefix and suffix variables
    We loop over the array twice, once to grow the prefix and once to grow the suffix, then we multiply them together during the suffix loop
    """
    r = []
    prefix = 1
    suffix = 1
    for i, n in enumerate(nums):
        r.append(prefix)
        prefix *= n
    n = len(nums)
    for i in range(n - 1, -1, -1):
        r[i] *= suffix
        suffix *= nums[i]

    return r


l1 = [1, 2, 3, 4]
l2 = [-1, 1, 0, -3, 3]

print(productExceptSelf1(l1))
print(productExceptSelf1(l2))
print(productExceptSelf2(l1))
print(productExceptSelf2(l2))
