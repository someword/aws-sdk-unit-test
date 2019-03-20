package main

import (
	"fmt"
	"testing"

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
	tests := []struct {
		name   string
		Output *ec2.DescribeInstancesOutput
		Error  error
	}{
		{"test 1", &ec2.DescribeInstancesOutput{}, nil},
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
