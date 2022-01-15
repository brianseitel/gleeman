package watcher

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type Watcher struct {
	Logger  *zap.Logger
	Paths   []string
	Handler func() error
}

func New(paths []string, handler func() error) Watcher {
	logger, _ := zap.NewDevelopment()
	return Watcher{
		Paths:   paths,
		Logger:  logger,
		Handler: handler,
	}
}

func (w Watcher) Start() {
	// creates a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// Lets us know when we're done
	done := make(chan bool)

	// Start the listener
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				w.Logger.Sugar().Infof("Detected change in file %s", event.Name)

				w.Logger.Sugar().Infof("Rebuilding files...")
				err := w.Handler()
				if err != nil {
					w.Logger.Sugar().Error(err)
				}
				// watch for errors
			case err := <-watcher.Errors:
				w.Logger.Sugar().Errorf("ERROR %v\n", err)
			}
		}
	}()

	// Add paths to the listener
	for _, path := range w.Paths {
		// out of the box fsnotify can watch a single file, or a single directory
		if err := watcher.Add(path); err != nil {
			fmt.Println("ERROR", err)
		}
	}

	// Done!
	<-done
}
