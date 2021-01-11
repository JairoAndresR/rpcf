package fixtures

func GetInvalidRegisterRequest() map[string]interface{} {
	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "",
		"names":    "John Doe",
	}
}

func GetRegisterRequest() map[string]interface{} {
	return map[string]interface{}{
		"email":    "test@gmail.com",
		"password": "1234567",
		"names":    "John Doe",
	}
}
