package views

import "testing"

func TestCleanPgDefault(t *testing.T) {
	cases := []struct {
		expr   string
		want   string
		wantOk bool
	}{
		{`''::character varying`, "", true},
		{`'abc'::character varying`, "abc", true},
		{`'it''s'::text`, "it's", true},
		{`'a b'`, "a b", true},
		{`'2024-01-01 00:00:00+08'::timestamp with time zone`, "2024-01-01 00:00:00+08", true},
		{`'1.50'::numeric(10,2)`, "1.50", true},
		{`'-1'::integer`, "-1", true},
		{`NULL::character varying`, "", false},
		{`NULL`, "", false},
		{`nextval('hg_test_seq'::regclass)`, "", false},
		{`now()`, "", false},
		{`CURRENT_TIMESTAMP`, "", false},
		{`0`, "0", true},
		{`1`, "1", true},
		{`3.14`, "3.14", true},
		{`0::smallint`, "0", true},
		{`true`, "true", true},
		{`false`, "false", true},
		{`1::integer + 2`, "", false},
		{`'a'::text || 'b'`, "", false},
		{`'x'::text::varchar`, "", false},
		{`'abc'::character varying(10)`, "abc", true},
		{`'{}'::jsonb`, "{}", true},
		{`'x'::text[]`, "x", true},
		{`1e3`, "1e3", true},
		{`-1.2E-3`, "-1.2E-3", true},
		{`2E+5`, "2E+5", true},
		{`.5`, ".5", true},
		{`1.`, "1.", true},
		{`+7`, "+7", true},
	}
	for _, c := range cases {
		got, ok := cleanPgDefault(c.expr)
		if got != c.want || ok != c.wantOk {
			t.Errorf("cleanPgDefault(%q) = (%q, %v), want (%q, %v)", c.expr, got, ok, c.want, c.wantOk)
		}
	}
}
