package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType(windowType string) {
	b.windowType = windowType
}

func (b *NormalBuilder) SetDoorType(doorType string) {
	b.doorType = doorType
}
func (b *NormalBuilder) SetNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *NormalBuilder) GetHouse() House {
	return House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
