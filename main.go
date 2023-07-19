package main

type Person struct {
	Src    string
	Dest   string
	Weight float64
}

type Elevator struct {
	CurrentFloor  int
	Direction     string
	CurrentWeight float64
	IsMobile      bool
	RequestFloors []int
	ButtonFloors  []int
	Passengers    map[int][]Person
}

func main() {

}
