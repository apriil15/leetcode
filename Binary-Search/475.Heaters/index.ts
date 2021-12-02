function findRadius(houses: number[], heaters: number[]): number {
  houses.sort((a, b) => a - b)
  heaters.sort((a, b) => a - b)

  let result = 0

  for (const house of houses) {
    if (house < heaters[0]) {
      result = Math.max(result, heaters[0] - house)
      continue
    }
    if (house > heaters[heaters.length - 1]) {
      result = Math.max(result, house - heaters[heaters.length - 1])
      continue
    }
    if (heaters.includes(house)) {
      continue
    }

    const index = getBiggerNumberIndex(heaters, house)
    const min = Math.min(heaters[index] - house, house - heaters[index - 1])

    result = Math.max(result, min)
  }

  return result
}

function getBiggerNumberIndex(heaters: number[], house: number): number {
  let left = 0
  let right = heaters.length - 1

  while (left <= right) {
    const mid = Math.floor((left + right) / 2)

    if (heaters[mid] < house) {
      if (heaters[mid + 1] > house) {
        return mid + 1
      } else {
        left = mid + 2
      }
    } else {
      right = mid - 1
    }
  }

  return left
}
