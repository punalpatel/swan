package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dataman-Cloud/swan/src/util"
	"github.com/boltdb/bolt"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func setupLogger(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)

	logrus.SetOutput(os.Stderr)

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

// waitForSignals wait for signals and do some clean up job.
func waitForSignals(unixSock string) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for sig := range signals {
		logrus.Debugf("Received signal %s , clean up...", sig)
		if _, err := os.Stat(unixSock); err == nil {
			logrus.Debugf("Remove %s", unixSock)
			if err := os.Remove(unixSock); err != nil {
				logrus.Errorf("Remove %s failed: %s", unixSock, err.Error())
			}
		}

		os.Exit(0)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "swan"
	app.Usage = "A general purpose mesos framework"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "0.0.0.0:9999",
			Usage: "API Server address <ip:port>",
		},
		cli.StringFlag{
			Name:  "sock",
			Value: "/var/run/swan.sock",
			Usage: "Unix socket for listening",
		},
		cli.StringFlag{
			Name:  "masters",
			Value: "127.0.0.1:5050",
			Usage: "masters address <ip:port>,<ip:port>...",
		},
		cli.StringFlag{
			Name:  "user",
			Value: "root",
			Usage: "mesos framework user",
		},
		cli.StringFlag{
			Name:  "log-level",
			Value: "info",
			Usage: "customize debug level [debug|info|error]",
		},
		cli.IntFlag{
			Name:  "raftid",
			Value: 1,
			Usage: "raft node id",
		},
		cli.StringFlag{
			Name:  "cluster",
			Value: "http://127.0.0.1.2221",
			Usage: "raft cluster peers addr",
		},
		cli.BoolFlag{
			Name:  "enable-dns-proxy",
			Usage: "enable dns proxy or not",
		},
		cli.BoolFlag{
			Name:  "enable-local-healthcheck",
			Usage: "Enable local healt check",
		},
		cli.BoolFlag{
			Name:  "standalone",
			Usage: "Run as standalone mode",
		},
		cli.StringFlag{
			Name:  "mode",
			Value: "mixed",
			Usage: "Server mode, manager|agent|mixed ",
		},
		cli.StringFlag{
			Name:  "work-dir",
			Value: "./data/",
			Usage: "swan data store dir",
		},
	}

	app.Action = func(c *cli.Context) error {
		config, err := util.NewConfig(c)
		if err != nil {
			os.Exit(1)
		}

		setupLogger(config.LogLevel)

		db, err := bolt.Open(fmt.Sprintf(c.String("work-dir")+".bolt.db.%d", config.Raft.RaftId), 0600, nil)
		if err != nil {
			logrus.Errorf("Init store engine failed:%s", err)
			return err
		}

		node, _ := NewNode(config, db)
		go func() {
			node.Start(context.Background())
		}()

		waitForSignals(config.HttpListener.UnixAddr)

		return nil
	}

	app.Run(os.Args)
}
