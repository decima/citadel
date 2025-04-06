package logger

import (
	"citadel-api/utils/build"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"
)

var once sync.Once

var log *zerolog.Logger

func Get() *zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.DebugLevel) // default to INFO
		}

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		var gitRevision string = "dev"

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		loggerInstance := zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Str("version", build.Version()).
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
		log = &loggerInstance
	})

	return log
}
