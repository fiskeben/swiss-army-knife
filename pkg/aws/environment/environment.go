package environment

import (
	"os"
	"github.com/spf13/pflag"
)

type EnvSettings struct {
	Quiet bool
}

func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.BoolVarP(&s.Quiet, "quiet", "q", false, "make it staph screaming 🤫")
}

func (s *EnvSettings) Init(fs *pflag.FlagSet) {
	for name, envar := range envMap {
		setFlagFromEnv(name, envar, fs)
	}
}

var envMap = map[string]string{
	"quiet": "SAK_AWS_QUIET",
}

func setFlagFromEnv(name, envar string, fs *pflag.FlagSet) {
	if fs.Changed(name) {
		return
	}
	if v, ok := os.LookupEnv(envar); ok {
		fs.Set(name, v)
	}
}
