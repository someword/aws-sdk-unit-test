package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockEC2Client struct {
	ec2iface.EC2API
	Output *ec2.DescribeInstancesOutput
	Error  error
}

func (m *mockEC2Client) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return m.Output, m.Error
}

func TestGetEKSEC2InstanceHostname(t *testing.T) {

	instances := []*ec2.Instance{
		&ec2.Instance{
			State:      &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameRunning)},
			InstanceId: aws.String(fmt.Sprintf("i-%s", "i1")),
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
		&ec2.Instance{
			State:      &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameStopped)},
			InstanceId: aws.String(fmt.Sprintf("i-%s", "i2")),
			Tags: []*ec2.Tag{
				{
					Key:   aws.String("Role"),
					Value: aws.String("Kubernetes"),
				},
				{
					Key:   aws.String("Team"),
					Value: aws.String("seahawks"),
				},
			},
		},
		&ec2.Instance{
			State:      &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameRunning)},
			InstanceId: aws.String(fmt.Sprintf("i-%s", "i3")),
			Tags: []*ec2.Tag{
				{
					Key:   aws.String("Role"),
					Value: aws.String("Kubernetes"),
				},
				{
					Key:   aws.String("Team"),
					Value: aws.String("rams"),
				},
			},
		},
	}

	r := &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			&ec2.Reservation{
				Instances: instances,
			},
		},
	}
	tests := []struct {
		name   string
		Output *ec2.DescribeInstancesOutput
		Error  error
	}{
		{"test 1", r, nil},
		{"test 2", &ec2.DescribeInstancesOutput{}, nil},
	}

	for _, test := range tests {
		m := &mockEC2Client{Output: test.Output}
		c := EC2Client{svc: m}
		h := c.GetHost()
		fmt.Printf("test name is %s\n", test.name)
		fmt.Println(h)

	}
}
