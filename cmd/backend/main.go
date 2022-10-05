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

package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/recsys-infra/simple-recsys-in-k8s/api"
	"github.com/recsys-infra/simple-recsys-in-k8s/web/backend"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strings"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":8080", "gRPC server endpoint")
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func newOpts() []grpc.DialOption {
	var op []grpc.DialOption
	op = append(op,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return op
}

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcServer := grpc.NewServer()
	api.RegisterBackendServiceServer(grpcServer, backend.NewService())

	gwMux := runtime.NewServeMux()
	if err := api.RegisterBackendServiceHandlerFromEndpoint(ctx, gwMux, *grpcServerEndpoint,
		newOpts(),
	); err != nil {
		glog.Errorf("failed to register handler")
		return
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", gwMux)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		glog.Errorf("failed to listen: %+v", err)
		return
	}

	srv := &http.Server{
		Addr:      *grpcServerEndpoint,
		Handler:   grpcHandlerFunc(grpcServer, httpMux),
		TLSConfig: nil,
	}

	if err := srv.Serve(l); err != nil {
		glog.Errorf("failed to serve: %+v", err)
		return
	}

	glog.Infof("server stopped!!!")
}
