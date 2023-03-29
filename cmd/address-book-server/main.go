package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jgoerz/go-kit-crud/internal/addressbook"
	"github.com/jgoerz/go-kit-crud/pkg/client/pb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var TestPKCS8PrivateKey string = `
-----BEGIN RSA PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCbwr4dnhiQg+g0
6Gzq/Dzdeufwai8V6ioqRJt4I2T+7FD16mBEQuFxrsgR8Tby+8k3vU6QFdTbw0M0
KcELWgOpLQf4zmHPFowGTnqq0WmtF62wueDmNCXiJQTWc8SArXHwdFugzcZU+v9A
q/IQnwqBUL0eLtsH3wsE/4Yp9RTmApaaB0gEQKwTDYfoXUbS8eOwsIitRIB7G9km
3YfmqsWgETcbEPkBZJYqEbh4EkhBrFK46Emmlj6rO/n/qRY/j3ulKIojosCIy+q+
4AQhSw9UVXjaQGT8/sW6rzdYmrAV7WQ7rLS7WTzZ4JpSMyCHe+oOvoqmbqvWBOb1
JrS1b9V3AgMBAAECggEATcKlQwgAX3Va4b7/UTjn8fJABKgeJaSntD5YF1wkOZgk
uwBtWubAwg5t13dC5X+J7wfVRt1/TM9op4wI0e/4T+cpSa9O6SHqeeOIHes6BK4D
imzhuEtkdkfg6GcXPN5aOZ79d4QDdb1w3Qp6aI3dor16DK17r6xMRgyDIEM4UbNF
tl6tRmbrEjr5pDTlsmOzmAeQ4BsYfgfCDPmCwnc2tqLPRsUFvqBOUoSJ3EID84XM
fnT+8NzR5ZsZDTiJxVdRpTEaruQABxNA43LPDy/6zQGAVlLsp3nD9gZLIHEpKufu
B/IzgHnBpyorW/sIRLeoUwdTxLTkbQZH5g1BxLlQoQKBgQDHJUZzjyl8Z0NTIl7X
gkjONOcuwhvbZapRXnEhywCO75ObrmKeR0Gk+GXk04aI7+5KJ52Guvb7nU09G/Xw
uGtC+F0XKfc2ptChfL6oD9VPMjp3bMuB65egKpzdV+K/ZEct8MI1EsB9Tc5yGWWv
pqwrMIxKj8Ci2k3ohd/GkF8F6QKBgQDIOqVab0SGb2KEhsiqgTVARQmlOTszIq4z
UfcAyCGi+xezywqdvQP3Ml5gogcDDR+6slyzIbzhGtHIQ4FkG9QfaefSTut5BT5/
lfK01CO4HQBoWgLEiKzl76MsRWIv+jxfExg6oEb7WZbK/BTwTxb4m2DYsjl5L7Xk
KFjph00EXwKBgQC/0cG4kY8uSvDoZNTh1JZ4OTDtMv9OJvEVC1kBad4Rz+ZoMGLB
fnVWiATtkmmmASWPu/TZz8ESv4OkdwhAZAK9MSnJpByBQdD3m4axrv6SGBmE6wBj
FiCooCMUeRDptZdyQtNt96/9gjJ2aMwvkuWHfG3FbA3rT0d3z2uqgWll8QKBgEre
Kt/ixPOjiGnXYAbpIzkx10ZxXOJk8E/+MOaY7oLbcmRm4kRS3b27lrB5RTft21Ra
xvCwB8j/1zsTirkc8rcASY9ItSFeRZ09OzBENkrshS9/oJNOK6Aad5/hHbKk1ZgT
MrcRIRlwyUKC+W1VlVhF+PNtyLG4lkGGmKBRWAnvAoGAMyTmCt+51qDF3P8ERX+U
C3FLysZqnEtaGB2rzOxqPI3tGT9R/PI8/zvJBjaenfF5Xs3Swy7t6tjYgsLUFHfB
H6wxyu8SdoVa0YRXBJzrbzqAuxKe35Jee9LjJZzxu2rgPLmY7WU9+KtaC0NNqdhj
+Sy/oU/bw4W4BVvBXr+PF1o=
-----END RSA PRIVATE KEY-----`

func main() {
	// Log Setup, "pretty" for DebugLevel, for InfoLevel and above (and
	// production) use JSON structured output.  Do not run DebugLevel in
	// production.  The ConsoleWriter is documented as having poor performance
	// log.Info().Msgf("Setting log level to: %v", env.GetServiceConfig().LogLevel)
	// zerolog.SetGlobalLevel(env.GetServiceConfig().LogLevel)
	// if env.GetServiceConfig().LogLevel <= zerolog.DebugLevel {
	// 	log.Logger = log.Output(
	// 		zerolog.ConsoleWriter{
	// 			Out:     os.Stdout,
	// 			NoColor: false,
	// 		},
	// 	)
	// }
	level, _ := zerolog.ParseLevel("debug")
	zerolog.SetGlobalLevel(level)

	var (
		httpAddr = flag.String("http", ":3333", "http listen address")
		gRPCAddr = flag.String("grpc", ":3334", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	repo := addressbook.NewInMemoryRepository()

	block, _ := pem.Decode([]byte(TestPKCS8PrivateKey))
	if block == nil {
		panic("Couldn't decode private key")
	}

	pKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	privateKey, ok := pKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("could not type assert private key to *rsa.PrivateKey, got: %T\n", pKey)
		panic("could not type assert private key to *rsa.PrivateKey")
	}

	srv := addressbook.NewService(repo, privateKey)
	endpoints := addressbook.MakeEndpoints(srv)
	errChan := make(chan error)

	// Trap interrupt and term signals; and send a message to errChan so we can
	// terminate.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport
	go func() {
		log.Info().Msgf("HTTP ListenAndServe '*%v'", *httpAddr)
		handler := addressbook.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// gRPC transport
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		log.Info().Msgf("gRPCServer.Serve '*%v'", *gRPCAddr)
		handler := addressbook.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterAddressBookServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	// Block on interrupt or term signal
	log.Fatal().Msgf("%v", <-errChan)
}
