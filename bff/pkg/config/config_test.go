package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const defaultValue = "blink"

type TestConfigWithOverride struct {
	CustomThing string
}

func (c *TestConfigWithOverride) Prefix() string {
	return "customtype"
}
func (c *TestConfigWithOverride) Override() error {
	if c.CustomThing == "" {
		c.CustomThing = "overriddenvalue"
	}
	return nil
}

type TestConfig struct {
	UntaggedField             string
	FieldWithDefaultValue     string `envconfig:"default=blink"`
	FieldWithNameOtherThanEnv string `envconfig:"BONK_POO"`
	DunderProp                string
}

func (c *TestConfig) Prefix() string {
	return "bonk"
}

func TestLoad(t *testing.T) {
	A := assert.New(t)
	err := Load("test.env")
	A.NoError(err)
	A.Equal(os.Getenv("FOO_BAR"), "foobar")
	A.Equal(os.Getenv("BAH_QUX"), "boom")
}

func TestUnmarshal(t *testing.T) {
	A := assert.New(t)
	os.Setenv("BONK_UNTAGGEDFIELD", "howdy")
	cfg := TestConfig{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal("howdy", cfg.UntaggedField)
	os.Unsetenv("BONK_UNTAGGEDFIELD")
}

func TestDefaultValue(t *testing.T) {
	A := assert.New(t)
	cfg := TestConfig{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal(defaultValue, cfg.FieldWithDefaultValue)
}
func TestExplicitValue(t *testing.T) {
	A := assert.New(t)
	os.Setenv("BONK_POO", "tada")
	cfg := TestConfig{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal("tada", cfg.FieldWithNameOtherThanEnv)
	os.Unsetenv("BONK_POO")
}
func TestDunderProps(t *testing.T) {
	A := assert.New(t)
	os.Setenv("BONK_DUNDER_PROP", "dunders")
	cfg := TestConfig{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal("dunders", cfg.DunderProp)
	os.Unsetenv("BONK_DUNDER_PROP")
}
func TestUnmarshalAllWithUnconstructedInnerConfigs(t *testing.T) {
	type appCfg struct {
		Bonk *TestConfig
	}
	A := assert.New(t)
	cfg := &appCfg{}
	MustUnmarshalAll(cfg)
	// default value
	A.Equal(defaultValue, cfg.Bonk.FieldWithDefaultValue)
}
func TestUnmarshalAllWithInnerConfigs(t *testing.T) {
	type appCfg struct {
		Bonk *TestConfig
	}
	A := assert.New(t)
	cfg := &appCfg{
		Bonk: &TestConfig{
			UntaggedField: "wink",
		},
	}
	MustUnmarshalAll(cfg)
	// default value
	A.Equal(defaultValue, cfg.Bonk.FieldWithDefaultValue)
	A.Equal("wink", cfg.Bonk.UntaggedField)
}
func TestOverride(t *testing.T) {

	os.Setenv("CUSTOMTYPE_CUSTOM_THING", "myenvvalue")

	A := assert.New(t)
	cfg := TestConfigWithOverride{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal("myenvvalue", cfg.CustomThing)
	os.Unsetenv("CUSTOMTYPE_CUSTOM_THING")

}

func TestOverrideWithNoEnvValue(t *testing.T) {
	A := assert.New(t)
	cfg := TestConfigWithOverride{}
	_, err := UnmarshalConfig(&cfg)
	A.NoError(err)
	A.Equal("overriddenvalue", cfg.CustomThing)

}
