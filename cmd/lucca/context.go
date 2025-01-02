package main

import (
	"context"

	"github.com/nlm/lucca-api/api"
	"github.com/nlm/lucca-api/cmd/lucca/config"
)

type ContextKey string

var (
	ctxKeyConfig ContextKey = "config"
	ctxKeyClient ContextKey = "client"
)

type Context struct {
	context.Context
}

func NewContext(ctx context.Context) Context {
	return Context{
		Context: ctx,
	}
}

func (ctx Context) WithClient(client *api.Client) Context {
	return Context{
		Context: context.WithValue(ctx.Context, ctxKeyClient, client),
	}
}

func (ctx Context) WithConfig(config *config.Config) Context {
	return Context{
		Context: context.WithValue(ctx.Context, ctxKeyConfig, config),
	}
}

func (ctx Context) Client() *api.Client {
	return ctx.Context.Value(ctxKeyClient).(*api.Client)
}

func (ctx Context) Config() *config.Config {
	return ctx.Context.Value(ctxKeyConfig).(*config.Config)
}
