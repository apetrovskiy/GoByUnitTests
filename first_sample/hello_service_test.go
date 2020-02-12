package main

import (
	"fmt"
	"github.com/cucumber/godog"
)

type GreetingData struct {
	username string
	language string
}

var g GreetingData

func iGreetUser(username string) error {
	g.username = username
	return nil
}

func iSayHelloIn(language string) error {
	g.language = language
	return nil
}

func thePhraseIs(expectedResult string) error {
	actualResult := hello(g.username, g.language)
	if expectedResult != actualResult {
		return fmt.Errorf("expected greeting %s, actual greeting %s", expectedResult, actualResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I greet user "(.*)"$`, iGreetUser)
	s.Step(`^I say Hello in "(.*)"$`, iSayHelloIn)
	s.Step(`^the phrase is "(.*)"$`, thePhraseIs)
}
