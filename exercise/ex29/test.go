package ex29

import "fmt"

func Test() {
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "email is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")))                               // CreateUserRequest{}, "password is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}")))                               // CreateUserRequest{}, "password confirmation is required"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")))                          // CreateUserRequest{}, "password does not match with the confirmation"
	fmt.Println(DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}"))) // CreateUserRequest{Email:"email@test.com", Password:"passwordtest", PasswordConfirmation:"passwordtest"}, nil
}
