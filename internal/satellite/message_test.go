package satellite

import (
	"reflect"
	"testing"
)

func TestDeleteElementFromArray_empty(t *testing.T) {
	var array = []string{}
	newarray := deleteElementFromArray(array, 1)

	if !reflect.DeepEqual(array, newarray) {
		t.Errorf("Erro to delete empty array")
	}
}

func TestDeleteElementFromArray_firstElement(t *testing.T) {
	var array = []string{"e1", "e2", "e3"}
	var expectarray = []string{"e2", "e3"}

	newarray := deleteElementFromArray(array, 0)
	if !reflect.DeepEqual(expectarray, newarray) {
		t.Errorf("Failure to delete first element from array")
	}
}

func TestDeleteElementFromArray_lastElement(t *testing.T) {
	var array = []string{"e1", "e2", "e3"}
	var expectarray = []string{"e1", "e2"}

	newarray := deleteElementFromArray(array, len(array)-1)
	if !reflect.DeepEqual(expectarray, newarray) {
		t.Errorf("Failure to delete last element from array")
	}
}

func TestDeleteElementFromArray_nthElement(t *testing.T) {
	var array = []string{"e1", "e2", "e3", "e4", "e5"}

	var expectarray1 = []string{"e1", "e3", "e4", "e5"}
	newarray := deleteElementFromArray(array, 1)
	if !reflect.DeepEqual(expectarray1, newarray) {
		t.Errorf("Failure to delete nth element from array")
	}
}

// deleteEmptyWord

func TestDeleteEmptyWord_empty(t *testing.T) {
	var array = []string{}
	newarray := deleteEmptyWord(array)

	if !reflect.DeepEqual(array, newarray) {
		t.Errorf("Erro to delete empty word from empty array")
	}
}

func TestDeleteEmptyWord_noemptywords(t *testing.T) {
	var array = []string{"e1", "e2", "e3"}

	newarray := deleteEmptyWord(array)
	if !reflect.DeepEqual(array, newarray) {
		t.Errorf("Erro to delete empty word from array which contains no empty word")
	}
}

func TestDeleteEmptyWord_singleemptyword(t *testing.T) {
	var array = []string{"e1", "e2", "", "e3"}
	var expectarray = []string{"e1", "e2", "e3"}

	newarray := deleteEmptyWord(array)
	if !reflect.DeepEqual(expectarray, newarray) {
		t.Errorf("Erro to delete empty word from array which contains single empty word")
	}
}

func TestDeleteEmptyWord_multipleemptyword(t *testing.T) {
	var array = []string{"", "e1", "e2", "", "e3", ""}
	var expectarray = []string{"e1", "e2", "e3"}

	newarray := deleteEmptyWord(array)
	if !reflect.DeepEqual(expectarray, newarray) {
		t.Errorf("Erro to delete empty word from array which contains multiple empty word")
	}
}
