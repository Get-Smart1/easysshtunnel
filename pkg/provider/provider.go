package provider

import "easytunnel/pkg/middelware"

type Provider interface {
	Initialize(middleware *middelware.Middleware)
}
