package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kdar/logrus-cloudwatchlogs"
	log "github.com/sirupsen/logrus"
)

func main() {

	r := registerRoutes()

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	group := os.Getenv("AWS_CLOUDWATCHLOGS_GROUP_NAME")
	stream := os.Getenv("AWS_CLOUDWATCHLOGS_STREAM_NAME")

	sess := session.Must(session.NewSession())

	hook, err := logrus_cloudwatchlogs.NewHook(group, stream, sess.Config)
	if err != nil {
		log.Fatal(err)
	}

	log.AddHook(hook)
	log.SetFormatter(logrus_cloudwatchlogs.NewProdFormatter())

	/* Log format:
	{
		"app": "application",
		"host": "ip-172-31-59-246",
		"level": "error",
		"msg": "Signin Error: 2 validation error(s) found.",
		"time": 1512147448
	}
	*/

	log.Info("Listening on :", port)
	r.Run(":" + port)
}
