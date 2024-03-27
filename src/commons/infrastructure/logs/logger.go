package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
    once.Do(func() {
        zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
        zerolog.TimeFieldFormat = time.RFC3339Nano

        logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
        if err != nil {
            logLevel = int(zerolog.InfoLevel) // default to INFO
        }

        var output io.Writer = zerolog.ConsoleWriter{
            Out:        os.Stdout,
            TimeFormat: time.RFC3339,
        }

        env := os.Getenv("APP_ENV")
        if  env != "production" && env != ""  {
            logPath :=  os.Getenv("LOGS_DIR")

            if logPath == "" {
                logPath = os.Getenv("HOME") + "/var/log/go-app.log"
            }

            fileLogger := &lumberjack.Logger{
                Filename:   logPath,
                MaxSize:    5, //
                MaxBackups: 10,
                MaxAge:     15,
                Compress:   true,
            }

            output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
        }

        log = zerolog.New(output).
            Level(zerolog.Level(logLevel)).
            With().
            Timestamp().
            Logger()
    })

    return log
}