// time: O(n)
// space: O(n)
function twoSum_map(numbers: number[], target: number): number[] {
  const map = new Map<number, number>();

  for (let i = 0; i < numbers.length; i++) {
    const theOtherNumber = target - numbers[i];

    if (!map.has(theOtherNumber)) {
      map.set(numbers[i], i + 1);
    } else {
      return [map.get(theOtherNumber)!, i + 1];
    }
  }
  throw new Error("");
}

// time: O(n)
// space: O(1)
function twoSum_twoPointers(numbers: number[], target: number): number[] {
  let left = 0;
  let right = numbers.length - 1;

  while (left < right) {
    if (numbers[left] + numbers[right] === target) {
      return [left + 1, right + 1];
    }
    if (numbers[left] + numbers[right] > target) {
      right--;
    }
    if (numbers[left] + numbers[right] < target) {
      left++;
    }
  }
  throw new Error("");
}
