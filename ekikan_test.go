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

func Test_KyoriWoHyoji(t *testing.T) {
	expected := "茗荷谷駅と新大塚駅までは1.2kmです"
	actual := KyoriWoHyoji("myogadani", "shinotsuka")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}

func Test_KyoriWoHyoji_Not_Connected(t *testing.T) {
	expected := "茗荷谷駅と千駄木駅はつながっていません"
	actual := KyoriWoHyoji("myogadani", "sendagi")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}

func Test_KyoriWoHyoji_DoNot_Exist(t *testing.T) {
	expected := "I dont existという駅は存在しません"
	actual := KyoriWoHyoji("myogadani", "I dont exist")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}

	actual = KyoriWoHyoji("I dont exist", "myogadani")
	if actual != expected {
		t.Errorf("got %v instead of %v", actual, expected)
	}
}