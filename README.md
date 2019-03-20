# Examples


When executing the command I get back the hostname of an instance which matches the filters passed in on lines 35-46 in main.go
`❯ go run main.go
ip-10-52-27-176.us-west-2.compute.internal
`



When running the unit test the returned ec2 instance should not have been returned because it should have been filtered out becauess it's state is stopped and it does not have the right tags (eg: Role = Kubernetes)

`❯ go test -v ./...
=== RUN   TestGetEKSEC2InstanceHostname
--- FAIL: TestGetEKSEC2InstanceHostname (0.00s)
    main_test.go:71: GetHost(): test name: test 1 expected: ip-1.1.1.1.us-west-2.computer.internal actual: ip-2.2.2.2.us-west-2.computer.internal
        main_test.go:71: GetHost(): test name: test 2 expected: blaj actual: No matching hosts found
        FAIL
        FAIL    test    0.027s
`

