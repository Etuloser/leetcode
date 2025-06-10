function moveZeros(nums: number[]): void {
  let slow = 0;
  for (let fast = 0; fast < nums.length; fast++) {
    if (nums[fast] !== 0) {
      nums[slow] = nums[fast];
      slow++;
    }
  }
  nums.fill(slow, 0);
}
// 快指针一直走，遇到不为零的交换到慢指针，慢指针进1，最后fill覆盖填0
