package builder

// SetRecipient устанавливает имя получателя
func (b *EmailBuilder) SetRecipient(name string) *EmailBuilder {
	b.data.RecipientName = name
	return b
}

// SetSubject устанавливает тему письма
func (b *EmailBuilder) SetSubject(subject string) *EmailBuilder {
	b.data.Subject = subject
	return b
}

// SetMessage устанавливает основное сообщение
func (b *EmailBuilder) SetMessage(msg string) *EmailBuilder {
	b.data.Message = msg
	return b
}
