package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aymerick/raymond"
	"github.com/caarlos0/env/v6"
)

type Environments struct {
	Workflow   string `env:"GITHUB_WORKFLOW"`
	Action     string `env:"GITHUB_ACTION"`
	Actor      string `env:"GITHUB_ACTOR"`
	Repository string `env:"GITHUB_REPOSITORY"`
	Commit     string `env:"GITHUB_SHA"`
	EventName  string `env:"GITHUB_EVENT_NAME"`
	EventPath  string `env:"GITHUB_EVENT_PATH"`
	Ref        string `env:"GITHUB_REF"`
}

type Inputs struct {
	Key     string `env:"INPUT_KEY"`
	Type    string `env:"INPUT_TYPE" envDefault:"text"`
	Content string `env:"INPUT_CONTENT"`
	Status  string `env:"INPUT_STATUS"`
}

type Content struct {
	Content string `json:"content"`
}

type Message struct {
	Type     string   `json:"msgtype"`
	Text     *Content `json:"text,omitempty"`
	Markdown *Content `json:"markdown,omitempty"`
}

func main() {
	var err error

	environments := Environments{}
	if err := env.Parse(&environments); err != nil {
		fmt.Printf("failed to parse envrionments: %s\n", err)
		os.Exit(1)
	}

	inputs := Inputs{}
	if err := env.Parse(&inputs); err != nil {
		fmt.Printf("failed to parse inputs: %s\n", err)
		os.Exit(1)
	}

	ctx := map[string]interface{}{
		"github": environments,
		"inputs": inputs,
	}

	content := &Content{}
	content.Content, err = raymond.Render(inputs.Content, ctx)
	if err != nil {
		fmt.Printf("failed to render message content: %s\n", err)
		os.Exit(1)
	}

	message := Message{}
	switch inputs.Type {
	case "text":
		message.Text = content
	case "markdown":
		message.Markdown = content
	default:
		fmt.Printf("invalid message type %s\n", inputs.Type)
		os.Exit(1)
	}
	message.Type = inputs.Type

	err = send(inputs.Key, message)
	if err != nil {
		fmt.Printf("failed to send message: %s", err)
		os.Exit(1)
	}
}

func send(key string, message Message) error {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(message)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", key)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	status := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("resposne %d: %s", status, string(body))

	return nil
}
