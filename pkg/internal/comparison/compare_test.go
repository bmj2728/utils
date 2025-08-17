package comparison

import "testing"

var (
	slStr1 = []string{"Hello", "world"}
	slStr2 = []string{"Hello", "world"}
	slStr3 = []string{"world", "Hello"}
	slStr4 = []string{"Hello", "world", "hello"}
	slStr5 []string

	slInt1 = []int{1, 2, 3}
	slInt2 = []int{1, 2, 3}
	slInt3 = []int{3, 2, 1}
	slInt4 = []int{1, 2, 3, 4}
	slInt5 []int

	slFloat1 = []float32{1.0, 2.0, 3.0}
	slFloat2 = []float32{1.0, 2.0, 3.0}
	slFloat3 = []float32{1.111111111111111, 2.222222222222222, 3.333333333333333}
	slFloat4 = []float32{1.111111111111111, 2.222222222222222, 3.333333333333333}
	slFloat5 = []float32{3.0, 2.0, 1.0}
	slFloat6 = []float32{1.0, 2.0, 3.0, 4.0}
	slFloat7 []float32
)

func TestCompareSlices(t *testing.T) {
	if CompareSlices(slStr1, slStr2, false) != true {
		t.Errorf("1 Expected true, got false")
	}
	if CompareSlices(slStr1, slStr3, false) != true {
		t.Errorf("2 Expected true, got false")
	}
	if CompareSlices(slStr1, slStr4, false) != false {
		t.Errorf("3 Expected false, got true")
	}
	if CompareSlices(slStr1, slStr5, false) != false {
		t.Errorf("4 Expected false, got true")
	}
	if CompareSlices(slStr1, slStr1, false) != true {
		t.Errorf("5 Expected true, got false")
	}
	if CompareSlices(slStr5, slStr5, true) != true {
		t.Errorf("6 Expected true, got false")
	}
	if CompareSlices(slStr5, slStr5, false) != false {
		t.Errorf("7 Expected false, got true")
	}

	if CompareSlices([]interface{}{"Hello", "world"}, []interface{}{1, 2, 3}, false) != false {
		t.Errorf("8 Expected true, got false")
	}

	if CompareSlices(slInt1, slInt1, false) != true {
		t.Errorf("9 Expected true, got false")
	}
	if CompareSlices(slInt1, slInt2, false) != true {
		t.Errorf("10 Expected true, got false")
	}
	if CompareSlices(slInt1, slInt3, false) != true {
		t.Errorf("11 Expected true, got false")
	}
	if CompareSlices(slInt1, slInt4, false) != false {
		t.Errorf("12 Expected true, got false")
	}

	if CompareSlices(slInt1, slInt5, false) != false {
		t.Errorf("13 Expected true, got false")
	}
	if CompareSlices(slInt5, slInt5, true) != true {
		t.Errorf("14 Expected true, got false")
	}
	if CompareSlices(slInt5, slInt5, false) != false {
		t.Errorf("15 Expected true, got false")
	}

	if CompareSlices(slFloat1, slFloat1, false) != true {
		t.Errorf("16 Expected true, got false")
	}
	if CompareSlices(slFloat1, slFloat2, false) != true {
		t.Errorf("17 Expected true, got false")
	}
	if CompareSlices(slFloat1, slFloat3, false) != false {
		t.Errorf("18 Expected false, got true")
	}
	if CompareSlices(slFloat1, slFloat4, false) != false {
		t.Errorf("19 Expected true, got false")
	}
	if CompareSlices(slFloat1, slFloat5, false) != true {
		t.Errorf("20 Expected true, got false")
	}
	if CompareSlices(slFloat1, slFloat6, false) != false {
		t.Errorf("21 Expected true, got false")
	}
	if CompareSlices(slFloat1, slFloat7, false) != false {
		t.Errorf("22 Expected true, got false")
	}
	if CompareSlices(slFloat7, slFloat7, true) != true {
		t.Errorf("23 Expected true, got false")
	}
	if CompareSlices(slFloat7, slFloat7, false) != false {
		t.Errorf("24 Expected false, got true")
	}
	if CompareSlices(slFloat3, slFloat4, true) != true {
		t.Errorf("25 Expected true, got false")
	}
	if CompareSlices([]rune{'w', 'o', 'r', 'd'}, []rune{'o', 'd', 'w', 'r'}, false) != true {
		t.Errorf("26 Expected true, got false")
	}
}
