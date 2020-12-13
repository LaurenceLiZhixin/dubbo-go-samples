/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkg

import (
	"context"
	"github.com/apache/dubbo-go-samples/general/dubbo3/protobuf"
	"github.com/apache/dubbo-go/remoting/dubbo3"
	_ "github.com/apache/dubbo-go/protocol/dubbo3/impl"
)


type GrpcGreeterImpl struct {
	SayHello func(ctx context.Context, in *protobuf.HelloRequest, out *protobuf.HelloReply) error
}

func (u *GrpcGreeterImpl) Reference() string {
	return "GrpcGreeterImpl"
}

func (u *GrpcGreeterImpl) GetDubboStub(cc *dubbo3.TripleConn) protobuf.GreeterClient {
	return protobuf.NewGreeterDubbo3Client(cc)
}