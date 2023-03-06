package main

import (
	"VRoidRipper/requests"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/http"
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
			headers := map[string]string{
				"X-Api-Version": "20211020",
			}
			resp, err := requests.GetWithHeaders("https://hub.vroid.com/api/character_models/"+item+"/optimized_preview", headers)
			if err != nil {
				fmt.Println(err)
			}
			err = os.WriteFile(item+".vrm", resp.Body(), 0755)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Finished!")
		} else {
			fmt.Println("Bad Input")
		}
	})

	content := container.NewVBox(
		widget.NewLabel("Enter VRoid character model URL:"),
		entry,
		button,
	)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
