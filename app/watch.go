package app

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/axllent/golp/utils"
	"github.com/radovskyb/watcher"
)

// ByLen struct for sorting shortest to longest
type ByLen []string

func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// WatchSrcDirs will start a watcher at the highest level of the src directories
func WatchSrcDirs() {

	watcherMap := []watchMap{}

	if err := DeleteDistDirs(); err != nil {
		Log().Errorf("Error deleting dist directories: %s", err)
		os.Exit(1)
	}

	for _, proc := range Conf.Process {
		mapPaths := make(map[string]bool)

		for _, src := range proc.Src {
			dir := filepath.Dir(src)

			mapPaths[dir] = true
		}

		paths := []string{}

		for p := range mapPaths {
			if !utils.IsDir(p) {
				Log().Warnf("'%s' directory \"%s\" not found, ignoring", proc.Type, p)
			} else {
				paths = append(paths, p)
			}
		}

		// sort slice from shortest to longest to process the topDir correctly
		sort.Sort(ByLen(paths))

		// find the highest level directory for each src slice to watch
		paths = topDir(paths)

		for _, p := range paths {
			a, err := filepath.Abs(p)
			if err != nil {
				panic(err)
			}
			watcherMap = append(watcherMap, watchMap{
				Path:          a,
				ProcessStruct: proc,
			})
		}

		if err := proc.Process(); err != nil {
			Log().Errorf("Error processing: %s", err)
		}
	}

	w := watcher.New()

	// if SetMaxEvents is not set, the default is to send all events
	// ie: if a directory is renamed then multiple events are sent
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case event := <-w.Event:
				for _, w := range watcherMap {
					if event.Path == w.Path || strings.HasPrefix(event.Path, w.Path+string(os.PathSeparator)) {
						if err := w.ProcessStruct.Process(); err != nil {
							Log().Error(err.Error())
						}
					}
				}
			case err := <-w.Error:
				Log().Error(err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	for _, wm := range watcherMap {
		if err := w.AddRecursive(wm.Path); err != nil {
			Log().Error(err.Error())
		}

		Log().Debugf("watching %s", wm.Path)
	}

	Log().Info("watching for changes...")

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		Log().Error(err.Error())
	}
}

// contains will return true if the string has any matching
// parent folder in a slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.HasPrefix(e, a+string(os.PathSeparator)) {
			return true
		}
	}
	return false
}

func topDir(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		if !contains(list, item) {
			list = append(list, item)
		}
	}
	return list
}
