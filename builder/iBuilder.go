package builder

type IBuilder interface {
	SetWindowType(windowType string)
	SetDoorType(doorType string)
	SetNumFloor(numFloor int)
	GetHouse() House
}

type Builder int64

const (
	Normal Builder = 0
	Igloo  Builder = 1
)

func (b Builder) GetBuilder() IBuilder {
	switch b {
	case Normal:
		return newNormalBuilder()
	case Igloo:
		return newIglooBuilder()
	}
	return nil
}
