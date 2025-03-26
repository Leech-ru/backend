package order

import "sync"

// TODO включить функцию на проде
func sendOrderEmails(mailService mailService, customerEmail string, body string) error {
	return nil
	var wg sync.WaitGroup
	var err error
	var mu sync.Mutex

	recipients := []string{corporateEmail, customerEmail}
	wg.Add(len(recipients))

	for _, email := range recipients {
		go func(email string) {
			defer wg.Done()

			if sendErr := mailService.SendEmail(
				[]string{email},
				emailSubject,
				body,
			); sendErr != nil {
				mu.Lock()
				err = sendErr
				mu.Unlock()
			}
		}(email)
	}

	wg.Wait()
	return err
}
