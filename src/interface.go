package lists

type ListInterface interface {
	Length() int
	Append(element rune)
	Insert(element rune, index int) error
	Delete(index int) (rune, error)
	DeleteAll(element rune)
	Get(index int) (rune, error)
	Clone() ListInterface
	Reverse()
	FindFirst(element rune) int
	FindLast(element rune) int
	Clear()
	Extend(other ListInterface)
}