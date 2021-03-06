package ble

import (
	"bytes"

	"gobot.io/x/gobot"
)

var _ gobot.Driver = (*DeviceInformationDriver)(nil)

// DeviceInformationDriver represents the Device Information Service for a BLE Peripheral
type DeviceInformationDriver struct {
	name       string
	connection gobot.Connection
	gobot.Eventer
}

// NewDeviceInformationDriver creates a DeviceInformationDriver
func NewDeviceInformationDriver(a *ClientAdaptor) *DeviceInformationDriver {
	n := &DeviceInformationDriver{
		name:       "DeviceInformation",
		connection: a,
		Eventer:    gobot.NewEventer(),
	}

	return n
}

// Connection returns the Driver's Connection to the associated Adaptor
func (b *DeviceInformationDriver) Connection() gobot.Connection { return b.connection }

// Name returns the Driver name
func (b *DeviceInformationDriver) Name() string { return b.name }

// SetName sets the Driver name
func (b *DeviceInformationDriver) SetName(n string) { b.name = n }

// adaptor returns BLE adaptor for this device
func (b *DeviceInformationDriver) adaptor() *ClientAdaptor {
	return b.Connection().(*ClientAdaptor)
}

// Start tells driver to get ready to do work
func (b *DeviceInformationDriver) Start() (err error) {
	return
}

// Halt stops driver (void)
func (b *DeviceInformationDriver) Halt() (err error) { return }

// GetModelNumber returns the model number for the BLE Peripheral
func (b *DeviceInformationDriver) GetModelNumber() (model string) {
	c, _ := b.adaptor().ReadCharacteristic("180a", "2a24")
	buf := bytes.NewBuffer(c)
	val := buf.String()
	return val
}

// GetFirmwareRevision returns the firmware revision for the BLE Peripheral
func (b *DeviceInformationDriver) GetFirmwareRevision() (revision string) {
	c, _ := b.adaptor().ReadCharacteristic("180a", "2a26")
	buf := bytes.NewBuffer(c)
	val := buf.String()
	return val
}

// GetHardwareRevision returns the hardware revision for the BLE Peripheral
func (b *DeviceInformationDriver) GetHardwareRevision() (revision string) {
	c, _ := b.adaptor().ReadCharacteristic("180a", "2a27")
	buf := bytes.NewBuffer(c)
	val := buf.String()
	return val
}

// GetManufacturerName returns the manufacturer name for the BLE Peripheral
func (b *DeviceInformationDriver) GetManufacturerName() (manufacturer string) {
	c, _ := b.adaptor().ReadCharacteristic("180a", "2a29")
	buf := bytes.NewBuffer(c)
	val := buf.String()
	return val
}

// GetPnPId returns the PnP ID for the BLE Peripheral
func (b *DeviceInformationDriver) GetPnPId() (model string) {
	c, _ := b.adaptor().ReadCharacteristic("180a", "2a50")
	buf := bytes.NewBuffer(c)
	val := buf.String()
	return val
}
