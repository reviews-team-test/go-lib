package introspect

import (
	"bytes"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testXML = `
<node>
  <interface name="org.bluez.Device1">
    <method name="Disconnect"/>
    <method name="Connect"/>
    <method name="ConnectProfile">
      <annotation name="org.freedesktop.DBus.GLib.CSymbol" value="impl_manager_activate_connection"/>
      <annotation name="org.freedesktop.DBus.GLib.Async" value=""/>
      <arg name="UUID" type="s" direction="in">
      	<annotation name="com.deepin.DBus.I18n" value="true"/>
      </arg>
    </method>
    <method name="DisconnectProfile">
      <arg name="UUID" type="s" direction="in"/>
    </method>
    <method name="Pair">
      <annotation name="org.freedesktop.DBus.Method.NoReply" value="true"/>
    </method>
    <method name="CancelPairing"/>

    <property name="Address" type="s" access="read">
      <annotation name="com.deepin.DBus.I18n" value="true"/>
    </property>
    <property name="Name" type="s" access="read">
      <annotation name="com.deepin.DBus.I18n" value="false"/>
    </property>
    <property name="Alias" type="s" access="readwrite"/>
    <property name="Class" type="u" access="read"/>
    <property name="Appearance" type="q" access="read"/>
    <property name="Icon" type="s" access="read"/>
    <property name="Paired" type="b" access="read"/>
    <property name="Trusted" type="b" access="readwrite"/>
    <property name="Blocked" type="b" access="readwrite"/>
    <property name="LegacyPairing" type="b" access="read"/>
    <property name="RSSI" type="n" access="read"/>
    <property name="Connected" type="b" access="read"/>
    <property name="UUIDs" type="as" access="read"/>
    <property name="Modalias" type="s" access="read"/>
    <property name="Adapter" type="o" access="read"/>
  </interface>
</node>
`

func TestParse(t *testing.T) {
	C.Convey("Create a reader ", t, func() {
		reader := bytes.NewBufferString(testXML)
		ninfo, err := Parse(reader)
		C.So(err, C.ShouldBeNil)
		C.Convey("Check interfaces", func() {
			C.So(len(ninfo.Interfaces), C.ShouldEqual, 1)
			ifc := ninfo.Interfaces[0]
			C.So(ifc.Name, C.ShouldEqual, "org.bluez.Device1")
			C.So(len(ifc.Methods), C.ShouldEqual, 6)
			C.So(len(ifc.Properties), C.ShouldEqual, 15)

			C.Convey("Check annotations", func() {
				m := ifc.Methods[2]
				C.So(m.Name, C.ShouldEqual, "ConnectProfile")
				C.So(m.Annotations[0].Name, C.ShouldEqual, "org.freedesktop.DBus.GLib.CSymbol")
				C.So(m.Annotations[0].Value, C.ShouldEqual, "impl_manager_activate_connection")
				C.So(m.Args[0].I18nField(), C.ShouldEqual, true)

				C.So(ifc.Methods[4].Name, C.ShouldEqual, "Pair")
				C.So(ifc.Methods[4].NoReply(), C.ShouldEqual, true)

				C.So(ifc.Methods[3].Args[0].I18nField(), C.ShouldEqual, false)

				p := ifc.Properties[0]
				C.So(p.Name, C.ShouldEqual, "Address")
				C.So(p.I18nField(), C.ShouldEqual, true)

				C.So(ifc.Properties[1].I18nField(), C.ShouldEqual, false)
				C.So(ifc.Properties[2].I18nField(), C.ShouldEqual, false)
				C.So(m.NoReply(), C.ShouldEqual, false)
			})
		})
	})
}
