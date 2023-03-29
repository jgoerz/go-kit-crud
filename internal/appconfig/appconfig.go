package appconfig

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

type ServiceConfig struct {
	ListenPortHTTP string
	ListenPortgRPC string
	LogLevel       zerolog.Level
	PrivateKey     *rsa.PrivateKey
	PublicKey      *rsa.PublicKey
}

type Environment struct {
	serviceConfig *ServiceConfig
	// dbReader      *DBConfig
	// dbWriter      *DBConfig
	// redis         *RedisConfig
	// amqp          *AMQPConfig
}

func NewEnvironment() (env *Environment, err []error) {
	env = &Environment{}
	err = env.initEnvironment()
	return env, err
}

func (e *Environment) ServiceConfig() *ServiceConfig {
	return e.serviceConfig
}

func (e *Environment) initEnvironment() (errs []error) {
	// Reference in case there are ParseTime "issues"
	// https://stackoverflow.com/questions/29341590/how-to-parse-time-from-database/29343013#29343013

	var (
		ok                                             bool
		portHTTP, portgRPC, logLevelStr, privateKeyStr string
		logLevel                                       zerolog.Level
		privateKey                                     *rsa.PrivateKey
		publicKey                                      *rsa.PublicKey
	)

	//
	// Application configuration
	//
	if portHTTP, ok = os.LookupEnv("ADDRESSBOOK_HTTP_PORT"); !ok || portHTTP == "" {
		errs = append(errs, errors.New("required ADDRESSBOOK_HTTP_PORT not set or empty"))
	}
	if _, e := strconv.Atoi(portHTTP); e == nil {
		// just checking to see if is a valid integer
		portHTTP = ":" + portHTTP
	} else {
		errs = append(errs, fmt.Errorf("required ADDRESSBOOK_HTTP_PORT is not a valid integer got: '%v'", portHTTP))
	}

	if portgRPC, ok = os.LookupEnv("ADDRESSBOOK_GRPC_PORT"); !ok || portgRPC == "" {
		errs = append(errs, errors.New("required ADDRESSBOOK_GRPC_PORT not set or empty"))
	}
	if _, e := strconv.Atoi(portgRPC); e == nil {
		// just checking to see if is a valid integer
		portgRPC = ":" + portgRPC
	} else {
		errs = append(errs, fmt.Errorf("required ADDRESSBOOK_GRPC_PORT is not a valid integer got: '%v'", portgRPC))
	}

	if logLevelStr, ok = os.LookupEnv("ADDRESSBOOK_SERVICE_LOG_LEVEL"); !ok || logLevelStr == "" {
		logLevel = zerolog.InfoLevel
	} else {
		var err error
		logLevel, err = zerolog.ParseLevel(logLevelStr)
		if err != nil {
			msg := "logLevel set in ADDRESSBOOK_SERVICE_LOG_LEVEL is not a valid level. "
			msg += "Must be one of: [trace, debug, info, warn, error, fatal, panic]"
			errs = append(errs, fmt.Errorf(msg))
		}
	}

	if privateKeyStr, ok = os.LookupEnv("ADDRESSBOOK_ENCRYPTION_KEY"); !ok || privateKeyStr == "" {
		errs = append(errs, errors.New("required ADDRESSBOOK_ENCRYPTION_KEY not set or empty"))
	} else {
		var err error
		privateKey, publicKey, err = parsePrivateKey(privateKeyStr)
		if err != nil {
			errs = append(errs, err)
		}
	}

	e.serviceConfig = &ServiceConfig{
		ListenPortHTTP: portHTTP,
		ListenPortgRPC: portgRPC,
		LogLevel:       logLevel,
		PrivateKey:     privateKey,
		PublicKey:      publicKey,
	}

	return errs
}

func parsePrivateKey(raw string) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(raw))
	if block == nil {
		// https://pkg.go.dev/encoding/pem
		return nil, nil, fmt.Errorf("appconfig.parsePrivateKey.pem.Decode: is format pem?")
	}

	pKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("appconfig.parsePrivateKey.x509.ParsePKCS8PrivateKey: %w", err)
	}

	privateKey, ok := pKey.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("appconfig.parsePrivateKey: could not type assert private key to *rsa.PrivateKey, got type: %T", pKey)
	}

	return privateKey, &privateKey.PublicKey, nil
}
