package main

import (
	"context"
	"flag"
	"fmt"
	booksapi "library"
	auther "library/gen/auther"
	books "library/gen/books"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")

		dbhost     = flag.String("dbhost", "localhost", "Database host")
		dbport     = flag.Int("dbport", 5432, "Database port")
		dbuser     = flag.String("dbuser", "postgres", "Database username")
		dbpassword = flag.String("dbpassword", "postgres", "Database password")
		dbname     = flag.String("dbname", "postgres", "Database name")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[booksapi] ", log.Ltime)
	}

	// Setup db connection
	var (
		booksRepo booksapi.BooksRepo
	)
	{
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			*dbhost, *dbport, *dbuser, *dbpassword, *dbname)
		db, err := sqlx.Open("postgres", psqlInfo)

		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open database %v: %s", psqlInfo, err)
			os.Exit(1)
		}

		if err := db.Ping(); err != nil {
			fmt.Fprintf(os.Stderr, "could not ping database %v: %s", psqlInfo, err)
			os.Exit(1)
		}

		booksRepo = booksapi.NewDbBooksRepo(db)
	}

	// Initialize the services.
	var (
		booksSvc  books.Service
		autherSvc auther.Service
	)
	{
		booksSvc = booksapi.NewBooks(logger, booksRepo)
		autherSvc = booksapi.NewAuther(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		booksEndpoints  *books.Endpoints
		autherEndpoints *auther.Endpoints
	)
	{
		booksEndpoints = books.NewEndpoints(booksSvc)
		autherEndpoints = auther.NewEndpoints(autherSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8088"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, booksEndpoints, autherEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
