package phonebook

type PhoneBook struct{
	Contacts map[string]*Contact
}
func NewPhoneBook() *PhoneBook{
	return &PhoneBook{
		Contacts: make(map[string]*Contact),
	}
}