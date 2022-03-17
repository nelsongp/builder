package builder

//MessageBuilder = Builder
type MessageBuilder interface {
	//En este caso estamos retornando la misma interface para poder establecer el patron de cadena, verlo en el main
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	Build() (*MessageRepresented, error)
}
