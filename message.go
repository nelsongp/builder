package builder

//Message = producto concreto
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Text      string `json:"text" xml:"text"`
}

//Message Represented = representacion del objeto
type MessageRepresented struct {
	Body   []byte
	Format string
}
