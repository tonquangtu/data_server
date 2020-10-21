package service

type PasswordEncoder interface {
	Encode(password string) string
}

type passwordEncoder struct {
}

func (encoder *passwordEncoder) Encode(password string) string {
	return password
}

func NewPasswordEncoder() PasswordEncoder {
	return &passwordEncoder{}
}
