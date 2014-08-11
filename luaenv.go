// Application to quickly setup a lua environment
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/aarzilli/golua/lua"
    "github.com/howeyc/fsnotify"
)

// List of lua files we are watching
type luas map[string]status

// States of individual lua files
type status int

// possible states of Luas
const (
    lua_loaded status = iota
    ignored
    edited
)

// how long to wait before loading updated Luas
const waitDelta = time.Duration(time.Second)

func main() {

    L := lua.NewState()
    L.OpenLibs()
    defer L.Close()

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        panic(err)
    }

    luasPath := "./"
    if len(os.Args) == 2 {
        luasPath = os.Args[1]
    }

    watcher.WatchFlags(luasPath, fsnotify.FSN_MODIFY)

    ls := luas{}

    var ev *fsnotify.FileEvent
    ls.Collect(luasPath)
    ls.Load(L)

newev:
    ls.Handle(ev, L)
    ev = <-watcher.Event
    goto newev
}

// Collect an initial list of lua files in the arguments directory
func (ls luas) Collect(base string) {

    f, err := os.Open(base)
    if err != nil {
        panic(err)
    }

    files, err := f.Readdir(0)

    for _, f := range files {

        if !isLuaFile(f.Name()) {
            continue
        }

        ls[filepath.Join(base, f.Name())] = edited
    }

    return
}

// Locker for us to queue multiple fsnotify events
var locker *time.Timer

// Handle fsnotify events, queues multiple events for waitDelta,
// then it reloads updated luas.
func (ls luas) Handle(ev *fsnotify.FileEvent, L *lua.State) {

    if ev == nil {
        return
    }

    if !isLuaFile(ev.Name) {
        return
    }

    ls[ev.Name] = edited

    if locker == nil {
        locker = time.NewTimer(waitDelta)
        go func() {
            <-locker.C
            ls.Load(L)
            locker = nil
        }()
    } else {
        locker.Reset(waitDelta)
    }

}

// Load all edited luas and run them into our state
func (ls luas) Load(L *lua.State) (err error) {

    for path, status := range ls {
        if status == edited {
            err := L.DoFile(path)
            if err != nil {
                fmt.Println("error: ", err)
            }
            ls[path] = lua_loaded
        }
    }

    return nil
}

// UTIL
func isLuaFile(path string) bool {
    return filepath.Ext(path) == ".lua"
}
