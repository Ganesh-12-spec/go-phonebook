package phonebook

type Contact struct {
	Name string
	Phone string
	Email string
	Groups []string //slice -muliple group mein ho sakta hai
}
func (c *Contact) UpdatePhone(newPhone string){
	c.Phone = newPhone
}
func (c *Contact) AddGroup(group string){

	for _,g := range c.Groups{
		if g == group{
			return
		}
	}
	c.Groups = append(c.Groups, group)
}