package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"os"
)

var (
	i int
)

func uiMain() {
	win := ui.NewWindow("Lufus", 300, 450, true)
	defer win.Show()

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	// Device
	vbox.Append(ui.NewLabel("Device"), false)
	driveList := ui.NewCombobox()
	// TODO: Replace it with proper list filtering
	driveList.Append("Pendrive")
	vbox.Append(driveList, false)

	// Partition scheme
	vbox.Append(ui.NewLabel("Partition scheme and target system type"), false)
	partionSchemeList := ui.NewCombobox()
	partionSchemeList.Append("MBR partition scheme for BIOS")
	partionSchemeList.Append("MBR partition for BIOS and UEFI-CSM")
	partionSchemeList.Append("GPT partition for UEFI computer")
	partionSchemeList.SetSelected(2)
	vbox.Append(partionSchemeList, false)

	// Filesystem
	vbox.Append(ui.NewLabel("File system"), false)
	fileSystemList := ui.NewCombobox()
	fileSystemList.Append("FAT32")
	fileSystemList.Append("NTFS")
	fileSystemList.Append("UDF")
	fileSystemList.Append("exFAT")
	fileSystemList.SetSelected(1)
	vbox.Append(fileSystemList, false)

	// Cluster size
	vbox.Append(ui.NewLabel("Cluster size"), false)
	clusterSizeList := ui.NewCombobox()
	clusterSizeList.Append("4096 bytes (default)")
	clusterSizeList.Append("1024 bytes")
	clusterSizeList.SetSelected(0)
	vbox.Append(clusterSizeList, false)

	// New volume label
	vbox.Append(ui.NewLabel("New volume name"), false)
	volumeLabel := ui.NewEntry()
	vbox.Append(volumeLabel, false)

	// Format options group
	fOptions := ui.NewGroup("Format options")
	vbox.Append(fOptions, false)
	{
		fvBox := ui.NewVerticalBox()
		fvBox.SetPadded(true)
		// Check device for bad blocks
		badBlocksCheck := ui.NewCheckbox("Check device for bad blocks")
		fvBox.Append(badBlocksCheck, false)

		fOptions.SetChild(fvBox)
	}
	win.SetChild(vbox)
	win.OnClosing(func(window *ui.Window) bool {
		os.Exit(1)
		return true
	})
	win.SetMargined(true)

}

func main() {
	err := ui.Main(uiMain)
	if err != nil {
		fmt.Printf("Failed to load gui: %s", err.Error())
		return
	}
}
