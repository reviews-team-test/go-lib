/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package sound

// #cgo pkg-config: glib-2.0 libcanberra
// #include <stdlib.h>
// #include "player.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// PlayThemeSound: play sound theme event
// @theme : the special sound theme
// @event : the special theme event id
// @device: the special backend card, default: the current sound card
// @driver: the special backend driver, as: 'pulse','alsa','null' ...
//          default: 'pulse'
//
// return: the error message
func PlayThemeSound(theme, event, device, driver string) error {
	cTheme := C.CString(theme)
	defer C.free(unsafe.Pointer(cTheme))
	cEvent := C.CString(event)
	defer C.free(unsafe.Pointer(cEvent))
	cDevice := C.CString(device)
	defer C.free(unsafe.Pointer(cDevice))
	cDriver := C.CString(driver)
	defer C.free(unsafe.Pointer(cDriver))

	ret := C.canberra_play_system_sound(cTheme, cEvent,
		cDevice, cDriver)
	if ret != 0 {
		msg := C.GoString(C.ca_strerror(ret))
		return fmt.Errorf("Play failed: %s", msg)
	}
	return nil
}

// PlaySoundFile: play sound file
// @file  : the file which needed to play
// @device: the special backend card, default: the current sound card
// @driver: the special backend driver, as: 'pulse','alsa','null' ...
//          default: 'pulse'
//
// return: the error message
func PlaySoundFile(file, device, driver string) error {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	cDevice := C.CString(device)
	defer C.free(unsafe.Pointer(cDevice))
	cDriver := C.CString(driver)
	defer C.free(unsafe.Pointer(cDriver))

	ret := C.canberra_play_sound_file(cFile, cDevice, cDriver)
	if ret != 0 {
		msg := C.GoString(C.ca_strerror(ret))
		return fmt.Errorf("Play failed: %s", msg)
	}
	return nil
}