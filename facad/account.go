package facad

type account struct {
	id string
}

func newAccount(id string) *account {
	return &account{
		id: id,
	}
}

func (a *account) verifyAccount(id string) bool {
	return a.id == id
}
