package main

import (
	"testing"
	"math"
)

func TestGlobalEkikan_GetEkikanKyori(t *testing.T) {
	sut := &GlobalEkikan{ekikanList: Ekikan_List}
	expected := 1.2
	actual := sut.GetEkikanKyori("茗荷谷", "新大塚")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}

func Test_GetEkikanKyori_In_ReveseOrder(t *testing.T) {
	sut := &GlobalEkikan{ekikanList: Ekikan_List}
	expected := 1.2
	actual := sut.GetEkikanKyori("新大塚", "茗荷谷")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}

func Test_GetEkikanKyori_Between_Directy_Not_Connected(t *testing.T) {
	sut := &GlobalEkikan{ekikanList: Ekikan_List}
	expected := math.Inf(+1)
	actual := sut.GetEkikanKyori("新大塚", "Im Not Connected")
	if actual != expected {
		t.Error("got non infinite number")
	}
}