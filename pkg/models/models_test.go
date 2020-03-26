package models_test

//func TestGenerator_GetConfigurator(t *testing.T) {
//	type fields struct {
//		Configurators map[string]models.ConfiguratorFactory
//	}
//	type args struct {
//		config string
//		params map[string]interface{}
//	}
//	configuratorMock := &mocks.Configurator{}
//	factory := func(options string, params map[string]interface{}) (models.Configurator, error) {
//		return configuratorMock, nil
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    models.Configurator
//		wantErr bool
//	}{
//		{
//			name:    `empty config`,
//			args:    args{config: ``, params: map[string]interface{}{}},
//			want:    nil,
//			wantErr: true,
//		},
//		{
//			name:    `config only dotEnv`,
//			fields: fields{Configurators: map[string]models.ConfiguratorFactory{`dotEnv`:factory}},
//			args:    args{config: `dotEnv`, params: map[string]interface{}{}},
//			want:    configuratorMock,
//			wantErr: false,
//		},
//		{
//			name:    `empty config dotEnv:.env`,
//			fields: fields{Configurators: map[string]models.ConfiguratorFactory{`dotEnv`:factory}},
//			args:    args{config: `dotEnv:.env`, params: map[string]interface{}{}},
//			want:    configuratorMock,
//			wantErr: false,
//		},
//		{
//			name:    `empty config dotEnv : .env`,
//			fields: fields{Configurators: map[string]models.ConfiguratorFactory{`dotEnv`:factory}},
//			args:    args{config: `dotEnv : .env`, params: map[string]interface{}{}},
//			want:    configuratorMock,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			g := &models.Generator{
//				Configurators: tt.fields.Configurators,
//			}
//			got, err := g.GetConfigurator(tt.args.config, tt.args.params)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetConfigurator() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetConfigurator() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGenerator_GetConfigurator1(t *testing.T) {
//	type fields struct {
//		Configurators map[string]ConfiguratorFactory
//	}
//	type args struct {
//		name   string
//		option string
//		params map[string]interface{}
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    models.Configurator
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			g := &models.Generator{
//				Configurators: tt.fields.Configurators,
//			}
//			got, err := g.GetConfigurator(tt.args.name, tt.args.option, tt.args.params)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetConfigurator() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetConfigurator() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}