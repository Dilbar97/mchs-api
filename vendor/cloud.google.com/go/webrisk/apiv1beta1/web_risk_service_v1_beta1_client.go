// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package webrisk

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	webriskpb "google.golang.org/genproto/googleapis/cloud/webrisk/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newWebRiskServiceV1Beta1ClientHook clientHook

// WebRiskServiceV1Beta1CallOptions contains the retry settings for each method of WebRiskServiceV1Beta1Client.
type WebRiskServiceV1Beta1CallOptions struct {
	ComputeThreatListDiff []gax.CallOption
	SearchUris            []gax.CallOption
	SearchHashes          []gax.CallOption
}

func defaultWebRiskServiceV1Beta1ClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("webrisk.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("webrisk.mtls.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultWebRiskServiceV1Beta1CallOptions() *WebRiskServiceV1Beta1CallOptions {
	return &WebRiskServiceV1Beta1CallOptions{
		ComputeThreatListDiff: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchUris: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchHashes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// WebRiskServiceV1Beta1Client is a client for interacting with Web Risk API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type WebRiskServiceV1Beta1Client struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	webRiskServiceV1Beta1Client webriskpb.WebRiskServiceV1Beta1Client

	// The call options for this service.
	CallOptions *WebRiskServiceV1Beta1CallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewWebRiskServiceV1Beta1Client creates a new web risk service v1 beta1 client.
//
// Web Risk v1beta1 API defines an interface to detect malicious URLs on your
// website and in client applications.
func NewWebRiskServiceV1Beta1Client(ctx context.Context, opts ...option.ClientOption) (*WebRiskServiceV1Beta1Client, error) {
	clientOpts := defaultWebRiskServiceV1Beta1ClientOptions()

	if newWebRiskServiceV1Beta1ClientHook != nil {
		hookOpts, err := newWebRiskServiceV1Beta1ClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	c := &WebRiskServiceV1Beta1Client{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultWebRiskServiceV1Beta1CallOptions(),

		webRiskServiceV1Beta1Client: webriskpb.NewWebRiskServiceV1Beta1Client(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *WebRiskServiceV1Beta1Client) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *WebRiskServiceV1Beta1Client) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *WebRiskServiceV1Beta1Client) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ComputeThreatListDiff gets the most recent threat list diffs.
func (c *WebRiskServiceV1Beta1Client) ComputeThreatListDiff(ctx context.Context, req *webriskpb.ComputeThreatListDiffRequest, opts ...gax.CallOption) (*webriskpb.ComputeThreatListDiffResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.ComputeThreatListDiff[0:len(c.CallOptions.ComputeThreatListDiff):len(c.CallOptions.ComputeThreatListDiff)], opts...)
	var resp *webriskpb.ComputeThreatListDiffResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.ComputeThreatListDiff(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchUris this method is used to check whether a URI is on a given threatList.
func (c *WebRiskServiceV1Beta1Client) SearchUris(ctx context.Context, req *webriskpb.SearchUrisRequest, opts ...gax.CallOption) (*webriskpb.SearchUrisResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.SearchUris[0:len(c.CallOptions.SearchUris):len(c.CallOptions.SearchUris)], opts...)
	var resp *webriskpb.SearchUrisResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.SearchUris(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchHashes gets the full hashes that match the requested hash prefix.
// This is used after a hash prefix is looked up in a threatList
// and there is a match. The client side threatList only holds partial hashes
// so the client must query this method to determine if there is a full
// hash match of a threat.
func (c *WebRiskServiceV1Beta1Client) SearchHashes(ctx context.Context, req *webriskpb.SearchHashesRequest, opts ...gax.CallOption) (*webriskpb.SearchHashesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.SearchHashes[0:len(c.CallOptions.SearchHashes):len(c.CallOptions.SearchHashes)], opts...)
	var resp *webriskpb.SearchHashesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.webRiskServiceV1Beta1Client.SearchHashes(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}