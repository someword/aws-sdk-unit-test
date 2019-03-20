package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EC2Client blah blah
type EC2Client struct {
	svc ec2iface.EC2API
}

// NewEC2Client returns and awsClient struct
func NewEC2Client(profile, region string) *EC2Client {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
		Profile:           profile,
	}))

	EC2Client := &EC2Client{
		svc: ec2.New(sess),
	}
	return EC2Client
}

// GetHost blah blah
func (e *EC2Client) GetHost() string {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("tag:Role"),
				Values: []*string{aws.String("Kubernetes")},
			},
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		},
	}
	instances, err := e.svc.DescribeInstances(params)
	fmt.Printf("instances are %#v\n", instances)

	if err != nil {
		log.Fatalf("Error is %s\n", err)
	}

	results := []*ec2.Instance{}
	for _, reservation := range instances.Reservations {
		results = append(results, reservation.Instances...)
	}
	if len(results) == 0 {
		return "No matching hosts found"
	}
	return *results[0].NetworkInterfaces[0].PrivateDnsName
}

func main() {

	n := NewEC2Client("tools-dev", "us-west-2")
	fmt.Println(n.GetHost())

}
