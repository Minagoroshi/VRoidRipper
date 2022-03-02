package main

import (
	"VRoidRipper/requests"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
	"regexp"
)

func main() {

	a := app.NewWithID("VRoid Ripper")

	Icon, _ := fyne.LoadResourceFromURLString("https://i.pinimg.com/originals/69/0d/25/690d2555dc24257e693e43a79a2e1afc.png")
	mainWindow := a.NewWindow("VRoid Ripper")
	mainWindow.SetIcon(Icon)
	mainWindow.Resize(fyne.NewSize(400, 0))
	mainWindow.SetMaster()
	regex := regexp.MustCompile(`https:\/\/hub.vroid.com\/en\/characters\/[\w-]{19}\/models\/`)

	entry := widget.NewEntry()

	button := widget.NewButton("Rip VRM", func() {
		if entry.Text != "" || regex.MatchString(entry.Text) {
			item := regex.ReplaceAllString(entry.Text, "")
			fmt.Println("Downloading...")
			resp := requests.Get("https://hub.vroid.com/api/character_models/" + item + "/optimized_preview")
			err := os.WriteFile(item+".vrm", resp.Body(), 0755)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Finished!")
		} else {
			fmt.Println("Bad Input")
		}
	})

	content := widget.NewCard("VRoid Ripper", "by Top", container.NewVBox(entry, button))

	mainWindow.SetContent(container.NewVBox(content))
	mainWindow.ShowAndRun()

}
