package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var excludedResponseBodyStatuses = [...]int{100, 101, 102, 103, 204, 205, 304}

type Context struct {
	Body       string              `json:"description"`
	Request    *http.Request       `json:"-"`
	StatusCode int                 `json:"statusCode"`
	Writer     http.ResponseWriter `json:"-"`
}

func respond(context *Context) {
	statusCode := context.StatusCode
	body := context.Body

	if statusCode == 0 {
		statusCode = 200
	}

	shouldExcludeBody := shouldExcludeBody(statusCode)

	if shouldExcludeBody {
		context.Writer.Header().Del("Content-Length")
		context.Writer.Header().Del("Content-Type")
	} else {
		if jsonRequested(context) {
			context.Writer.Header().Set("Content-Type", "application/json")
			data, _ := json.Marshal(context)
			body = string(data)
		} else {
			context.Writer.Header().Set("Content-Type", "text/plain")
		}
	}

	context.Writer.WriteHeader(statusCode)

	if !shouldExcludeBody {
		context.Writer.Write([]byte(body))
	}
}

func jsonRequested(context *Context) bool {
	value := context.Request.Header.Get("Accept")
	return value == "application/json"
}

func shouldExcludeBody(statusCode int) bool {
	for _, v := range excludedResponseBodyStatuses {
		if statusCode == v {
			return true
		}
	}

	return false
}

func main() {
	router := httprouter.New()

	// home page
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		context := &Context{
			Body:       "Request a desired status code response by adding the status code at the end of the url.",
			StatusCode: 200,
			Request:    request,
			Writer:     writer,
		}
		respond(context)
	})

	// dynamic status page
	router.GET("/:statusCode", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		statusCodeFromParams := params.ByName("statusCode")
		statusCode, _ := strconv.Atoi(statusCodeFromParams)
		statusText := http.StatusText(statusCode)

		context := &Context{
			Request: request,
			Writer:  writer,
		}

		if statusText == "" {
			context.Body = "Status code not found!"
			context.StatusCode = 404
		} else {
			context.Body = statusCodeFromParams + " " + statusText
			context.StatusCode = statusCode
		}

		respond(context)
	})

	fmt.Println("*************")
	fmt.Println("Server running at http://localhost:8080")
	fmt.Println("*************")
	http.ListenAndServe(":8080", router)
}
