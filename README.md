Abstract API Go SDK

https://www.abstractapi.com/

Implemented: Phone Validation and Email Validation.

Example usage:

```go
func main() {
	client := datavalidation.NewClient(<your_api_key_for_phone_validation>,
            your_api_key_for_email_validation)

	pv, err := client.PhoneValidation(<phone_number>)
	if err != nil {
		log.Fatal(err)
	}

	ev, err := client.EmailValidation(<email>)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pv)
	fmt.Println(ev)
}
