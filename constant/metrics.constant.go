package constant

type Domain string

const (
	Auth      Domain = "auth"
	CheckIn   Domain = "check_in"
	Group     Domain = "group"
	Object    Domain = "object"
	Pin       Domain = "pin"
	Selection Domain = "selection"
	User      Domain = "user"
)

func (d Domain) String() string {
	return string(d)
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)

func (m Method) String() string {
	return string(m)
}
