package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func instances() *ec2.DescribeInstancesOutput {
	foo := &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			{
				Instances: []*ec2.Instance{
					{
						State: &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameRunning)},
						Tags: []*ec2.Tag{
							{
								Key:   aws.String("Service"),
								Value: aws.String("whizbang"),
							},
							{
								Key:   aws.String("Team"),
								Value: aws.String("seahawks"),
							},
						},
						NetworkInterfaces: []*ec2.InstanceNetworkInterface{{PrivateDnsName: aws.String("ip-1.1.1.1.us-west-2.computer.internal")}},
					},
				},
			},
		},
	}

	return foo
}
