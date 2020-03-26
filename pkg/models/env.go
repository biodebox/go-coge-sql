package models

type (
	env struct {
		transformer Transformer
		hydrator    Hydrator
	}
)

func (e *env) GetTransformer() Transformer {
	panic("implement me")
}

func (e *env) GetHydrator() Hydrator {
	panic("implement me")
}

func EnvConfiguratorFactory(options string, params map[string]interface{}) (Configurator, error) {
	return &env{}, nil
}