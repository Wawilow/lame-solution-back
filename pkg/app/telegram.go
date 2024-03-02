package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", os.Getenv("TG_TOKEN"))
}

func getUpdates() {
	r, err := http.Get(fmt.Sprintf("%s/getUpdates", getUrl()))
	if err != nil || r.Status != "200 OK" {
		log.Fatal(err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))
}

func SendMessageToChat(text string) error {
	var url = fmt.Sprintf("%s/sendMessage", getUrl())
	body, err := json.Marshal(map[string]string{
		"chat_id": os.Getenv("TG_CHAT"),
		"text":    text,
	})
	if err != nil {
		return err
	}

	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Println(string(b))

	if err != nil || r.StatusCode != 200 {
		return errors.New(fmt.Sprintf("post request error: %s", err))
	}
	return nil
}
