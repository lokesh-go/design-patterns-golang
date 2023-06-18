package facad

type securityCode struct {
	code int
}

func newSC(code int) *securityCode {
	return &securityCode{
		code: code,
	}
}

func (s *securityCode) checkSecurity(code int) bool {
	return s.code == code
}
