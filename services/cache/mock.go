package services

import "fmt"

type MockedCache struct {
	data map[string]string
}

func NewMockedCache() *MockedCache {
	return &MockedCache{
		data: make(map[string]string),
	}
}

func (m *MockedCache) Get(key string) (string, error) {
	if value, ok := m.data[key]; ok {
		return value, nil
	}

	return "", fmt.Errorf("key %s not found", key)
}

func (m *MockedCache) Set(key string, value string) error {
	m.data[key] = value
	return nil
}
