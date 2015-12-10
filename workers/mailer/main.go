package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/facebookgo/inject"
	"github.com/facebookgo/startstop"
)

func main() {
	var graph inject.Graph
	var app Application

	err := graph.Provide(
		&inject.Object{Value: &app},
	)

	logrus.Info("Providing graph objects")

	if err != nil {
		logrus.WithError(err).Fatal("Error providing graph objects")
	}

	logrus.Info("Populating graph")

	if err := graph.Populate(); err != nil {
		logrus.WithError(err).Fatal("Error populating graph objects")
	}

	logrus.Info("Starting application")

	if err = startstop.Start(graph.Objects(), logrus.StandardLogger()); err != nil {
		logrus.WithError(err).Fatal("Error initializing objects")
	}

	logrus.Info("Tearing down application")

	if err = startstop.Stop(graph.Objects(), logrus.StandardLogger()); err != nil {
		logrus.WithError(err).Fatal("Error tearing down application")
	}
}
