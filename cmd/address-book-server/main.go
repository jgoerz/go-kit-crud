package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jgoerz/go-kit-crud/internal/addressbook"
	"github.com/jgoerz/go-kit-crud/internal/appconfig"
	"github.com/jgoerz/go-kit-crud/pkg/client/pb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {

	// Get configuration
	env, problems := appconfig.NewEnvironment()
	if len(problems) != 0 {
		for _, err := range problems {
			log.Err(err).Msg("")
		}
		os.Exit(1)
	}

	// Log Setup, "pretty" for DebugLevel, for InfoLevel and above (and
	// production) use JSON structured output.  Do not run DebugLevel in
	// production.  The ConsoleWriter is documented as having poor performance
	log.Info().Msgf("Setting log level to: %v", env.ServiceConfig().LogLevel)
	zerolog.SetGlobalLevel(env.ServiceConfig().LogLevel)
	if env.ServiceConfig().LogLevel <= zerolog.DebugLevel {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	}

	// FIXME This is a little awkward as we still have to set the env vars if we
	// want to override it with or use a CLI parameter.
	var (
		httpAddr = flag.String("http", env.ServiceConfig().ListenPortHTTP, "http listen address")
		gRPCAddr = flag.String("grpc", env.ServiceConfig().ListenPortgRPC, "gRPC listen address")
	)
	flag.Parse()

	ctx := context.Background()

	repo := addressbook.NewInMemoryRepository()

	srv := addressbook.NewService(repo, env.ServiceConfig().PrivateKey)
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
