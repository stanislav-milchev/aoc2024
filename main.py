def isSafe(nums: list[int], inc: bool) -> bool:
    # Base case: If there are fewer than 2 elements, it's "safe" by default
    if len(nums) < 2:
        return True

    # Check the first pair of elements
    if (nums[1] - nums[0] > 2) or (nums[0] - nums[1] > 2):
        return False  # Absolute difference is greater than 2

    if inc and nums[0] >= nums[1]:
        return False  # Not increasing
    if not inc and nums[0] <= nums[1]:
        return False  # Not decreasing

    # Recursive case: Check the rest of the list
    return isSafe(nums[1:], inc)
