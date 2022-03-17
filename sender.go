package builder

// Sender = Director
type Sender struct {
	builder MessageBuilder
}

// SetBuilder asigna el constructor
func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

// BuildMessage a concrete message via MessageBuilder
func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	//esto es lo del patron cadena (diferente a cadena de mando) que lo tenemos en el builder
	s.builder.SetRecipient(recipient).SetMessage(message)
	return s.builder.Build()
}
