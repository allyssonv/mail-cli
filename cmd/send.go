package cmd

import (
	"crypto/tls"
	"fmt"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send email based on yaml file parameters",
	Long: `
	Config file example:

	from: me@mail.test
	to: someone@mail.test
	carbonCopy: other@mail.test 
	subject: Auto de Zé Limeira
	bodyContentType: text/plain
	body: |
	  No sertão sob o sol da Borborema
	  Numa terra regada a pedra e osso
	  O lagarto equilibra seu pescoço
	  Com a cauda apontando a parte extrema
	host: 127.0.0.1
	port: 1025
	user: nil
	password: nil
	`,
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}

func send() {

	from := viper.Get("from").(string)
	to := viper.Get("to").(string)
	subject := viper.Get("subject").(string)
	bodyContentType := viper.Get("bodyContentType").(string)
	body := viper.Get("body").(string)
	host := viper.Get("host").(string)
	port := viper.Get("port").(int)
	user := viper.Get("user").(string)
	password := viper.Get("password").(string)

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", viper.Get("carbonCopy").(string), viper.Get("carbonCopy").(string))
	m.SetHeader("Subject", subject)
	m.SetBody(bodyContentType, body)
	//m.Attach("/home/user/lolcat.jpg")

	d := gomail.NewDialer(host, port, user, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Send email ERROR")
		panic(err)
	}
}
