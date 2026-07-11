package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/twilio/twilio-go/twiml"
)

func askQuestion(context *gin.Context, logger *zerolog.Logger, question string, answerList []string, action string) (e error) {

	builder := strings.Builder{}
	optionCounter := 1
	for _, answer := range answerList {
		builder.WriteString("Press " + strconv.Itoa(optionCounter) + " for " + answer)
	}

	userResponse, err := twiml.Voice([]twiml.Element{
		&twiml.VoiceGather{
			Action:      action,
			Method:      "GET",
			FinishOnKey: "",
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: builder.String(),
				},
			},
		},
		&twiml.VoiceSay{
			Message: "We didn't receive any input, so your survey will now be discontinued. Goodbye!",
		},
	})

	if err != nil {
		logger.Error().Err(err).Msg("will return http 500")
		context.String(http.StatusInternalServerError, err.Error())
		return err
	} else {
		context.Header("Content-Type", "text/xml")
		context.String(http.StatusOK, userResponse)
		return nil
	}
}

func playMessage(context *gin.Context, logger *zerolog.Logger, message string) {

	voicePrompt, err := twiml.Voice([]twiml.Element{&twiml.VoiceSay{
		Message: message,
	}})

	if err != nil {
		logger.Error().Err(err).Msg("will return http 500")
		context.String(http.StatusInternalServerError, err.Error())
	} else {
		context.Header("Content-Type", "text/xml")
		context.String(http.StatusOK, voicePrompt)
	}

}

func confirmContinue(context *gin.Context, logger *zerolog.Logger, action string) (e error) {

	userResponse, err := twiml.Voice([]twiml.Element{
		&twiml.VoiceGather{
			Action:      action,
			Method:      "GET",
			FinishOnKey: "",
			InnerElements: []twiml.Element{
				&twiml.VoiceSay{
					Message: "To continue, press 1. Otherwise, press 2.",
				},
			},
		},
		&twiml.VoiceSay{
			Message: "We didn't receive any input, so your survey will now be discontinued. Goodbye!",
		},
	})

	if err != nil {
		logger.Error().Err(err).Msg("will return http 500")
		context.String(http.StatusInternalServerError, err.Error())
		return err
	} else {
		context.Header("Content-Type", "text/xml")
		context.String(http.StatusOK, userResponse)
		return nil
	}
}
