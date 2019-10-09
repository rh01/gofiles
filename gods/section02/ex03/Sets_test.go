package set

import "testing"

func TestAddElement(t *testing.T) {
	var s = &Set{}
	s.New()

	s.AddElement(5)
	if _, exist := s.integerMap[5]; !exist {
		t.Errorf("AddElement method failed, expect %d, but got %v", 5, nil)
	}
}
