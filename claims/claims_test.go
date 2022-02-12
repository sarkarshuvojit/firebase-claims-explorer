package claims

import "testing"

func TestClaim(t *testing.T) {
	actual := GetName()
	expected := "Shuvojit"

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
