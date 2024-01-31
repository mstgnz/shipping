package cargo

type Response struct {
	Code    int
	Status  string
	Message string
	Data    []byte
	Length  int64
}
