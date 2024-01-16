function twoSum(nums: number[], target: number): number[] {
  const myMap: Map<number, number> = new Map();

  for (let i = 0; i < nums.length; i++) {
    let value: number = nums[i];
    let compl: number = target - value;
    if (myMap.has(compl) && myMap.get(compl) != i) {
      let r: number[] = [myMap.get(compl)!, i];
      return r;
    }
    myMap.set(value, i);
  }

  return [];
}
