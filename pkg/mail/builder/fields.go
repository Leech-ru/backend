package builder

// SetFIO устанавливает ФИО клиента
func (b *EmailBuilder) SetFIO(fio string) *EmailBuilder {
	b.data.FIO = fio
	return b
}

// SetPhoneNumber устанавливает номер телефона
func (b *EmailBuilder) SetPhoneNumber(phone string) *EmailBuilder {
	b.data.PhoneNumber = phone
	return b

} // SetEmail устанавливает почту
func (b *EmailBuilder) SetEmail(email string) *EmailBuilder {
	b.data.Email = email
	return b
}

// SetAddress устанавливает адрес доставки
func (b *EmailBuilder) SetAddress(address string) *EmailBuilder {
	b.data.Address = address
	return b
}

// SetComment устанавливает комментарий клиента
func (b *EmailBuilder) SetComment(comment *string) *EmailBuilder {
	if comment == nil {
		return b
	}
	b.data.Comment = *comment
	return b
}

// SetLeechSize1 устанавливает количество пиявок размера 1
func (b *EmailBuilder) SetLeechSize1(count *int) *EmailBuilder {
	if count == nil {
		return b
	}
	b.data.LeechSize1 = *count
	return b
}

// SetLeechSize2 устанавливает количество пиявок размера 2
func (b *EmailBuilder) SetLeechSize2(count *int) *EmailBuilder {
	if count == nil {
		return b
	}
	b.data.LeechSize2 = *count
	return b
}

// SetLeechSize3 устанавливает количество пиявок размера 3
func (b *EmailBuilder) SetLeechSize3(count *int) *EmailBuilder {
	if count == nil {
		return b
	}
	b.data.LeechSize3 = *count
	return b

} // SetTotalCount устанавливает общее количество пиявок
func (b *EmailBuilder) SetTotalCount(totalCount int) *EmailBuilder {
	b.data.TotalCount = totalCount
	return b
}

// SetPackageType устанавливает тип упаковки
func (b *EmailBuilder) SetPackageType(packageType int) *EmailBuilder {
	b.data.PackageType = packageType
	return b
}

// SetTotalPrice устанавливает итоговую сумму
func (b *EmailBuilder) SetTotalPrice(price float64) *EmailBuilder {
	b.data.TotalPrice = price
	return b
}
