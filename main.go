package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"os"
)

func uiMain() {
	win := ui.NewWindow("Lufus v0.1.0", 0, 0, true)
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
	formatOptions := ui.NewGroup("Format options")
	vbox.Append(formatOptions, false)

	// File options vBox
	fvBox := ui.NewVerticalBox()
	fvBox.SetPadded(true)

	// Check device for bad blocks
	badBlocksCheckbox := ui.NewCheckbox("Check device for bad blocks")
	fvBox.Append(badBlocksCheckbox, false)

	// Quick format
	quickFormatCheckbox := ui.NewCheckbox("Quick format")
	quickFormatCheckbox.SetChecked(true)
	fvBox.Append(quickFormatCheckbox, false)

	// Create a bootable disk ...
	createBootableDiskByHBox := ui.NewHorizontalBox()

	// Create disk checkbox
	createDiskCheckbox := ui.NewCheckbox("Create a bootable disk using")
	createDiskCheckbox.SetChecked(true)
	createBootableDiskByHBox.Append(createDiskCheckbox, false)

	// By .. combo box
	isoOrDosComboBox := ui.NewCombobox()
	isoOrDosComboBox.Append("ISO image")
	isoOrDosComboBox.Append("FreeDOS")
	isoOrDosComboBox.SetSelected(0)
	createBootableDiskByHBox.Append(isoOrDosComboBox, false)

	// Select iso image button
	selectIsoButton := ui.NewButton("Select")
	createBootableDiskByHBox.Append(selectIsoButton, false)
	// Add this hbox into format group vbox
	fvBox.Append(createBootableDiskByHBox, false)

	// Create extended labels ... checkbox
	createExtendedLabelsCheckbox := ui.NewCheckbox("Create extended labels and icon files")
	createExtendedLabelsCheckbox.SetChecked(true)
	fvBox.Append(createExtendedLabelsCheckbox, false)

	// Set format options vBox as child of format box
	formatOptions.SetChild(fvBox)

	// Main progress bar - writing files to drive etc.
	mainProgressBar := ui.NewProgressBar()
	mainProgressBar.SetValue(100)
	vbox.Append(mainProgressBar, false)

	// Actual status
	statusLabel := ui.NewLabel("READY")
	vbox.Append(statusLabel, false)

	// Bottom button bar
	bottomButtonBar := ui.NewHorizontalBox()
	bottomButtonBar.SetPadded(true)

	aboutButton := ui.NewButton("About..")
	bottomButtonBar.Append(aboutButton, false)

	logButton := ui.NewButton("Log")
	bottomButtonBar.Append(logButton, false)

	// Empty label as spacer
	bottomButtonBar.Append(ui.NewLabel(""), true)

	startButton := ui.NewButton("Start")
	bottomButtonBar.Append(startButton, false)

	closeButton := ui.NewButton("Close")
	bottomButtonBar.Append(closeButton, false)

	vbox.Append(bottomButtonBar, false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	// Time and devices count bar
	bottomStatusBar := ui.NewHorizontalBox()
	bottomStatusBar.SetPadded(true)

	devicesCount := ui.NewLabel("1 device found")
	bottomStatusBar.Append(devicesCount, false)

	// Empty label as spacer
	bottomStatusBar.Append(ui.NewLabel(""), true)

	bottomStatusBar.Append(ui.NewVerticalSeparator(), false)
	timer := ui.NewLabel("00:00:00")
	bottomStatusBar.Append(timer, false)

	vbox.Append(bottomStatusBar, false)

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
