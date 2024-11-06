package fetch

import "github.com/tinh-tinh/tinhtinh/core"

const FETCH core.Provide = "FETCH"

func Register(opt *Config) core.Module {
	return func(module *core.DynamicModule) *core.DynamicModule {
		fetchModule := module.New(core.NewModuleOptions{})

		fetchModule.NewProvider(core.ProviderOptions{
			Name:  FETCH,
			Value: Create(opt),
		})
		fetchModule.Export(FETCH)

		return fetchModule
	}
}

func Inject(module *core.DynamicModule) *Fetch {
	fetch, ok := module.Ref(FETCH).(*Fetch)
	if !ok {
		return nil
	}
	return fetch
}
