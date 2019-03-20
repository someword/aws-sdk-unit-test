package main

import (
	"errors"
	"reflect"
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
		name     string
		expected []string
		Output   *ec2.DescribeInstancesOutput
		Error    error
	}{
		{"test 1", []string{"ip-2.2.2.2.us-west-2.computer.internal", "ip-1.1.1.1.us-west-2.computer.internal"}, &ec2.DescribeInstancesOutput{
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
									Key:   aws.String("Role"),
									Value: aws.String("Kubernetes"),
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
		},
			nil},
		{"test 2", []string{}, &ec2.DescribeInstancesOutput{}, errors.New("No matching hosts found")},
	}

	for _, test := range tests {
		m := &mockEC2Client{Output: test.Output}
		c := EC2Client{svc: m}
		actual, err := c.GetHost()
		if !reflect.DeepEqual(actual, test.expected) || !reflect.DeepEqual(err, test.Error) {
			t.Errorf("GetHost(): test name: %s expected: %s actual: %s \n expected error: %s error: %s\n", test.name, test.expected, actual, test.Error, err)
		}
	}
}
