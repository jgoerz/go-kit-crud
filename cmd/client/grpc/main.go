package main

// The reference for this example
// https://github.com/PacktPublishing/Go-Programming-Blueprints/blob/master/Chapter10/vault/cmd/vaultcli/main.go

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	// FIXME clients should not be referencing internal implementation
	"github.com/jgoerz/go-kit-crud/internal/addressbook"
	client "github.com/jgoerz/go-kit-crud/pkg/client/addressbook"
	grpcclient "github.com/jgoerz/go-kit-crud/pkg/client/grpc"

	"google.golang.org/grpc"
)

func main() {
	var grpcAddr string

	flag.StringVar(&grpcAddr, "addr", LookupEnvOrString("ADDRESSBOOK_GRPC_PORT", "3334"), "gRPC address")
	flag.Parse()

	ctx := context.Background()
	conn, err := grpc.Dial(":"+grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	service := grpcclient.New(conn)

	args := flag.Args()
	var cmd string
	cmd, args = pop(args)

	switch cmd {
	case "create-contact":
		var rawTenantID, firstName, lastName, rawActive, address, someSecret string

		rawTenantID, args = pop(args)
		firstName, args = pop(args)
		lastName, args = pop(args)
		rawActive, args = pop(args)
		address, args = pop(args)
		someSecret, _ = pop(args)
		createContact(ctx, service, rawTenantID, firstName, lastName, rawActive, address, someSecret)

	case "read-contact":
		var rawID string
		rawID, _ = pop(args)

		readContact(ctx, service, rawID)

	case "list-contacts":
		var (
			rawPageToken string
			rawPageSize  string
		)
		rawPageToken, _ = pop(args)
		rawPageSize, _ = pop(args)

		listContacts(ctx, service, rawPageToken, rawPageSize)

	case "update-contact":
		var rawID, rawTenantID, firstName, lastName, rawActive, address, someSecret string

		rawID, args = pop(args)
		rawTenantID, args = pop(args)
		firstName, args = pop(args)
		lastName, args = pop(args)
		rawActive, args = pop(args)
		address, args = pop(args)
		someSecret, _ = pop(args)
		updateContact(ctx, service, rawID, rawTenantID, firstName, lastName, rawActive, address, someSecret)

	case "delete-contact":
		var rawID string
		rawID, _ = pop(args)

		deleteContact(ctx, service, rawID)

	default:
		log.Fatalln("unknown/unsupported command", cmd)
	}

}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// https://www.gmarik.info/blog/2019/12-factor-golang-flag-package/
func LookupEnvOrString(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func createContact(ctx context.Context, service addressbook.Service,
	rawTenantID, firstName, lastName, rawActive, address, someSecret string) {

	tenantID, err := strconv.ParseInt(rawTenantID, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	active, err := strconv.ParseBool(rawActive)
	if err != nil {
		log.Fatalln(err.Error())
	}

	input := &client.ContactRequest{
		TenantID:   tenantID,
		FirstName:  firstName,
		LastName:   lastName,
		Active:     active,
		Address:    address,
		SomeSecret: someSecret,
	}

	contact, err := service.CreateContact(ctx, input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(contact)
}

func readContact(ctx context.Context, service addressbook.Service, rawID string) {
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	input := &client.ReadContactRequest{
		ID: id,
	}

	contact, err := service.ReadContact(ctx, input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(contact)
}

func listContacts(ctx context.Context, service addressbook.Service, rawPageToken, rawPageSize string) {
	pageToken, err := strconv.ParseInt(rawPageToken, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	pageSize, err := strconv.ParseInt(rawPageSize, 10, 32)
	if err != nil {
		log.Fatalln(err.Error())
	}

	input := &client.ListContactsRequest{
		PageToken: pageToken,
		PageSize:  int32(pageSize),
	}

	list, err := service.ListContacts(ctx, input)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, contact := range list.ContactResponses {
		fmt.Println(contact)
	}
	fmt.Println("NextPageToken: ", list.NextPageToken)
}

func updateContact(ctx context.Context, service addressbook.Service,
	rawID, rawTenantID, firstName, lastName, rawActive, address, someSecret string) {

	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	tenantID, err := strconv.ParseInt(rawTenantID, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	active, err := strconv.ParseBool(rawActive)
	if err != nil {
		log.Fatalln(err.Error())
	}

	input := &client.ContactRequest{
		ID:         id,
		TenantID:   tenantID,
		FirstName:  firstName,
		LastName:   lastName,
		Active:     active,
		Address:    address,
		SomeSecret: someSecret,
	}

	contact, err := service.UpdateContact(ctx, input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(contact)
}

func deleteContact(ctx context.Context, service addressbook.Service, rawID string) {

	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}

	input := &client.DeleteContactRequest{
		ID: id,
	}

	contact, err := service.DeleteContact(ctx, input)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(contact)
}
