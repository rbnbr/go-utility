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

// GetValuesOfMap
// Returns a slice containing all values of the map someMap
func GetValuesOfMap[K comparable, V any](someMap map[K]V) []V {
	values := make([]V, len(someMap))

	i := 0
	for k := range someMap {
		values[i] = someMap[k]
		i++
	}

	return values
}
