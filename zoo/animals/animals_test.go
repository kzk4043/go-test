package animals

import "testing"

func TestElephantFeed(t *testing.T) {
	expected := "Grass"
	result := ElephantFeed()
	if result != expected {
		t.Errorf("ElephantFeed() = %v; want %v", result, expected)
	}
}

func TestMonkeyFeed(t *testing.T) {
	expected := "Banana"
	result := MonkeyFeed()
	if result != expected {
		t.Errorf("MonkeyFeed() = %v; want %v", result, expected)
	}
}

func TestRabbitFeed(t *testing.T) {
	expected := "Carrot"
	result := RabbitFeed()
	if result != expected {
		t.Errorf("RabbitFeed() = %v; want %v", result, expected)
	}
}