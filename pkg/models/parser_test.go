package models

import (
	"reflect"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type fields struct {
		PathSep      string
		ListSep      string
		NameSep      string
		PrefixExclud string
		AllSymb      string
	}
	type args struct {
		expression string
	}
	fs := fields{
		PathSep:      `.`,
		ListSep:      `|`,
		NameSep:      `:`,
		PrefixExclud: `!`,
		AllSymb:      `*`,
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: `only table`,
			fields: fs,
			args:args{expression:`table`},
			want: map[string]interface{}{
				`tables`: []string{`table`},
			},
		},
		{
			name: `only table with all column`,
			fields: fs,
			args:args{expression:`table(*)`},
			want: map[string]interface{}{
				`tables`: []string{`table`},
			},
		},
		{
			name: `only table with all column and one alias`,
			fields: fs,
			args:args{expression:`table(column_id:ID|*)`},
			want: map[string]interface{}{
				`tables`: []string{`table`},
				`alias_columns`: map[string]string{`column_id`:`ID`},
			},
		},
		{
			name: `only table with all column and one alias and one exclude`,
			fields: fs,
			args:args{expression:`table(column_id:ID|!column_version|*)`},
			want: map[string]interface{}{
				`tables`: []string{`table`},
				`excluded_columns`: []string{`column_version`},
				`alias_columns`: map[string]string{`column_id`:`ID`},
			},
		},
		{
			name: `only schema and table with all column and one alias and one exclude`,
			fields: fs,
			args:args{expression:`public.table(column_id:ID|!column_version|*)`},
			want: map[string]interface{}{
				`schemas`: []string{`public`},
				`tables`: []string{`table`},
				`excluded_columns`: []string{`column_version`},
				`alias_columns`: map[string]string{`column_id`:`ID`},
			},
		},
		{
			name: `database, schema and table with all column and one alias and one exclude`,
			fields: fs,
			args:args{expression:`database.public.table(column_id:ID|!column_version|*)`},
			want: map[string]interface{}{
				`databases`: []string{`database`},
				`schemas`: []string{`public`},
				`tables`: []string{`table`},
				`excluded_columns`: []string{`column_version`},
				`alias_columns`: map[string]string{`column_id`:`ID`},
			},
		},
		{
			name: `only database`,
			fields: fs,
			args:args{expression:`database.*.*`},
			want: map[string]interface{}{
				`databases`: []string{`database`},
			},
		},
		{
			name: `only database and schema`,
			fields: fs,
			args:args{expression:`database.public.*`},
			want: map[string]interface{}{
				`databases`: []string{`database`},
				`schemas`: []string{`public`},
			},
		},
		{
			name: `multiple databases and schema`,
			fields: fs,
			args:args{expression:`(database1|database2).public.*`},
			want: map[string]interface{}{
				`databases`: []string{`database1`, `database2`},
				`schemas`: []string{`public`},
			},
		},
		{
			name: `database and multiple schema`,
			fields: fs,
			args:args{expression:`database.(schema1|schema2).*`},
			want: map[string]interface{}{
				`databases`: []string{`database`},
				`schemas`: []string{`schema1`,`schema2`},
			},
		},
		{
			name: `database and exclude schema`,
			fields: fs,
			args:args{expression:`database.(*|!schema2).*`},
			want: map[string]interface{}{
				`databases`: []string{`database`},
				`excluded_schemas`: []string{`schema2`},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				PathSep:      tt.fields.PathSep,
				ListSep:      tt.fields.ListSep,
				NameSep:      tt.fields.NameSep,
				PrefixExclud: tt.fields.PrefixExclud,
				AllSymb:      tt.fields.AllSymb,
			}
			got, err := p.Parse(tt.args.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}