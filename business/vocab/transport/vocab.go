package transport

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/labstack/echo/v4"
	"gitlab.com/sannonthachai/find-the-hidden-backend/util"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func (h *Handler) GetVocabByChapter(c echo.Context) error {
	chapter := c.QueryParam("chapter")
	chapterInt, err := strconv.Atoi(chapter)
	if err != nil {
		fmt.Println("Error strconv")
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}
	result, err := h.vocabService.GetVocabByChapter(chapterInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.CreateErrorResponse(nil, "99", "ErrorUnexpected", "ErrorUnexpected"))
	}

	return c.JSON(http.StatusOK, util.CreateSuccessResponse(result))
}

func (h *Handler) TextToSpeech(c echo.Context) error {
	// Instantiates a client.
	ctx := context.Background()
	text := c.QueryParam("text")

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		fmt.Println("Error new client: ", err)
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		fmt.Println("Error synthesizeSpeech: ", err)
	}

	// The resp's AudioContent is binary.
	// filename := "output.mp3"
	// err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	// if err != nil {
	// 	fmt.Println("Error writeFile: ", err)
	// }

	return c.Blob(http.StatusOK, "audio/mp3", resp.AudioContent)
}
