package types

type AuthConfiguration struct {
	Type         string
	Username     string
	Password     string
	PrivateKey   string
	SkipInsecure bool `json:"skipInsecure"`
}
