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

package backend

import (
	"context"
	"strconv"
	"time"

	"github.com/recsys-infra/simple-recsys-in-k8s/api"

	"github.com/golang/glog"
)

// RecSys 推荐入口
func (s *service) RecSys(parentCtx context.Context, req *api.Request) (rsp *api.Response, err error) {
	glog.Infof("req: %s", req.String())

	return &api.Response{
		RequestId: req.GetRequestId(),
		UserId:    strconv.Itoa(int(time.Now().Unix())),
	}, nil
}
