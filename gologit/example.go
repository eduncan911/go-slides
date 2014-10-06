// +build OMIT

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Superman struct {
	EyesOpened bool
}

type XrayVision struct{}

var Log = new(logger)

// START OMIT
func (s *Superman) Xray(person string) *XrayVision {
	Log.Infof("Xray() starting for '%s'", person)

	if !s.EyesOpened {
		Log.Info("Opening eyes")
		s.OpenEyes()
		Log.Debug("Eyes were opened")
	}

	xv, err := s.beginXrayOn(person)
	if err != nil {
		Log.Fatalf("beginXrayOn() for '%s' returned error: %s", person, err)
		os.Exit(1)
	}

	Log.Debug("Xray() completed.")
	return xv
}

// STOP OMIT

func (s *Superman) beginXrayOn(person string) (*XrayVision, error) {
	return nil, errors.New(fmt.Sprintf("%s was wearing lead armor!", person))
	//return new(XrayVision), nil
}

func (s *Superman) OpenEyes() {
	s.EyesOpened = true
}

func main() {

	s := new(Superman)
	s.Xray("Lex Luthor")

}

type logger struct{}

func (l *logger) Debug(message string) {
	fmt.Println(strings.Join([]string{"DEBUG 490a82f6 ", message}, ""))
}
func (l *logger) Debugf(message string, args ...interface{}) {
	fmt.Printf(strings.Join([]string{"DEBUG 490a82f6 ", message, "\n"}, ""), args...)
}
func (l *logger) Info(message string) {
	fmt.Println(strings.Join([]string{"INFO 490a82f6 ", message}, ""))
}
func (l *logger) Infof(message string, args ...interface{}) {
	fmt.Printf(strings.Join([]string{"INFO 490a82f6 ", message, "\n"}, ""), args...)
}
func (l *logger) Fatal(message string) {
	fmt.Println(strings.Join([]string{"FATAL 490a82f6 ", message}, ""))
}
func (l *logger) Fatalf(message string, args ...interface{}) {
	fmt.Printf(strings.Join([]string{"FATAL 490a82f6 ", message, "\n"}, ""), args...)
}
