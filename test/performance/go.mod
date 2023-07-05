// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go/test/performance

replace github.com/hyperledger/fabric-sdk-go => ../../

require (
	github.com/golang/protobuf v1.5.2
	github.com/hyperledger/fabric-protos-go v0.0.0-20211118165945-23d738fc3553
	github.com/hyperledger/fabric-sdk-go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.1
	golang.org/x/net v0.5.0
	google.golang.org/grpc v1.53.0
)

go 1.14
