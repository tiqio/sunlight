package core

type Session struct {
	id int
}

func NewSession(id int) *Session {
	v := &Session{id: id}
	
}
