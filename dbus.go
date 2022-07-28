package gonnarain

import (
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

const (
	dbusName     = "com.github.centodev.gonnaraind"
	dbusNamePath = "/com/github/centodev/gonnaraind"
)

const intro = `
<node>
	<interface name="` + dbusName + `">
		<method name="DisasterAlert">
			<arg direction="in" type="s"/>
		</method>
	</interface>` + introspect.IntrospectDataString + `</node> `

type DbusInterface interface {
	DisasterAlert(msg string) *dbus.Error
}

func exportDbusInterface(conn *dbus.Conn, idbus DbusInterface) error {
	err := conn.Export(idbus, dbusNamePath, dbusName)

	if err != nil {
		return fmt.Errorf("dbus export failed: %w", err)
	}

	err = conn.Export(introspect.Introspectable(intro), dbusNamePath,
		"org.freedesktop.DBus.Introspectable")

	if err != nil {
		return fmt.Errorf("dbus export failed: %w", err)
	}

	reply, err := conn.RequestName(dbusName,
		dbus.NameFlagDoNotQueue)

	if err != nil {
		return fmt.Errorf("dbus request name failed: %w", err)
	}

	if reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("dbus name already taken")
	}

	return nil
}
