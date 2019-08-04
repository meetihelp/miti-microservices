package Mail

import(
	"net/smtp"
	"log"
)

func Send_email(email string,url string){
	auth := smtp.PlainAuth("", "kumar93sunny91@iitkgp.ac.in", "01.14.2000", "iitkgpmail.iitkgp.ac.in")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"gjhakgp@gmail.com"}
	msg := []byte("To: "+email+"\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		url+".\r\n")
	err := smtp.SendMail("iitkgpmail.iitkgp.ac.in:25", auth, "kumar93sunny91@iitkgp.ac.in", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}