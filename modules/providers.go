package modules

import (
	"io"
	"net"
	"reflect"
)

type FilterProvider interface {
	Configuration
	Filter(targets []System) []System
}

type CommandProvider interface {
	Configuration
	Command(targets []System, resp CommandResponse) error
}

type NetworkingProvider interface {
	Configuration
	Dial() net.Conn
	Listen() net.Listener
}

type EncryptionProvider interface {
	Configuration
	NewStream(conn io.ReadWriter) io.ReadWriter
}

type DataProvider interface {
	Configuration
	Store
}

type Store interface {
	Lock(key string) error
	Unlock(key string) error
	Get(key string, result interface{}) error
	Set(key string, data interface{}) error
}

type DiscoveryProvider interface {
	Configuration
	GetNetworkingConfig() (Configuration, error)
}

func addrOf(val interface{}) interface{} {
	return &val
}

type ProviderType reflect.Type

var (
	TypeCommandProvider    ProviderType = reflect.TypeOf(new(CommandProvider)).Elem()
	TypeNetworkingProvider ProviderType = reflect.TypeOf(new(NetworkingProvider)).Elem()
	TypeEncryptionProvider ProviderType = reflect.TypeOf(new(EncryptionProvider)).Elem()
	TypeDataProvider       ProviderType = reflect.TypeOf(new(DataProvider)).Elem()
	TypeDiscoveryProvider  ProviderType = reflect.TypeOf(new(DiscoveryProvider)).Elem()
)

var AllProviders = []ProviderType{
	TypeCommandProvider,
	TypeNetworkingProvider,
	TypeEncryptionProvider,
	TypeDataProvider,
	TypeDiscoveryProvider,
}

func GetProviderType(module interface{}) ProviderType {
	inputModule := reflect.TypeOf(module)
	for _, provider := range AllProviders {
		if inputModule.AssignableTo(provider) {
			return provider
		}
	}

	return nil
}
