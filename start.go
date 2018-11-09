package webapp

import (
	// #cgo CFLAGS: -I cef
	// #cgo darwin CFLAGS: -x objective-c
	// #cgo darwin LDFLAGS: -framework Cocoa -F cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -Lcef/Release -lcef
	// #include "platform_common.h"
	"C"
	"runtime"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
)

// Start the user interface. This should only be called on the main OS thread.
// Only returns if an error occurs during initialization.
func Start() error {
	runtime.LockOSThread()
	if err := platformPrepareForStart(); err != nil {
		return err
	}
	args := (*C.cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	settings := (*C.cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))
	settings.size = C.sizeof_struct__cef_settings_t
	settings.no_sandbox = 1
	settings.command_line_args_disabled = 1
	if C.cef_initialize(args, settings, nil, nil) != 1 {
		return errs.New("Unable to initialize CEF")
	}
	C.cef_run_message_loop()
	C.cef_shutdown()
	atexit.Exit(0)
	return nil
}

// WillFinishStartupCallback is called right before application startup has
// completed. This is a good point to create any windows your app wants to
// display.
var WillFinishStartupCallback = func() {}

// DidFinishStartupCallback is called once application startup has completed.
var DidFinishStartupCallback = func() {}

// WillActivateCallback is called right before the application is
// activated.
var WillActivateCallback = func() {}

// DidActivateCallback is called once the application is activated.
var DidActivateCallback = func() {}

// WillDeactivateCallback is called right before the application is
// deactivated.
var WillDeactivateCallback = func() {}

// DidDeactivateCallback is called once the application is deactivated.
var DidDeactivateCallback = func() {}
