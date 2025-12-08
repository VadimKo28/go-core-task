package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    int
		wantValue int
		wantFound bool
	}{
		{
			name:      "add new element",
			key:       "test_key",
			value:     42,
			wantValue: 42,
			wantFound: true,
		},
		{
			name:      "add element with zero value",
			key:       "zero_key",
			value:     0,
			wantValue: 0,
			wantFound: true,
		},
	}

	m := genMap()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Add(tt.key, tt.value)

			if !m.Exists(tt.key) {
				t.Errorf("Exists(%q) = false, want true", tt.key)
			}

			gotValue, gotFound := m.Get(tt.key)
			if gotFound != tt.wantFound {
				t.Errorf("Get(%q) found = %v, want %v", tt.key, gotFound, tt.wantFound)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Get(%q) value = %d, want %d", tt.key, gotValue, tt.wantValue)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name        string
		setupKeys   []string
		setupValues []int
		removeKey   string
		wantExists  bool
		wantFound   bool
	}{
		{
			name:        "remove existing element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			removeKey:   "key1",
			wantExists:  false,
			wantFound:   false,
		},
		{
			name:        "remove non-existing element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			removeKey:   "nonexistent",
			wantExists:  false,
			wantFound:   false,
		},
		{
			name:        "remove from empty map",
			setupKeys:   []string{},
			setupValues: []int{},
			removeKey:   "any_key",
			wantExists:  false,
			wantFound:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := genMap()

			for i, key := range tt.setupKeys {
				m.Add(key, tt.setupValues[i])
			}

			m.Remove(tt.removeKey)

			gotExists := m.Exists(tt.removeKey)
			if gotExists != tt.wantExists {
				t.Errorf("Exists(%q) = %v, want %v", tt.removeKey, gotExists, tt.wantExists)
			}

			_, gotFound := m.Get(tt.removeKey)
			if gotFound != tt.wantFound {
				t.Errorf("Get(%q) found = %v, want %v", tt.removeKey, gotFound, tt.wantFound)
			}

			for i, key := range tt.setupKeys {
				if key != tt.removeKey {
					if !m.Exists(key) {
						t.Errorf("Exists(%q) = false, want true (other element should still exist)", key)
					}
					gotValue, _ := m.Get(key)
					if gotValue != tt.setupValues[i] {
						t.Errorf("Get(%q) = %d, want %d (other element value should be unchanged)", key, gotValue, tt.setupValues[i])
					}
				}
			}
		})
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name      string
		setupKeys []string
		setupValues []int
	}{
		{
			name:        "copy empty map",
			setupKeys:   []string{},
			setupValues: []int{},
		},
		{
			name:        "copy map with elements",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := genMap()

			for i, key := range tt.setupKeys {
				original.Add(key, tt.setupValues[i])
			}

			copyMap := original.Copy()

			if copyMap == original {
				t.Errorf("Copy() returned same pointer, want different pointer")
			}

			for i, key := range tt.setupKeys {
				if !copyMap.Exists(key) {
					t.Errorf("Copy() missing key %q in copy", key)
				}

				gotValue, gotFound := copyMap.Get(key)
				if !gotFound {
					t.Errorf("Get(%q) in copy found = false, want true", key)
				}
				if gotValue != tt.setupValues[i] {
					t.Errorf("Get(%q) in copy = %d, want %d", key, gotValue, tt.setupValues[i])
				}

				if !original.Exists(key) {
					t.Errorf("Original map missing key %q after copy", key)
				}

				origValue, origFound := original.Get(key)
				if !origFound {
					t.Errorf("Get(%q) in original found = false, want true", key)
				}
				if origValue != tt.setupValues[i] {
					t.Errorf("Get(%q) in original = %d, want %d", key, origValue, tt.setupValues[i])
				}
			}

			if len(tt.setupKeys) > 0 {
				testKey := tt.setupKeys[0]
				testValue := tt.setupValues[0]

				original.Add(testKey, testValue+1000)
				copyValue, _ := copyMap.Get(testKey)
				if copyValue != testValue {
					t.Errorf("Modifying original affected copy: Get(%q) in copy = %d, want %d", testKey, copyValue, testValue)
				}

				copyMap.Add(testKey, testValue+2000)
				origValue, _ := original.Get(testKey)
				if origValue != testValue+1000 {
					t.Errorf("Modifying copy affected original: Get(%q) in original = %d, want %d", testKey, origValue, testValue+1000)
				}
			}
		})
	}
}

func TestExists(t *testing.T) {
	tests := []struct {
		name        string
		setupKeys   []string
		setupValues []int
		checkKey    string
		wantExists  bool
	}{
		{
			name:        "exists in empty map",
			setupKeys:   []string{},
			setupValues: []int{},
			checkKey:    "any_key",
			wantExists:  false,
		},
		{
			name:        "exists for single element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			checkKey:    "key1",
			wantExists:  true,
		},
		{
			name:        "not exists for single element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			checkKey:    "nonexistent",
			wantExists:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := genMap()

			for i, key := range tt.setupKeys {
				m.Add(key, tt.setupValues[i])
			}

			gotExists := m.Exists(tt.checkKey)
			if gotExists != tt.wantExists {
				t.Errorf("Exists(%q) = %v, want %v", tt.checkKey, gotExists, tt.wantExists)
			}

			for _, key := range tt.setupKeys {
				if !m.Exists(key) {
					t.Errorf("Exists(%q) = false, want true (other element should exist)", key)
				}
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name        string
		setupKeys   []string
		setupValues []int
		getKey      string
		wantValue   int
		wantFound   bool
	}{
		{
			name:        "get from empty map",
			setupKeys:   []string{},
			setupValues: []int{},
			getKey:      "any_key",
			wantValue:   0,
			wantFound:   false,
		},
		{
			name:        "get existing single element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			getKey:      "key1",
			wantValue:   10,
			wantFound:   true,
		},
		{
			name:        "get non-existing single element",
			setupKeys:   []string{"key1"},
			setupValues: []int{10},
			getKey:      "nonexistent",
			wantValue:   0,
			wantFound:   false,
		},
		{
			name:        "get element with zero value",
			setupKeys:   []string{"zero_key"},
			setupValues: []int{0},
			getKey:      "zero_key",
			wantValue:   0,
			wantFound:   true,
		},
		{
			name:        "get element with negative value",
			setupKeys:   []string{"neg_key"},
			setupValues: []int{-5},
			getKey:      "neg_key",
			wantValue:   -5,
			wantFound:   true,
		},
		{
			name:        "get element with large value",
			setupKeys:   []string{"large_key"},
			setupValues: []int{999999},
			getKey:      "large_key",
			wantValue:   999999,
			wantFound:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := genMap()

			for i, key := range tt.setupKeys {
				m.Add(key, tt.setupValues[i])
			}

			gotValue, gotFound := m.Get(tt.getKey)
			if gotFound != tt.wantFound {
				t.Errorf("Get(%q) found = %v, want %v", tt.getKey, gotFound, tt.wantFound)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Get(%q) value = %d, want %d", tt.getKey, gotValue, tt.wantValue)
			}

			for i, key := range tt.setupKeys {
				gotValue, gotFound := m.Get(key)
				if !gotFound {
					t.Errorf("Get(%q) found = false, want true (other element should exist)", key)
				}
				if gotValue != tt.setupValues[i] {
					t.Errorf("Get(%q) = %d, want %d (other element value should be unchanged)", key, gotValue, tt.setupValues[i])
				}
			}
		})
	}
}
