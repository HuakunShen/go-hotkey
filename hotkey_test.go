// Copyright 2021 The golang.design Initiative Authors.
// All rights reserved. Use of this source code is governed
// by a MIT license that can be found in the LICENSE file.
//
// Written by Changkun Ou <changkun.de>

package hotkey_test

import (
	"golang.design/x/hotkey"
	"os"
	"testing"

	"golang.design/x/hotkey/mainthread"
)

// The test cannot be run twice since the mainthread loop may not be terminated:
// go test -v -count=1
func TestMain(m *testing.M) {
	mainthread.Init(func() { os.Exit(m.Run()) })
}

func TestKeyMap(t *testing.T) {
	keyMapKeys := []string{"KeySpace", "Key1", "Key2", "Key3", "Key4", "Key5", "Key6", "Key7", "Key8", "Key9", "Key0", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "KeyReturn", "KeyEscape", "KeyDelete", "KeyTab", "KeyLeft", "KeyRight", "KeyUp", "KeyDown", "KeyF1", "KeyF2", "KeyF3", "KeyF4", "KeyF5", "KeyF6", "KeyF7", "KeyF8", "KeyF9", "KeyF10", "KeyF11", "KeyF12", "KeyF13", "KeyF14", "KeyF15", "KeyF16", "KeyF17", "KeyF18", "KeyF19", "KeyF20"}
	if len(hotkey.KeyMap) != len(keyMapKeys) {
		t.Errorf("Wrong Keymap Size")
	}
	for _, key := range keyMapKeys {
		if _, ok := hotkey.KeyMap[key]; !ok {
			t.Errorf("Key: %s doesn't exist", key)
		}
	}
}
