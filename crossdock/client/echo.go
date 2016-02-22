package client

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/yarpc/yarpc-go"
	"github.com/yarpc/yarpc-go/crossdock/server"
	"github.com/yarpc/yarpc-go/encoding/json"
	"github.com/yarpc/yarpc-go/transport"
	"github.com/yarpc/yarpc-go/transport/http"

	"golang.org/x/net/context"
)

// EchoBehavior asserts that a server response is the same as the request
func EchoBehavior(addr string) Result {
	yarpc := yarpc.New(yarpc.Config{
		Name: "client",
		Outbounds: transport.Outbounds{
			"yarpc-test": http.Outbound(fmt.Sprintf("http://%v:8081", addr)),
		},
	})
	client := json.New(yarpc.Channel("yarpc-test"))
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	var response server.EchoResponse
	token := randString(5)

	_, err := client.Call(
		ctx,
		// TODO get Service from client
		&json.Request{
			Caller:    "client",
			Service:   "yarpc-test",
			Procedure: "echo",
			Body:      &server.EchoRequest{Token: token},
			TTL:       3 * time.Second,
		},
		&response,
	)

	if err != nil {
		return Result{
			Passed:  false,
			Message: fmt.Sprintf("Got err: %v", err),
		}
	}
	if response.Token != token {
		return Result{
			Passed:  false,
			Message: fmt.Sprintf("Got %v, wanted %v", response.Token, token),
		}
	}
	return Result{
		Passed:  true,
		Message: fmt.Sprintf("Server said: %v", response.Token),
	}
}

func randString(length int64) string {
	bs, err := ioutil.ReadAll(io.LimitReader(rand.Reader, length))
	if err != nil {
		panic(err)
	}
	return base64.RawStdEncoding.EncodeToString(bs)
}