package core

import (
	"errors"
	"fmt"
	"net/url"
)

type GamelContext interface {
	Service
	Name() string
	WithName(string) GamelContext

	Status() ContextStatus

	GetComponent(name string) (Component, error)
	GetComponentFromURI(uri string) (Component, error)

	AddRoute(route Route) GamelContext
}

type ContextStatus int

const (
	Idle ContextStatus = iota
	Started
	Error
	Stopped
	Suspended
	Resumed
)

type DefaultGamelContext struct {
	name   string
	status ContextStatus
	routes []Route
}

func NewGamelContext() GamelContext {
	return &DefaultGamelContext{
		name:   "gamel",
		status: Idle,
	}
}

func (context *DefaultGamelContext) Start() error {
	fmt.Println("Starting Gamel Context...")
	for i := 0; i < len(context.routes); i++ {
		err := context.routes[i].Start()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Started
	fmt.Println("Gamel Context Started")
	return nil
}

func (context *DefaultGamelContext) Stop() error {
	fmt.Println("Stopping Gamel Context...")
	for i := 0; i < len(context.routes); i++ {
		err := context.routes[i].Stop()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Stopped
	fmt.Println("Gamel Context Stopped")
	return nil
}

func (context *DefaultGamelContext) Suspend() error {
	fmt.Println("Suspending Gamel Context...")
	for i := 0; i < len(context.routes); i++ {
		err := context.routes[i].Suspend()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Suspended
	fmt.Println("Gamel Context Suspended")
	return nil
}

func (context *DefaultGamelContext) Resume() error {
	fmt.Println("Resuming Gamel Context...")
	for i := 0; i < len(context.routes); i++ {
		err := context.routes[i].Resume()
		if err != nil {
			context.status = Error
			return err
		}
	}
	context.status = Resumed
	fmt.Println("Gamel Context Resumed")
	return nil
}

func (context DefaultGamelContext) Name() string {
	return context.name
}

func (context *DefaultGamelContext) WithName(name string) GamelContext {
	context.name = name
	return context
}

func (context DefaultGamelContext) Status() ContextStatus {
	return context.status
}

func (context DefaultGamelContext) GetComponent(name string) (Component, error) {

	switch name {
	case "timer":
		return &TimerComponent{}, nil
	case "log":
		return &LogComponent{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot find component %s", name))
	}
}


func (context DefaultGamelContext) GetComponentFromURI(uri string) (Component, error){
	parsedUrl, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	name := parsedUrl.Scheme
	return context.GetComponent(name)
}

func (context *DefaultGamelContext) AddRoute(route Route) GamelContext {
	context.routes = append(context.routes, route)
	return context
}
