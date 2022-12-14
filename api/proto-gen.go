// Copyright 2022 Kai Huang(seakia@live.cn).

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

//go:generate protoc -I=. --go_out=. --go_opt paths=source_relative common.proto
//go:generate protoc -I=. --go-grpc_out=. --go-grpc_opt paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=backend.yaml backend.proto
