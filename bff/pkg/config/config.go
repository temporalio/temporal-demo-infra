package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	//"github.com/kelseyhightower/envconfig"
	"github.com/vrischmann/envconfig"
)

var configFile = ".env"

type prefixer interface {
	Prefix() string
}
type overrider interface {
	Override() error
}

func MustLoad(filenames ...string) {
	if len(filenames) == 0 {
		filenames = []string{configFile}
	}
	// optionally load .env file
	if err := Load(filenames...); err != nil {
		log.Fatalln("Error loading:", filenames)
	}
}
func Load(filenames ...string) error {
	// optionally load .env file
	if err := godotenv.Load(filenames...); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to load %v: %w", filenames, err)
	}
	return nil
}

func MustUnmarshalConfig(obj prefixer) {
	_, err := UnmarshalConfig(obj)
	if err != nil {
		panic(err.Error())
	}
}
func UnmarshalConfig(obj prefixer) (prefixer, error) {
	if obj == nil {
		return nil, fmt.Errorf("prefixer cannot be nil")
	}
	if err := envconfig.InitWithOptions(obj, envconfig.Options{
		Prefix:          obj.Prefix(),
		AllOptional:     true,
		LeaveNil:        true,
		AllowUnexported: false,
	}); err != nil {
		return nil, fmt.Errorf("failed to process %T from environment with prefix %s: %w", obj, obj.Prefix(), err)
	}

	if ovr, ok := obj.(overrider); ok {
		if err := ovr.Override(); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func MustUnmarshalAll(cfg interface{}) {
	val := reflect.ValueOf(cfg)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		fv := f.Interface()

		if f.IsNil() {
			// this creates a new instance of the field and assigns it
			fvv := reflect.New(f.Type().Elem())
			f.Set(fvv)
			fv = fvv.Interface()
		}

		if p, ok := fv.(prefixer); ok {
			MustUnmarshalConfig(p)
		}

	}
}
