package main

import "github.com/twilio/twilio-go/twiml"

var ivrGoodbye = &twiml.VoiceSay{
	Message: "Goodbye.",
}

var ivrToContinue = &twiml.VoiceSay{
	Message: "To continue, press 1. Otherwise, press 2.",
}

var ivrGreeting1 = &twiml.VoiceSay{
	Message: `Welcome to the Jack dating interactive survey experience.
    In a moment, you will be asked a brief series of questions to 
    help determine your compatibility in a potential relationship with Jack.
    Please use the keypad on your phone to answer each question.
    Your responses will not be recorded until the end of the survey experience.
    You may hang up at any time to end the survey and erase your survey responses.`,
}
