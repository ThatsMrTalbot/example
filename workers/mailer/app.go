package main

import (
	"net/mail"

	"github.com/ThatsMrTalbot/example"
	"github.com/ThatsMrTalbot/example/queues"
	"github.com/Sirupsen/logrus"
	"github.com/jeffail/tunny"
)

// NOTE: Must be run from the example directory in order to pick up the config

const (
	// NumWorkers is the number of concurrent workers
	NumWorkers = 10
)

// Application is the main enrty point for the program
type Application struct {
	Config *example.Config `inject:""`
	Queues *example.Queues `inject:""`
	Mailer *Mailer         `inject:""`
}

// Start application
func (app *Application) Start() error {
	logrus.Info("Creating worker pool")

	pool, err := tunny.CreatePool(NumWorkers, func(object interface{}) interface{} {
		input, ok := object.(*queues.MailJob)

		if !ok {
			logrus.Errorf("Incorrect type, expecting MailJob got %T", input)
			return nil
		}

		logrus.Infof("Sending email to %s", input.Email)

		client := app.Mailer.Client()
		email := client.NewMail()
		email.AddRecipient(&mail.Address{Address: input.Email})
		email.SetHTML("Some content here")

		err := client.Send(email)

		if err != nil {
			logrus.Errorf("Error sending email: %s", err.Error())
		}

		return nil
	}).Open()

	if err != nil {
		logrus.Errorf("Error starting worker pool: %s", err.Error())
	}

	defer pool.Close()

	consumer, err := app.Queues.MailJobQueue.Consumer()
	if err != nil {
		logrus.Fatalf("Error getting queue consumer: %s", err.Error())
	}

	for {
		job, err := consumer.Get()

		if err != nil {
			logrus.Errorf("Error getting item from queue: %s", err.Error())
			continue
		}

		pool.SendWorkAsync(job, nil)
	}
}
