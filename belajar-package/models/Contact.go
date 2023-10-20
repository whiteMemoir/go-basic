package models

type Contact struct {
	Phone			string
	Instagram	string
}

func NewContact(phone, instragram string) *Contact {
	return &Contact{
		Phone: phone,
		Instagram: instragram,
	}
}