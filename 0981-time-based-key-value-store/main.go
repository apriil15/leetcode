package main

func main() {

}

type TimeMap struct {
	m map[string][]Node
}

type Node struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]Node),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.m[key] = append(tm.m[key], Node{
		value:     value,
		timestamp: timestamp,
	})
}

// time: O(log n)
func (tm *TimeMap) Get(key string, timestamp int) string {
	nodes, ok := tm.m[key]
	if !ok {
		return ""
	}

	left := 0
	right := len(nodes) - 1

	// corner case
	if nodes[left].timestamp > timestamp {
		return ""
	}
	if nodes[right].timestamp <= timestamp {
		return nodes[right].value
	}

	resIndex := 0
	for left <= right {
		mid := left + (right-left)/2
		if nodes[mid].timestamp == timestamp {
			return nodes[mid].value
		}

		if nodes[mid].timestamp < timestamp {
			resIndex = mid // find nearest timestamp
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return nodes[resIndex].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
