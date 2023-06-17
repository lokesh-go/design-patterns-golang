package builder

// Student ...
type mbaStudent struct {
	name    string
	rollNo  int
	stream  string
	college string
}

func newMBAStudentBuilder() Builder {
	return &mbaStudent{}
}

// setCollege implements Builder.
func (s *mbaStudent) SetCollege() {
	s.college = "MBA College"
}

// setName implements Builder.
func (s *mbaStudent) SetName() {
	s.name = "Pankaj"
}

// setRollNo implements Builder.
func (s *mbaStudent) SetRollNo() {
	s.rollNo = 1
}

// setStream implements Builder.
func (s *mbaStudent) SetStream() {
	s.stream = "MBA"
}
