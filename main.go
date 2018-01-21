package main

import (
    "fmt"
    "github.com/godbus/dbus"
    flag "github.com/spf13/pflag"
    "os"
)

func main() {
    var bgImg = flag.StringP("backgroundfile", "b", "", "image file for background")
    var inSrc = flag.StringArrayP("inputsources", "i", []string{}, "list of xkb keyboard names for input sources")
    flag.Parse()

    user := dbusUser(os.Geteuid())
    if *bgImg != "" {
        fmt.Println("TODO: set background")
    }
    showBackgroundFile(user)
    if len(*inSrc) > 0 {
        fmt.Println("TODO: set input sources")
    }
    showInputSources(user)
}

func dbusUser(uid int) dbus.BusObject {
	bus, err := dbus.SystemBus()
	if err != nil { panic(err) }

    var acct = dbus.ObjectPath(fmt.Sprintf("/org/freedesktop/Accounts/User%d", uid))
    return bus.Object("org.freedesktop.Accounts", acct)
}

func showBackgroundFile(user dbus.BusObject) {
	var bgImg string
    err := user.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.Accounts.User", "BackgroundFile").Store(&bgImg)
	if err != nil { panic(err) }
	fmt.Printf("BackgroundFile=%v\n", bgImg)
}

func showInputSources(user dbus.BusObject) {
	var inSrc []map[string]string
    err := user.Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.Accounts.User", "InputSources").Store(&inSrc)
	if err != nil { panic(err) }
	fmt.Printf("InputSources=%v\n", inSrc)
}
