// Copyright © 2025 Mark Summerfield. All rights reserved.

package main

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/mark-summerfield/ini"
	"github.com/mark-summerfield/ufile"
)

type Config struct {
	Filename       string
	CursorBlink    bool
	WindowGeometry []byte
	WindowState    []byte
	MostRecentFile string
	// TODO recent files
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
		config.MostRecentFile = cfg.Str(ini.UNNAMED,
			CONFIG_MOST_RECENT_FILE, "")
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
	cfg.SetStr(CONFIG_WINDOW, CONFIG_WINDOW_GEOMETRY,
		base64.StdEncoding.EncodeToString(me.WindowGeometry))
	cfg.SetStr(CONFIG_WINDOW, CONFIG_WINDOW_STATE,
		base64.StdEncoding.EncodeToString(me.WindowState))
	cfg.SetStr(ini.UNNAMED, CONFIG_MOST_RECENT_FILE, me.MostRecentFile)
	return cfg.SaveFile(me.Filename)
}

func (me *Config) Save() error { return me.SaveTo(me.Filename) }
