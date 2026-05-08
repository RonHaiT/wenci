package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

type config struct {
	rootDir        string
	watchRelDirs   []string
	ignoreSuffixes []string
	debounce       time.Duration
}

func main() {
	cfg, err := newConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("watcher init error: %v", err)
	}
	defer w.Close()

	if err := addRecursiveWatches(w, cfg.rootDir, cfg.watchRelDirs); err != nil {
		log.Fatalf("watch init error: %v", err)
	}

	log.Printf("watching: %s (dirs: %s)", cfg.rootDir, strings.Join(cfg.watchRelDirs, ", "))
	log.Printf("debounce: %s", cfg.debounce)

	var (
		runMu   sync.Mutex
		running bool
		timer   *time.Timer
	)
	resetTimer := func() {
		runMu.Lock()
		defer runMu.Unlock()
		if timer == nil {
			timer = time.AfterFunc(cfg.debounce, func() {
				if err := runGenCtrlOnce(ctx, cfg.rootDir, &runMu, &running); err != nil && !errors.Is(err, context.Canceled) {
					log.Printf("gen ctrl failed: %v", err)
				}
			})
			return
		}
		timer.Reset(cfg.debounce)
	}

	for {
		select {
		case <-ctx.Done():
			log.Printf("stopped")
			return
		case evt, ok := <-w.Events:
			if !ok {
				return
			}
			if shouldIgnoreEvent(cfg, evt) {
				continue
			}
			if evt.Op&fsnotify.Create != 0 {
				_ = tryAddDirWatch(w, evt.Name)
			}
			resetTimer()
		case err, ok := <-w.Errors:
			if !ok {
				return
			}
			log.Printf("watch error: %v", err)
		}
	}
}

func newConfig() (*config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	root := filepath.Clean(cwd)

	return &config{
		rootDir:      root,
		watchRelDirs: []string{"api"},
		ignoreSuffixes: []string{
			".tmp", ".swp", ".swo", ".bak", ".old",
			"~",
		},
		debounce: 600 * time.Millisecond,
	}, nil
}

func addRecursiveWatches(w *fsnotify.Watcher, root string, relDirs []string) error {
	for _, rel := range relDirs {
		abs := filepath.Join(root, rel)
		info, err := os.Stat(abs)
		if err != nil {
			return fmt.Errorf("stat %s: %w", abs, err)
		}
		if !info.IsDir() {
			return fmt.Errorf("%s is not a directory", abs)
		}
		if err := filepath.WalkDir(abs, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				return nil
			}
			if isHiddenDir(path) {
				return filepath.SkipDir
			}
			return w.Add(path)
		}); err != nil {
			return err
		}
	}
	return nil
}

func tryAddDirWatch(w *fsnotify.Watcher, path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return nil
	}
	if !fi.IsDir() {
		return nil
	}
	if isHiddenDir(path) {
		return nil
	}
	return w.Add(path)
}

func isHiddenDir(path string) bool {
	base := filepath.Base(path)
	if base == "." || base == ".." {
		return false
	}
	return strings.HasPrefix(base, ".")
}

func shouldIgnoreEvent(cfg *config, evt fsnotify.Event) bool {
	name := evt.Name
	for _, suf := range cfg.ignoreSuffixes {
		if strings.HasSuffix(name, suf) {
			return true
		}
	}
	base := filepath.Base(name)
	if strings.HasPrefix(base, ".") {
		return true
	}
	if strings.Contains(strings.ToLower(name), string(filepath.Separator)+".git"+string(filepath.Separator)) {
		return true
	}
	return false
}

func runGenCtrlOnce(ctx context.Context, root string, mu *sync.Mutex, running *bool) error {
	mu.Lock()
	if *running {
		mu.Unlock()
		return nil
	}
	*running = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		*running = false
		mu.Unlock()
	}()

	cmd := exec.CommandContext(ctx, "gf", "gen", "ctrl")
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	log.Printf("running: gf gen ctrl")
	start := time.Now()
	err := cmd.Run()
	cost := time.Since(start).Truncate(10 * time.Millisecond)
	if err != nil {
		if errors.Is(ctx.Err(), context.Canceled) {
			return context.Canceled
		}
		return fmt.Errorf("command error (%s): %w", cost, err)
	}
	log.Printf("done (%s)", cost)

	// On Windows, fsnotify can emit additional events in bursts; a short sleep reduces immediate re-trigger.
	if runtime.GOOS == "windows" {
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
