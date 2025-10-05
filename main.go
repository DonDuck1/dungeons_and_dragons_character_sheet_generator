package main

import (
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Printf(`Usage:
  %s init
  %s create -name CHARACTER_NAME -race RACE -class CLASS -level N -str N -dex N -con N -int N -wis N -cha N
  %s view -name CHARACTER_NAME
  %s list
  %s change-level -name CHARACTER_NAME -level LEVEL
  %s delete -name CHARACTER_NAME
  %s equip -name CHARACTER_NAME -weapon WEAPON_NAME -slot SLOT
  %s equip -name CHARACTER_NAME -armor ARMOR_NAME
  %s equip -name CHARACTER_NAME -shield SHIELD_NAME
  %s learn-spell -name CHARACTER_NAME -spell SPELL_NAME
  %s forget-spell -name CHARACTER_NAME -spell SPELL_NAME
  %s prepare-spell -name CHARACTER_NAME -spell SPELL_NAME 
`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := os.Args[1]

	switch cmd {
	case "test":
		// csvEquipmentRepository, err := infrastructure.NewCsvEquipmentRepository("./data/5e-SRD-Equipment.csv")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		dndApiGateway := infrastructure.NewDndApiGateway("https://www.dnd5eapi.co")

		InitialiseBackgrounds(*dndApiGateway)
	case "init":
		csvEquipmentRepository, err := infrastructure.NewCsvEquipmentRepository("./data/5e-SRD-Equipment.csv")
		if err != nil {
			log.Fatal(err)
		}

		dndApiGateway := infrastructure.NewDndApiGateway("https://www.dnd5eapi.co")

		InitialiseArmorAndShields(*csvEquipmentRepository, *dndApiGateway)
		InitialiseBackgrounds(*dndApiGateway)
		InitialiseRaces(*dndApiGateway)
		InitialiseSpells(*dndApiGateway)
		InitialiseWeapons(*csvEquipmentRepository, *dndApiGateway)

		os.Exit(0)
	case "create":
		// You could use the Flag package like this
		// But feel free to do it differently!
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		name := createCmd.String("name", "", "character name (required)")
		err := createCmd.Parse(os.Args[2:])
		if *name == "" || err != nil {
			fmt.Println("name is required")
			createCmd.Usage()
			os.Exit(2)
		}

	case "view":

	case "list":

	case "change-level":

	case "delete":

	case "equip":

	case "learn-spell":

	case "forget-spell":

	case "prepare-spell":

	default:
		usage()
		os.Exit(2)
	}
}
