// Code generated by goa v3.0.3, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen library/design

package cli

import (
	"flag"
	"fmt"
	autherc "library/gen/http/auther/client"
	booksc "library/gen/http/books/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `auther signin
books (list|reserve|pickup|return|subscribe)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` auther signin --username "user" --password "password"` + "\n" +
		os.Args[0] + ` books list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		autherFlags = flag.NewFlagSet("auther", flag.ContinueOnError)

		autherSigninFlags        = flag.NewFlagSet("signin", flag.ExitOnError)
		autherSigninUsernameFlag = autherSigninFlags.String("username", "REQUIRED", "Username used to perform signin")
		autherSigninPasswordFlag = autherSigninFlags.String("password", "REQUIRED", "Password used to perform signin")

		booksFlags = flag.NewFlagSet("books", flag.ContinueOnError)

		booksListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		booksReserveFlags      = flag.NewFlagSet("reserve", flag.ExitOnError)
		booksReserveBodyFlag   = booksReserveFlags.String("body", "REQUIRED", "")
		booksReserveBookIDFlag = booksReserveFlags.String("bookid", "REQUIRED", "id of the Book")
		booksReserveTokenFlag  = booksReserveFlags.String("token", "REQUIRED", "")

		booksPickupFlags      = flag.NewFlagSet("pickup", flag.ExitOnError)
		booksPickupBodyFlag   = booksPickupFlags.String("body", "REQUIRED", "")
		booksPickupBookIDFlag = booksPickupFlags.String("bookid", "REQUIRED", "id of the Book")
		booksPickupTokenFlag  = booksPickupFlags.String("token", "REQUIRED", "")

		booksReturnFlags     = flag.NewFlagSet("return", flag.ExitOnError)
		booksReturnBodyFlag  = booksReturnFlags.String("body", "REQUIRED", "")
		booksReturnTokenFlag = booksReturnFlags.String("token", "REQUIRED", "")

		booksSubscribeFlags      = flag.NewFlagSet("subscribe", flag.ExitOnError)
		booksSubscribeBookIDFlag = booksSubscribeFlags.String("bookid", "REQUIRED", "id of the Book")
		booksSubscribeTokenFlag  = booksSubscribeFlags.String("token", "REQUIRED", "")
	)
	autherFlags.Usage = autherUsage
	autherSigninFlags.Usage = autherSigninUsage

	booksFlags.Usage = booksUsage
	booksListFlags.Usage = booksListUsage
	booksReserveFlags.Usage = booksReserveUsage
	booksPickupFlags.Usage = booksPickupUsage
	booksReturnFlags.Usage = booksReturnUsage
	booksSubscribeFlags.Usage = booksSubscribeUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "auther":
			svcf = autherFlags
		case "books":
			svcf = booksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "auther":
			switch epn {
			case "signin":
				epf = autherSigninFlags

			}

		case "books":
			switch epn {
			case "list":
				epf = booksListFlags

			case "reserve":
				epf = booksReserveFlags

			case "pickup":
				epf = booksPickupFlags

			case "return":
				epf = booksReturnFlags

			case "subscribe":
				epf = booksSubscribeFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "auther":
			c := autherc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "signin":
				endpoint = c.Signin()
				data, err = autherc.BuildSigninPayload(*autherSigninUsernameFlag, *autherSigninPasswordFlag)
			}
		case "books":
			c := booksc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "reserve":
				endpoint = c.Reserve()
				data, err = booksc.BuildReservePayload(*booksReserveBodyFlag, *booksReserveBookIDFlag, *booksReserveTokenFlag)
			case "pickup":
				endpoint = c.Pickup()
				data, err = booksc.BuildPickupPayload(*booksPickupBodyFlag, *booksPickupBookIDFlag, *booksPickupTokenFlag)
			case "return":
				endpoint = c.Return()
				data, err = booksc.BuildReturnPayload(*booksReturnBodyFlag, *booksReturnTokenFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = booksc.BuildSubscribePayload(*booksSubscribeBookIDFlag, *booksSubscribeTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// autherUsage displays the usage of the auther command and its subcommands.
func autherUsage() {
	fmt.Fprintf(os.Stderr, `The auther service serves authentication methods
Usage:
    %s [globalflags] auther COMMAND [flags]

COMMAND:
    signin: Creates a valid JWT

Additional help:
    %s auther COMMAND --help
`, os.Args[0], os.Args[0])
}
func autherSigninUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] auther signin -username STRING -password STRING

Creates a valid JWT
    -username STRING: Username used to perform signin
    -password STRING: Password used to perform signin

Example:
    `+os.Args[0]+` auther signin --username "user" --password "password"
`, os.Args[0])
}

// booksUsage displays the usage of the books command and its subcommands.
func booksUsage() {
	fmt.Fprintf(os.Stderr, `The books service serves operations on books: list, reserve, pickedUp, returned, subscribe
Usage:
    %s [globalflags] books COMMAND [flags]

COMMAND:
    list: List books
    reserve: Mark book as reserved. Once a book is reserved timer starts with timeout for the book to become picked up. Timeout is configurable. 
			Once timeout is expired book becomes available
    pickup: Mark book as picked up
    return: Mark book as returned
    subscribe: Subscribe the caller on the next 'book's become available

Additional help:
    %s books COMMAND --help
`, os.Args[0], os.Args[0])
}
func booksListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] books list

List books

Example:
    `+os.Args[0]+` books list
`, os.Args[0])
}

func booksReserveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] books reserve -body JSON -bookid INT64 -token STRING

Mark book as reserved. Once a book is reserved timer starts with timeout for the book to become picked up. Timeout is configurable. 
			Once timeout is expired book becomes available
    -body JSON: 
    -bookid INT64: id of the Book
    -token STRING: 

Example:
    `+os.Args[0]+` books reserve --body '{
      "subscriber_id": "Molestias est."
   }' --bookid 3733984170189243710 --token "Occaecati sint molestias assumenda sed placeat."
`, os.Args[0])
}

func booksPickupUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] books pickup -body JSON -bookid INT64 -token STRING

Mark book as picked up
    -body JSON: 
    -bookid INT64: id of the Book
    -token STRING: 

Example:
    `+os.Args[0]+` books pickup --body '{
      "subscriber_id": "Et unde eveniet autem aliquam laboriosam."
   }' --bookid 6061653207355838457 --token "Repudiandae nostrum et amet."
`, os.Args[0])
}

func booksReturnUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] books return -body JSON -token STRING

Mark book as returned
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` books return --body '{
      "book_id": 5640400661650167868,
      "subscriber_id": "Dolorum labore quas quidem sit fuga rerum."
   }' --token "Vel qui totam incidunt aut possimus mollitia."
`, os.Args[0])
}

func booksSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] books subscribe -bookid INT64 -token STRING

Subscribe the caller on the next 'book's become available
    -bookid INT64: id of the Book
    -token STRING: 

Example:
    `+os.Args[0]+` books subscribe --bookid 2021374031697615583 --token "Debitis exercitationem."
`, os.Args[0])
}
