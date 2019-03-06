package main

import (
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"

	log "github.com/sirupsen/logrus"

	"github.com/itmecho/rabbithole/internal/spammer"

	"github.com/itmecho/rabbithole/pkg/rabbitmqclient"
)

var (
	version    string
	commitHash string

	flagHost    = flag.String("host", "amqp://guest:guest@localhost:5672/", "Full URL of the RabbitMQ server")
	flagQueue   = flag.String("queue", "rabbithole", "Name of the queue to send messages to")
	flagCount   = flag.Int("count", 100, "Number of messages to send to the queue")
	flagMinSize = flag.Int("min-size", 100, "Minimum size of messages in bytes")
	flagMaxSize = flag.Int("max-size", 10000, "Maximum size of messages in bytes")

	flagVerbose = flag.Bool("verbose", false, "Enable verbose logging")
	flagVersion = flag.Bool("version", false, "Show version information")
)

func init() {
	flag.Parse()

	if *flagVersion {
		fmt.Printf("rabbithole version %s+%s\n", version, commitHash)
		os.Exit(0)
	}

	if *flagVerbose {
		log.Info("Enabling verbose logging")
		log.SetLevel(log.DebugLevel)
	}

	if *flagMinSize > *flagMaxSize {
		log.Debugf("min size is more than max size - setting max size to min size")
		*flagMaxSize = *flagMinSize
	}

}

func main() {
	client, err := rabbitmqclient.NewClient(*flagHost, *flagQueue)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	startTime := time.Now()

	err = spammer.Send(client, spammer.SendConfig{
		Count:   *flagCount,
		MinSize: *flagMinSize,
		MaxSize: *flagMaxSize,
	})
	if err != nil {
		log.Fatal(err)
	}

	endTime := time.Since(startTime)

	fmt.Println("Done!")
	fmt.Println(*flagCount, "messages sent in", endTime.Seconds(), "seconds")
}
