package maps

// GetKeysOfMap
// Returns the keys contained in someMap
func GetKeysOfMap[K comparable, V any](someMap map[K]V) []K {
	keys := make([]K, len(someMap))

	i := 0
	for k := range someMap {
		keys[i] = k
		i++
	}

	return keys
}
