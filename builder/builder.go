package builder

// Builder ...
type Builder interface {
	SetName()
	SetRollNo()
	SetStream()
	SetCollege()
}

// NewStudentBuilder ...
func NewStudentBuilder(builderType string) Builder {
	switch builderType {
	case "mca":
		{
			return newMCAStudentBuilder()
		}
	case "mba":
		{
			return newMBAStudentBuilder()
		}
	}
	return nil
}
