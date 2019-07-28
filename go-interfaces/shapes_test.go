package interfaces

import (
	"testing"

	"github.com/aszenz/basic-golang/go-test-helpers"
)

func TestPermiter(t *testing.T) {
	checkPermiter := func(t *testing.T, shape Shape, want float64, shapeName string) {
		got := shape.Perimeter()
		testhelpers.CheckNumbersEqual(t, got, want, shapeName)

	}
	permiterTests := []struct {
		shapeName     string
		shape         Shape
		wantPerimeter float64
	}{
		{shapeName: "Rectangle", shape: Rectangle{Width: 50.0, Height: 50.0}, wantPerimeter: 200.0},
		{shapeName: "Circle", shape: Circle{Radius: 1}, wantPerimeter: 6.283185307179586},
	}
	for _, tt := range permiterTests {
		t.Run(tt.shapeName, func(t *testing.T) {
			checkPermiter(t, tt.shape, tt.wantPerimeter, tt.shapeName)
		})
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64, shapeName string) {
		got := shape.Area()
		testhelpers.CheckNumbersEqual(t, got, want, shapeName)
	}
	areaTests := []struct {
		shapeName string
		shape     Shape
		wantArea  float64
	}{
		{shapeName: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, wantArea: 72.0},
		{shapeName: "Circle", shape: Circle{Radius: 10}, wantArea: 314.1592653589793},
		{shapeName: "Circle", shape: Triangle{Base: 12, Altitude: 5}, wantArea: 30},
	}

	for _, tt := range areaTests {
		t.Run(tt.shapeName, func(t *testing.T) {
			checkArea(t, tt.shape, tt.wantArea, tt.shapeName)
		})
	}
}
