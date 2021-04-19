package user

import (
	"testing"
)

func TestFullname(t *testing.T) {
	var userParams []Option
	userParams = append(userParams, OptionUsername("bonj"))
	userParams = append(userParams, OptionFirstName("James"))
	userParams = append(userParams, OptionLastName("Bond"))

	myUser := New(userParams...)

	if myUser.Fullname() != "James Bond" {
		t.Errorf("myUser.Fullname() = %s; want James Bond", myUser.Fullname())
	}
}

func TestPasswordLength(t *testing.T) {
	var userParams []Option
	userParams = append(userParams, OptionUsername("bonj"))
	userParams = append(userParams, OptionFirstName("James"))
	userParams = append(userParams, OptionLastName("Bond"))

	myUser := New(userParams...)
	myUser.CreateRandomPassword(17)

	if len(myUser.password.plaintext) != 17 {
		t.Errorf("len(myUser.password.plaintext) = %d; want 17", len(myUser.password.plaintext))
	}
}

func BenchmarkGeneration(b *testing.B) {
	var userParams []Option
	userParams = append(userParams, OptionUsername("bonj"))
	userParams = append(userParams, OptionFirstName("James"))
	userParams = append(userParams, OptionLastName("Bond"))

	myUser := New(userParams...)

	for i := 0; i < b.N; i++ {
		myUser.CreateRandomPassword(17)
	}
}
