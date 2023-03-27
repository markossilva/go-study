package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType(windowType string) {
	b.windowType = windowType
}

func (b *IglooBuilder) SetDoorType(doorType string) {
	b.doorType = doorType
}

func (b *IglooBuilder) SetNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *IglooBuilder) GetHouse() House {
	return House{
		DoorType:   b.doorType,
		WindowType: b.windowType,
		Floor:      b.floor,
	}
}
