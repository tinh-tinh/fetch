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

func RegisterSchema[M any](opt *Config) core.Module {
	return func(module *core.DynamicModule) *core.DynamicModule {
		fetchModule := module.New(core.NewModuleOptions{})

		fetchModule.NewProvider(core.ProviderOptions{
			Name:  FETCH,
			Value: CreateSchema[M](opt),
		})
		fetchModule.Export(FETCH)

		return fetchModule
	}
}

func Inject(module *core.DynamicModule) *Fetch[any] {
	fetch, ok := module.Ref(FETCH).(*Fetch[any])
	if !ok {
		return nil
	}
	return fetch
}

func InjectSchema[M any](module *core.DynamicModule) *Fetch[M] {
	fetch, ok := module.Ref(FETCH).(*Fetch[M])
	if !ok {
		return nil
	}
	return fetch
}
