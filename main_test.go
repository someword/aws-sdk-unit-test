package main

import (
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
	tests := []struct {
		name, expected string
		Output         *ec2.DescribeInstancesOutput
		Error          error
	}{
		{"test 1", "ip-1.1.1.1.us-west-2.computer.internal", &ec2.DescribeInstancesOutput{
			Reservations: []*ec2.Reservation{
				{
					Instances: []*ec2.Instance{
						{
							State: &ec2.InstanceState{Name: aws.String(ec2.InstanceStateNameStopped)},
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
							NetworkInterfaces: []*ec2.InstanceNetworkInterface{{PrivateDnsName: aws.String("ip-2.2.2.2.us-west-2.computer.internal")}},
						},
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
		}, nil},
		{"test 2", "blaj", &ec2.DescribeInstancesOutput{}, nil},
	}

	for _, test := range tests {
		m := &mockEC2Client{Output: test.Output}
		c := EC2Client{svc: m}
		actual := c.GetHost()
		if actual != test.expected {
			t.Errorf("GetHost(): test name: %s expected: %s actual: %s\n", test.name, test.expected, actual)
		}
	}
}
