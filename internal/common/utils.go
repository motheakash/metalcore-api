package common

func validate_phone(phone_number string) bool {
	return len(phone_number) == 10 || len(phone_number) == 13
}
