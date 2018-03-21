package serverdisplay

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/PonGoLan/game/authoritative-server/instances"
	tm "github.com/buger/goterm"
)

const dashdashdash = "------------------------"

func PrintGeneralInformations(im *instances.InstancesManager) {
	const format = "%v\t%v\n"
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 1)
	fmt.Fprintf(w, format, "Number of instances", "Number of players")
	fmt.Fprintf(w, format, dashdashdash, dashdashdash)
	fmt.Fprintf(w, format, im.NumberOfInstances(), 0)
	w.Flush()
	fmt.Printf("\n\n")
}

func Print() {
	tm.Clear()

	instancesManager := instances.Get()
	print("\033[H\033[2J")

	PrintGeneralInformations(instancesManager)

	const format = "%v\t%v\t%v\n"

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 1)
	fmt.Fprintf(w, format, "ID", "Number of players", "Ticks")
	fmt.Fprintf(w, format, dashdashdash, dashdashdash, dashdashdash)
	for _, instance := range instancesManager.GetInstances() {
		fmt.Fprintf(w, format, instance.GetRoomName(), instance.GetNumberOfPlayersConnected(), instance.GetTicks())
	}
	w.Flush()
}
