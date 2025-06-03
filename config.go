// Copyright © 2025 Mark Summerfield. All rights reserved.

package main

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mark-summerfield/ini"
	"github.com/mark-summerfield/ufile"
	"github.com/mark-summerfield/unum"
)

type Config struct {
	Filename       string
	CursorBlink    bool
	WindowGeometry []byte
	WindowState    []byte
	MostRecentFile string
	RecentFiles    RecentFiles
}

func NewConfig(filename string) *Config {
	return &Config{Filename: filename, CursorBlink: DEFAULT_CURSOR_BLINK}
}

func NewConfigFrom(filename string) *Config {
	config := NewConfig(filename)
	cfg := ini.NewIni()
	if err := cfg.MergeFile(filename); err != nil {
		log.Printf("failed to read config in %q: %v\n", filename, err)
	} else {
		config.CursorBlink = cfg.Bool(ini.UNNAMED, CONFIG_CURSOR_BLINK,
			DEFAULT_CURSOR_BLINK)
		raw := cfg.Str(CONFIG_WINDOW, CONFIG_WINDOW_GEOMETRY, "")
		if raw != "" {
			if geometry, err := base64.StdEncoding.DecodeString(
				raw); err == nil {
				config.WindowGeometry = geometry
			}
		}
		raw = cfg.Str(CONFIG_WINDOW, CONFIG_WINDOW_STATE, "")
		if raw != "" {
			if state, err := base64.StdEncoding.DecodeString(
				raw); err == nil {
				config.WindowState = state
			}
		}
		filename := cfg.Str(ini.UNNAMED, CONFIG_MOST_RECENT_FILE, "")
		if filename != "" && ufile.FileExists(filename) {
			config.MostRecentFile = filename
		}
		maximum := unum.Clamp(0, cfg.Int(ini.UNNAMED,
			CONFIG_RECENT_FILES, DEFAULT_MAX_RECENT_FILES),
			DEFAULT_MAX_RECENT_FILES)
		config.RecentFiles = NewRecentFiles(maximum)
		// .Add() adds newest first so we must add in reverse order
		for i := config.RecentFiles.maximum; i > 0; i-- {
			filename := cfg.Str(CONFIG_RECENT_FILES,
				CONFIG_RECENT_FILE+strconv.Itoa(i), "")
			if filename != "" {
				config.RecentFiles.Add(filename)
			}
		}

	}
	return config
}

func (me *Config) SaveTo(filename string) error {
	me.Filename = filename
	if me.Filename == "" {
		return errors.New("no filename to save to")
	}
	if dir := filepath.Dir(me.Filename); !ufile.PathExists(dir) {
		if err := os.MkdirAll(dir, 0o750); err != nil {
			return err
		}
	}
	cfg := ini.NewIni()
	cfg.SetBool(ini.UNNAMED, CONFIG_CURSOR_BLINK, me.CursorBlink)
	cfg.SetComment(ini.UNNAMED, CONFIG_CURSOR_BLINK, "true or false")
	cfg.SetStr(CONFIG_WINDOW, CONFIG_WINDOW_GEOMETRY,
		base64.StdEncoding.EncodeToString(me.WindowGeometry))
	cfg.SetStr(CONFIG_WINDOW, CONFIG_WINDOW_STATE,
		base64.StdEncoding.EncodeToString(me.WindowState))
	cfg.SetStr(ini.UNNAMED, CONFIG_MOST_RECENT_FILE, me.MostRecentFile)
	cfg.SetInt(ini.UNNAMED, CONFIG_MAX_RECENT_FILES, me.RecentFiles.maximum)
	cfg.SetComment(ini.UNNAMED, CONFIG_MAX_RECENT_FILES, "0-9")
	i := 1
	for _, filename := range me.RecentFiles.Files() {
		if ufile.FileExists(filename) {
			cfg.SetStr(CONFIG_RECENT_FILES,
				CONFIG_RECENT_FILE+strconv.Itoa(i), filename)
			i++
		}
	}
	return cfg.SaveFile(me.Filename)
}

func (me *Config) Save() error { return me.SaveTo(me.Filename) }
