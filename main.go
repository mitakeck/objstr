package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

func main() {
	var (
		atoken   string
		asecret  string
		bucket   string
		endPoint string
	)

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&atoken, "t", "", "specify a access token")
	f.StringVar(&asecret, "s", "", "specify a access token secret")
	f.StringVar(&bucket, "b", "", "specify a bucket")
	f.StringVar(&endPoint, "e", "https://b.sakurastorage.jp", "specify a end point url")

	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if atoken == "" || asecret == "" || bucket == "" {
		usage := `$ objstr -t ACCESS_TOKEN -s ACCESS_TOKEN_SECRET -b BUCKET [-e END_POINT_URL]

-t ACCESS_TOKEN           access token
-s ACCESS_TOKEN_SECRET    access token secret
-b BUCKET                 specify a target bucket name
-e END_PONT_URL           specify a end point url (default: "https://b.sakurastorage.jp")`
		fmt.Println(usage)
		return
	}

	check(atoken, asecret, bucket, endPoint)
}

func check(atoken string, asecret string, bucket string, endPoint string) {
	auth, err := aws.GetAuth(atoken, asecret)

	if err != nil {
		panic(err.Error())
	}

	client := s3.New(auth, aws.Region{
		Name:       bucket,
		S3Endpoint: endPoint,
	})

	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.ListBuckets()

	if err != nil {
		log.Fatal(err)
	}

	log.Print(fmt.Sprintf("%T %+v", resp.Buckets[0], resp.Buckets[0]))
}
