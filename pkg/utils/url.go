package pkgUtils

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/grpc/metadata"
)

func ParamsFromGrpc(ctx context.Context) (url.Values, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("params is missing")
	}

	rawQueries := md.Get("x-raw-query")
	if len(rawQueries) == 0 {
		return nil, fmt.Errorf("no query parameters found")
	}
	return url.ParseQuery(rawQueries[0])
}
