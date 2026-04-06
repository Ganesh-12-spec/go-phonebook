package phonebook

type PhoneBook struct{
	Contacts map[string]*Contact
}
func NewPhoneBook() *PhoneBook{
	return &PhoneBook{
		Contacts: make(map[string]*Contact),
	}
}
func (pb *PhoneBook) Add(c *Contact){
	pb.Contacts[c.Name] = c
}
func(pb *PhoneBook) Get(name string) (*Contact, bool){
	c, ok := pb.Contacts[name]
	return c, ok
}

func (pb *PhoneBook) Delete(name string){
	delete(pb.Contacts,name)
}
func (pb *PhoneBook) Search(query string) []*Contact{
	var results []*Contacts 
	for _, c := range pb.contacts{
		if string.Contains (
			string.ToLower(c.Name),
			string.ToLower(query),{
	results = append(results,c)
}
	}
	return results
}
// All contacts list
func (pb *PhoneBook) List() {
    if len(pb.contacts) == 0 {
        fmt.Println("Phonebook khali hai")
        return
    }
    for _, c := range pb.contacts {
        fmt.Printf("%-15s %-15s %s\n", c.Name, c.Phone, c.Email)
        if len(c.Groups) > 0 {
            fmt.Printf("  Groups: %v\n", c.Groups)
        }
    }
}

// Size
func (pb *PhoneBook) Size() int {
    return len(pb.contacts)
}