package main

import (
	_ "github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/api"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/controller"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/internal/phonebook"
	"github.com/morteza-shahrabi-farahani/golang-exercises/mastering-go/Phone-book/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PhoneBook []phonebook.Entry

const CSVFILE = "../data/data.csv"

func main() {
	// Register prometheus metrics
	metrics := metrics.RegisterMetrics()
	for _, metric := range metrics {
		prometheus.MustRegister(metric)
	}

	controller.StartHander()
}
