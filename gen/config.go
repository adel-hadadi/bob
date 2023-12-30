package gen

import (
	"github.com/stephenafamo/bob/gen/drivers"
)

// Config for the running of the commands
type Config struct {
	// Struct tags to generate
	Tags []string `yaml:"tags"`
	// Disable generating factory for models.
	NoFactory bool `yaml:"no_factory"`
	// Disable generated go test files
	NoTests bool `yaml:"no_tests"`
	// Disable back referencing in the loaded relationship structs
	NoBackReferencing bool `yaml:"no_back_referencing"`
	// Delete the output folder (rm -rf) before generation to ensure sanity
	Wipe bool `yaml:"wipe"`
	// Decides the casing for go structure tag names. camel, title or snake (default snake)
	StructTagCasing string `yaml:"struct_tag_casing"`
	// Relationship struct tag name
	RelationTag string `yaml:"relation_tag"`
	// List of column names that should have tags values set to '-' (ignored during parsing)
	TagIgnore []string `yaml:"tag_ignore"`

	Aliases       Aliases       `yaml:"aliases"`
	Replacements  []Replace     `yaml:"replacements"`
	Relationships Relationships `yaml:"relationships"`
	Inflections   Inflections   `yaml:"inflections"`

	// Customize the generator name in the top level comment of generated files
	// >>   Code generated by **GENERATOR NAME**. DO NOT EDIT.
	// defaults to "BobGen [driver] [version]"
	Generator string `yaml:"generator" toml:"generator" json:"generator"`
}

// Replace replaces a column type with something else
type Replace struct {
	Tables  []string       `yaml:"tables" toml:"tables" json:"tables"`
	Match   drivers.Column `yaml:"match" toml:"match" json:"match"`
	Replace drivers.Column `yaml:"replace" toml:"replace" json:"replace"`
}

type Inflections struct {
	Plural        map[string]string `yaml:"plural"`
	PluralExact   map[string]string `yaml:"plural_exact"`
	Singular      map[string]string `yaml:"singular"`
	SingularExact map[string]string `yaml:"singular_exact"`
	Irregular     map[string]string `yaml:"irregular"`
}
