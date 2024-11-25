package store

import "sync"

var (
	receipts = make(map[string]int)
	mu       sync.Mutex
)

func SaveReceipt(id string, points int) {
	mu.Lock()
	defer mu.Unlock()
	receipts[id] = points
}

func GetPoints(id string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	points, exists := receipts[id]
	return points, exists
}
