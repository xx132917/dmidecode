package dmidecode

import (
	"github.com/yumaojun03/dmidecode/parser/baseboard"
	"github.com/yumaojun03/dmidecode/parser/bios"
	"github.com/yumaojun03/dmidecode/parser/chassis"
	"github.com/yumaojun03/dmidecode/parser/memory"
	"github.com/yumaojun03/dmidecode/parser/onboard"
	"github.com/yumaojun03/dmidecode/parser/port"
	"github.com/yumaojun03/dmidecode/parser/processor"
	"github.com/yumaojun03/dmidecode/parser/slot"
	"github.com/yumaojun03/dmidecode/parser/system"
	"github.com/yumaojun03/dmidecode/smbios"
)

// New 实例化
func New() (*Decoder, error) {
	ss, err := smbios.ReadStructures()
	if err != nil {
		return nil, err
	}

	d := new(Decoder)
	d.total = len(ss)

	for i := range ss {
		switch smbios.StructureType(ss[i].Header.Type) {
		case smbios.BIOS:
			d.bios = append(d.bios, ss[i])
		case smbios.System:
			d.system = append(d.system, ss[i])
		case smbios.BaseBoard:
			d.baseBoard = append(d.baseBoard, ss[i])
		case smbios.Chassis:
			d.chassis = append(d.chassis, ss[i])
		case smbios.OnBoardDevices:
			d.onBoardDevice = append(d.onBoardDevice, ss[i])
		case smbios.OnBoardDevicesExtendedInformation:
			d.onBoardDevices = append(d.onBoardDevices, ss[i])
		case smbios.PortConnector:
			d.portConnector = append(d.portConnector, ss[i])
		case smbios.Processor:
			d.processor = append(d.processor, ss[i])
		case smbios.Cache:
			d.cache = append(d.cache, ss[i])
		case smbios.PhysicalMemoryArray:
			d.physicalMemoryArray = append(d.physicalMemoryArray, ss[i])
		case smbios.MemoryDevice:
			d.memoryDevice = append(d.memoryDevice, ss[i])
		case smbios.SystemSlots:
			d.systemSlots = append(d.systemSlots, ss[i])
		default:
		}
	}

	return d, nil
}

// Decoder decoder
type Decoder struct {
	total int

	bios                []*smbios.Structure
	system              []*smbios.Structure
	baseBoard           []*smbios.Structure
	chassis             []*smbios.Structure
	onBoardDevice       []*smbios.Structure
	onBoardDevices      []*smbios.Structure
	portConnector       []*smbios.Structure
	processor           []*smbios.Structure
	cache               []*smbios.Structure
	physicalMemoryArray []*smbios.Structure
	memoryDevice        []*smbios.Structure
	systemSlots         []*smbios.Structure
}

// BIOS 解析bios信息
func (d *Decoder) BIOS() ([]*bios.Information, error) {
	infos := make([]*bios.Information, 0, len(d.bios))
	for i := range d.bios {
		info, err := bios.Parse(d.bios[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// System 解析system信息
func (d *Decoder) System() ([]*system.Information, error) {
	infos := make([]*system.Information, 0, len(d.system))
	for i := range d.system {
		info, err := system.Parse(d.system[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// BaseBoard 解析baseboard信息
func (d *Decoder) BaseBoard() ([]*baseboard.Information, error) {
	infos := make([]*baseboard.Information, 0, len(d.baseBoard))
	for i := range d.baseBoard {
		info, err := baseboard.Parse(d.baseBoard[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Chassis 解析chassis信息
func (d *Decoder) Chassis() ([]*chassis.Information, error) {
	infos := make([]*chassis.Information, 0, len(d.chassis))
	for i := range d.chassis {
		info, err := chassis.Parse(d.chassis[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Onboard 解析onboard信息
func (d *Decoder) Onboard() ([]*onboard.ExtendedInformation, error) {
	infos := make([]*onboard.ExtendedInformation, 0, len(d.onBoardDevices))
	for i := range d.onBoardDevices {
		info, err := onboard.Parse(d.onBoardDevices[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Onboard 解析onboard信息
func (d *Decoder) OnboardDevice() ([]string, error) {
	infos := make([]string, 0, len(d.onBoardDevices))
	for i := range d.onBoardDevice {
		if d.onBoardDevice[i].Strings == nil {
			continue
		}
		for _, desc := range d.onBoardDevice[i].Strings {
			infos = append(infos, desc)
		}
	}

	return infos, nil
}

// PortConnector 解析port connector信息
func (d *Decoder) PortConnector() ([]*port.Information, error) {
	infos := make([]*port.Information, 0, len(d.portConnector))
	for i := range d.portConnector {
		info, err := port.Parse(d.portConnector[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Processor 解析processor信息
func (d *Decoder) Processor() ([]*processor.Processor, error) {
	infos := make([]*processor.Processor, 0, len(d.processor))
	for i := range d.processor {
		info, err := processor.ParseProcessor(d.processor[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// ProcessorCache 解析processor cache信息
func (d *Decoder) ProcessorCache() ([]*processor.Cache, error) {
	infos := make([]*processor.Cache, 0, len(d.cache))
	for i := range d.cache {
		info, err := processor.ParseCache(d.cache[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// MemoryArray 解析memory array信息
func (d *Decoder) MemoryArray() ([]*memory.PhysicalMemoryArray, error) {
	infos := make([]*memory.PhysicalMemoryArray, 0, len(d.physicalMemoryArray))
	for i := range d.physicalMemoryArray {
		info, err := memory.ParseMemoryArray(d.physicalMemoryArray[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// MemoryDevice 解析memory device信息
func (d *Decoder) MemoryDevice() ([]*memory.MemoryDevice, error) {
	infos := make([]*memory.MemoryDevice, 0, len(d.memoryDevice))
	for i := range d.memoryDevice {
		info, err := memory.ParseMemoryDevice(d.memoryDevice[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// Slot 解析memory device信息
func (d *Decoder) Slot() ([]*slot.SystemSlot, error) {
	infos := make([]*slot.SystemSlot, 0, len(d.systemSlots))
	for i := range d.systemSlots {
		info, err := slot.Parse(d.systemSlots[i])
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// ALL decode all
func (d *Decoder) ALL() (*InformationSet, error) {
	errs := NewErrorSet()
	sets := NewInformationSet()

	biosInfos, err := d.BIOS()
	errs.checkOrAdd(err)
	sets.addBios(biosInfos)

	systemInfos, err := d.System()
	errs.checkOrAdd(err)
	sets.addSystem(systemInfos)

	bbInfos, err := d.BaseBoard()
	errs.checkOrAdd(err)
	sets.addBaseBoard(bbInfos)

	csInfos, err := d.Chassis()
	errs.checkOrAdd(err)
	sets.addChassis(csInfos)

	obInfos, err := d.Onboard()
	errs.checkOrAdd(err)
	sets.addOnboard(obInfos)

	pcInfos, err := d.PortConnector()
	errs.checkOrAdd(err)
	sets.addPortConnector(pcInfos)

	processorInfos, err := d.Processor()
	errs.checkOrAdd(err)
	sets.addProcessor(processorInfos)

	pcacheInfos, err := d.ProcessorCache()
	errs.checkOrAdd(err)
	sets.addCache(pcacheInfos)

	maInfos, err := d.MemoryArray()
	errs.checkOrAdd(err)
	sets.addMemoryArray(maInfos)

	mdInfos, err := d.MemoryDevice()
	errs.checkOrAdd(err)
	sets.addMemoryDevice(mdInfos)

	slotsInfos, err := d.Slot()
	errs.checkOrAdd(err)
	sets.addSlot(slotsInfos)

	return sets, errs.Error()
}
