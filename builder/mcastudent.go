package builder

// Student ...
type mcaStudent struct {
	name    string
	rollNo  int
	stream  string
	college string
}

func newMCAStudentBuilder() Builder {
	return &mcaStudent{}
}

// setCollege implements Builder.
func (s *mcaStudent) SetCollege() {
	s.college = "MCA College"
}

// setName implements Builder.
func (s *mcaStudent) SetName() {
	s.name = "Ganesh"
}

// setRollNo implements Builder.
func (s *mcaStudent) SetRollNo() {
	s.rollNo = 2
}

// setStream implements Builder.
func (s *mcaStudent) SetStream() {
	s.stream = "MCA"
}
