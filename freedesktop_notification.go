package gonnarain

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func sendNotification(conn *dbus.Conn, msg string) error {
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "", uint32(0),
		"", "It's going to rain?", msg, []string{},
		map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		return fmt.Errorf("notification send failed: %w", call.Err)
	}

	return nil
}
