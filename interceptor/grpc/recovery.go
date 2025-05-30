package grpcinterceptor

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/getsentry/sentry-go"
	"google.golang.org/grpc"

	"github.com/ggsrc/blueprint-go-toolkits/zerolog/log"
)

func SentryUnaryServerInterceptor(sentryDSN string, sentryEnv string) grpc.UnaryServerInterceptor {
	err := sentry.Init(sentry.ClientOptions{Dsn: sentryDSN, Environment: sentryEnv})
	if err != nil {
		log.Err(err).Msg("sentry init failed, ignore it and continue...")
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Ctx(ctx).Error().
					Str("panic.stack", string(debug.Stack())).
					Err(fmt.Errorf("[panic] %v", r)).
					Msgf("%s grpc server panic", strings.Trim(info.FullMethod, "/"))
				err = fmt.Errorf("server Internal Error")
				hub := sentry.CurrentHub()
				hub.Recover(r)
			}
		}()

		resp, err = handler(ctx, req)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("grpc server error")
		}
		return resp, err
	}
}

func SentryUnaryClientInterceptor(sentryDSN string, sentryEnv string) grpc.UnaryClientInterceptor {
	err := sentry.Init(sentry.ClientOptions{Dsn: sentryDSN, Environment: sentryEnv})
	if err != nil {
		log.Err(err).Msg("sentry init failed")
	}

	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Ctx(ctx).Error().
					Str("panic.stack", string(debug.Stack())).
					Err(fmt.Errorf("[panic] %v", r)).
					Msgf("%s grpc client panic", strings.Trim(method, "/"))
				err = fmt.Errorf("server Internal Error")
			}
		}()

		err = invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("grpc client error")
		}
		return err
	}
}
