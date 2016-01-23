package config

const (
	// Validation
	EMAIL_INVALID string = "Veuillez entrer une adresse email valide."
	EMAIL_EMPTY   string = "Veuillez entrer votre adresse email."
	NAME_EMPTY    string = "Veuillez entrer votre nom dans les champs vierges."
	PHONE_EMPTY   string = "Veuillez entrer votre numéro de téléphone."
	MESSAGE_EMPTY string = "Veuillez entrer un message."
	THANK_YOU     string = "Merci pour votre demande! Je vais vous contacter bientôt."

	// Errors
	MAIL_NOT_SENT_ERR string = "Could not connect to mail server. Please try again soon.!"

	// Emails
	CONTACT_EMAIL string = `
		<table style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 100%%; line-height: 1.6em; width: 100%%; margin: 0; padding: 0;">
			<tr style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 100%%; line-height: 1.6em; margin: 0; padding: 0;">
				<td style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 100%%; line-height: 1.6em; margin: 0; padding: 0;">

					<h1 style="font-family: 'Helvetica Neue', Helvetica, Arial, 'Lucida Grande', sans-serif; font-size: 36px; line-height: 1.2em; color: #111111; font-weight: 200; margin: 40px 0 10px; padding: 0;">
						New Message!
					</h1>

					<p style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 1.6em; font-weight: normal; margin: 0 0 10px; padding: 0;">
						<strong>Name: </strong>%s
					</p>

					<p style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 14px; line-height: 1.6em; font-weight: normal; margin: 0 0 10px; padding: 0;">
						<strong>Email: </strong>%s
					</p>

					<p style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 14px; line-height: 1.6em; font-weight: normal; margin: 0 0 10px; padding: 0;">
						<strong>Phone: </strong>%s
					</p>

					<p style="font-family: 'Helvetica Neue', 'Helvetica', Helvetica, Arial, sans-serif; font-size: 16px; line-height: 1.6em; font-weight: normal; margin: 0 0 10px; padding: 0;">
						<strong>Message: </strong>%s
					</p>

				</td>
			</tr>
        </table>
	`
)
