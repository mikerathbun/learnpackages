package packageutil

import "fmt"

// Person is a test struct
type Person struct {
	Name string
	Sex  string
}

func (p Person) String() string {
	return fmt.Sprintf("The person's name is %s and the sex is %s.", p.Name, p.Sex)
}
