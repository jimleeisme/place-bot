package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// SET YOUR VALUES HERE
	coordX := 123
	coordY := 123
	myBearerToken := ""

	client := &http.Client{}
	bigBodyTemplate := `{"operationName":"setPixel","variables":{"input":{"actionName":"r/replace:set_pixel","PixelMessageData":{"coordinate":{"x":%d,"y":%d},"colorIndex":14,"canvasIndex":0}}},"query":"mutation setPixel($input: ActInput!) {\n act(input: $input) {\n data {\n ... on BasicMessage {\n id\n data {\n ... on GetUserCooldownResponseMessageData {\n nextAvailablePixelTimestamp\n __typename\n }\n ... on SetPixelResponseMessageData {\n timestamp\n __typename\n }\n __typename\n }\n __typename\n }\n __typename\n }\n __typename\n }\n}\n"}`
	requestBody := fmt.Sprintf(bigBodyTemplate, coordX, coordY)

	req, err := http.NewRequest("POST", "https://gql-realtime-2.reddit.com/query", bytes.NewBuffer([]byte(requestBody)))

	req.Header.Add("accept", "*/*")
	req.Header.Add("apollographql-client-name", "mona-lisa")
	req.Header.Add("apollographql-client-version", "0.0.1")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", myBearerToken))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
