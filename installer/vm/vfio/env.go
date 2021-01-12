package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	vfio "workspace/linux/environment/vfio-utils"
)

var (
	iommuGroup = flag.Int("iommu-group", 0, "IOMMU group ID")
	pciAddress = flag.String("pci-address", "", "PCI address of vfio device")
)

func main() {

	flag.Parse()
	if (*iommuGroup == 0) || (*pciAddress == "") {
		fmt.Printf("Missing input parameters, exiting...")
		os.Exit(1)
	}
	groupStatus := vfio.GroupStatus{
		Argsz: uint32(unsafe.Sizeof(vfio.GroupStatus{})),
	}

	// Attempting to open /dev/vfio/vfio
	container, err := syscall.Open("/dev/vfio/vfio", syscall.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("Something happened while opening /dev/vfio/vfio, error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Open container succeeded, handle: %d\n", container)

	groupPath := fmt.Sprintf("/dev/vfio/%d", *iommuGroup)
	group, err := syscall.Open(groupPath, syscall.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("Something happened while opening %s, error: %+v\n", groupPath, err)
		os.Exit(1)
	}
	// Status before setting group's container
	fmt.Printf("Open group succeeded, handle: %d\n", group)
	if err := vfio.GetGroupStatus(group, &groupStatus); err != nil {
		fmt.Printf("Fail to get group status with error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group: %d status Flags: %b before setting up to a container: %d\n", group, groupStatus.Flags, container)
	if (groupStatus.Flags & vfio.VFIO_GROUP_FLAGS_VIABLE) != vfio.VFIO_GROUP_FLAGS_VIABLE {
		fmt.Printf("The group is not viable, exiting...\n")
		os.Exit(1)
	}
	// Setting up group's container
	if err := vfio.SetGroupContainer(group, container); err != nil {
		fmt.Printf("Fail to set group's container with error: %+v\n", err)
		os.Exit(1)
	}
	// Status after setting group's container
	if err := vfio.GetGroupStatus(group, &groupStatus); err != nil {
		fmt.Printf("Fail to get group status with error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group: %d status Flags: %b after setting up to a container: %d\n", group, groupStatus.Flags, container)
	if (groupStatus.Flags & vfio.VFIO_GROUP_FLAGS_VIABLE) != vfio.VFIO_GROUP_FLAGS_VIABLE {
		fmt.Printf("The group is not viable, exiting...\n")
		os.Exit(1)
	}

	found, err := vfio.CheckExtension(container, vfio.VFIO_TYPE1_IOMMU)
	if err != nil {
		fmt.Printf("Failed to check for supported extension: %04x with error: %+v\n", vfio.VFIO_TYPE1_IOMMU, err)
		os.Exit(1)
	}
	if found {
		fmt.Printf("Device: %d supports VFIO_TYPE1_IOMMU\n", container)
		if err := vfio.SetContainerIOMMU(container, vfio.VFIO_TYPE1_IOMMU); err != nil {
			fmt.Printf("Failed to set container %d for VFIO_TYPE1_IOMMU type with error: %+v\n", container, err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Device: %d does not support VFIO_TYPE1_IOMMU\n", container)
		if err := vfio.SetContainerIOMMU(container, vfio.VFIO_NOIOMMU_IOMMU); err != nil {
			fmt.Printf("Failed to set container %d for VFIO_NOIOMMU_IOMMU type with error: %+v\n", container, err)
			os.Exit(1)
		}
	}

	if err := vfio.GetGroupStatus(group, &groupStatus); err != nil {
		fmt.Printf("Fail to get group status with error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group %d status: %+v Flags: %b \n", group, groupStatus, groupStatus.Flags)
	device, err := vfio.GetGroupFD(group, *pciAddress)
	if err != nil {
		fmt.Printf("Fail to get group file descriptor %+v.\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group %d file descriptor is %d\n", group, device)
	deviceInfo := vfio.DeviceInfo{
		Argsz: uint32(unsafe.Sizeof(vfio.DeviceInfo{})),
	}
	if err := vfio.GetDeviceInfo(device, &deviceInfo); err != nil {
		fmt.Printf("Fail to get device info with error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Group %d device info: %+v\n", group, deviceInfo)

	for r := uint32(0); r < deviceInfo.NumRegions; r++ {
		regionInfo := vfio.RegionInfo{
			Argsz: uint32(unsafe.Sizeof(vfio.RegionInfo{})),
			Index: r,
		}
		if err := vfio.GetRegionInfo(device, &regionInfo); err != nil {
			fmt.Printf("Fail to get region %d info with error: %+v\n", r, err)
			continue
		}
		fmt.Printf("Region %d info: Flags: %08b Size: %#x Offset: %#x\n", r, regionInfo.Flags, regionInfo.Size, regionInfo.Offset)
	}

	for i := uint32(0); i < deviceInfo.NumIRQs; i++ {
		irqInfo := vfio.IRQInfo{
			Argsz: uint32(unsafe.Sizeof(vfio.IRQInfo{})),
			Index: i,
		}
		if err := vfio.GetIRQInfo(device, &irqInfo); err != nil {
			fmt.Printf("Fail to get IRQ index %d info with error: %+v\n", i, err)
			continue
		}
		fmt.Printf("IRQ index %d info: Flags: %08b Count: %d\n", i, irqInfo.Flags, irqInfo.Count)
	}
}
