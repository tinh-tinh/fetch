package fetch

import "github.com/tinh-tinh/tinhtinh/v2/core"

const FETCH core.Provide = "FETCH"

// Register registers a Fetch module with the given configuration.
// The Fetch module will be available as a provider named "FETCH"
// and can be injected into other modules.
func Register(opt *Config) core.Modules {
	return func(module core.Module) core.Module {
		fetchModule := module.New(core.NewModuleOptions{})

		fetchModule.NewProvider(core.ProviderOptions{
			Name:  FETCH,
			Value: Create(opt),
		})
		fetchModule.Export(FETCH)

		return fetchModule
	}
}

// Inject returns the Fetch instance registered with the given module.
// It returns nil if the Fetch module is not registered with the given module.
func Inject(module core.Module) *Fetch {
	fetch, ok := module.Ref(FETCH).(*Fetch)
	if !ok {
		return nil
	}
	return fetch
}
