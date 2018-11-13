package email

import "bitbucket.org/Budry/availability-checker/src/sites"

func SendFailNotificationMessage(result *sites.Result) {
	message := &Message{}
	message.From = "info@zaruba-ondrej.cz"
	message.To = []string{"info@zaruba-ondrej.cz"}
	message.Subject = "Site " + result.Site.Url + " is not working!"
	message.Body = "Site " + result.Site.Url + " is not working!\nErrors:\n"
	for _, errorMessage := range result.Errors  {
		message.Body += "\n * " + errorMessage + "\n"
	}

	Send(message)
}