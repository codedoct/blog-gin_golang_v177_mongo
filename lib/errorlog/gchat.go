package errorlog

import (
	"encoding/json"

	"blog-gin_golang_v177_mongo/lib/env"
	"blog-gin_golang_v177_mongo/lib/httpcli"
)

type ErrorData struct {
	Error  string
	Params interface{}
	Method string
	Path   string
	Host   string
	Header interface{}
}

func SendErrorLogGChat(data *ErrorData) error {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	dataParams, _ := json.Marshal(data.Params)
	dataHeaders, _ := json.Marshal(data.Header)
	body := "*Error:* " + data.Error + "\n" +
		"*Params:* " + string(dataParams) + "\n" +
		"*Method:* " + data.Method + "\n" +
		"*Path:* " + data.Path + "\n" +
		"*Host:* " + data.Host + "\n" +
		"*Header:* " + string(dataHeaders)

	message := map[string]string{"text": body}
	webhookUrl := env.String("GCHAT_WEBHOOK", "")
	if webhookUrl != "" {
		httpClient := httpcli.HttpClient{
			ReqHeaders: headers,
			URL:        webhookUrl,
			ReqBody:    message,
		}
		_, err := httpcli.Post(&httpClient)
		if err != nil {
			return err
		}
	}

	return nil
}
