package main

import (
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"dungeons_and_dragons_character_sheet_generator/services"
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Printf(`Usage:
  go run . init
  go run . create -name "CHARACTER_NAME" -race "RACE" -class "CLASS" -level N -str N -dex N -con N -int N -wis N -cha N
  go run . view -name "CHARACTER_NAME"
  go run . list
  go run . change-level -name "CHARACTER_NAME" -level LEVEL
  go run . delete -name "CHARACTER_NAME"
  go run . equip -name "CHARACTER_NAME" -weapon "WEAPON_NAME" -slot SLOT
  go run . equip -name "CHARACTER_NAME" -armor "ARMOR_NAME"
  go run . equip -name "CHARACTER_NAME" -shield "SHIELD_NAME"
  go run . learn-spell -name "CHARACTER_NAME" -spell "SPELL_NAME"
  go run . forget-spell -name "CHARACTER_NAME" -spell "SPELL_NAME"
  go run . prepare-spell -name "CHARACTER_NAME" -spell "SPELL_NAME"

  Location: %s
`, os.Args[0])
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := os.Args[1]

	if cmd == "init" {
		csvEquipmentRepository, err := infrastructure.NewCsvEquipmentRepository("./data/5e-SRD-Equipment.csv")
		if err != nil {
			log.Fatal(err)
		}
		dndApiGateway := infrastructure.NewDndApiGateway("https://www.dnd5eapi.co")

		armorAndShieldService := services.NewArmorAndShieldService(csvEquipmentRepository, dndApiGateway)
		armorAndShieldService.InitialiseArmorAndShields()

		backgroundService := services.NewBackgroundService(dndApiGateway)
		backgroundService.InitialiseBackgrounds()

		classService := services.NewClassService(dndApiGateway)
		classService.InitialiseClasses()

		raceService := services.NewRaceService(dndApiGateway)
		raceService.InitialiseRaces()

		spellService := services.NewSpellService(dndApiGateway)
		spellService.InitialiseSpells()

		weaponService := services.NewWeaponService(csvEquipmentRepository, dndApiGateway)
		weaponService.InitialiseWeapons()

		os.Exit(0)
	}

	jsonBackgroundRepository, err := infrastructure.NewJsonBackgroundRepository("./data/backgrounds.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonClassRepository, err := infrastructure.NewJsonClassRepository("./data/classes.json")
	if err != nil {
		log.Fatal(err)
	}

	jsonRaceRepository, err := infrastructure.NewJsonRaceRepository("./data/races.json")
	if err != nil {
		log.Fatal(err)
	}

	characterService := services.NewCharacterService(
		jsonBackgroundRepository,
		jsonCharacterRepository,
		jsonClassRepository,
		jsonRaceRepository,
	)

	switch cmd {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)

		characterName := createCmd.String("name", "", "character name (required)")
		potentialRaceName := createCmd.String("race", "", "race name (required)")
		potentialMainClassName := createCmd.String("class", "", "main class name (required)")
		level := createCmd.Int("level", -999, "main class level")
		strengthValue := createCmd.Int("str", -999, "strength score value")
		dexterityValue := createCmd.Int("dex", -999, "dexterity score value")
		constitutionValue := createCmd.Int("con", -999, "constitution score value")
		intelligenceValue := createCmd.Int("int", -999, "intelligence score value")
		wisdomValue := createCmd.Int("wis", -999, "wisdom score value")
		charismaValue := createCmd.Int("cha", -999, "strength score value")

		err = createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *characterName == "" {
			fmt.Println("Character name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		if *potentialRaceName == "" {
			fmt.Println("Race name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		if *potentialMainClassName == "" {
			fmt.Println("Class name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		if *level == -999 {
			fmt.Println("No level was provided, level has been set to 1")
			*level = 1
		} else if *level < 1 {
			fmt.Printf("Provided level (%d) is too low, level has been set to 1 instead\n", *level)
			*level = 1
		} else if *level > 20 {
			fmt.Printf("Provided level (%d) is too high, level has been set to 20 instead\n", *level)
			*level = 20
		}

		if *strengthValue == -999 {
			fmt.Println("No strength score was provided, strength score has been set to 10")
			*strengthValue = 10
		} else if *strengthValue < 1 {
			fmt.Printf("Provided strength score (%d) is too low, strength score has been set to 1 instead\n", *strengthValue)
			*strengthValue = 1
		} else if *strengthValue > 20 {
			fmt.Printf("Provided strength score (%d) is too high, strength score has been set to 20 instead\n", *strengthValue)
			*strengthValue = 20
		}

		if *dexterityValue == -999 {
			fmt.Println("No dexterity score was provided, dexterity score has been set to 10")
			*dexterityValue = 10
		} else if *dexterityValue < 1 {
			fmt.Printf("Provided dexterity score (%d) is too low, dexterity score has been set to 1 instead\n", *dexterityValue)
			*dexterityValue = 1
		} else if *dexterityValue > 20 {
			fmt.Printf("Provided dexterity score (%d) is too high, dexterity score has been set to 20 instead\n", *dexterityValue)
			*dexterityValue = 20
		}

		if *constitutionValue == -999 {
			fmt.Println("No constitution score was provided, constitution score has been set to 10")
			*constitutionValue = 10
		} else if *constitutionValue < 1 {
			fmt.Printf("Provided constitution score (%d) is too low, constitution score has been set to 1 instead\n", *constitutionValue)
			*constitutionValue = 1
		} else if *constitutionValue > 20 {
			fmt.Printf("Provided constitution score (%d) is too high, constitution score has been set to 20 instead\n", *constitutionValue)
			*constitutionValue = 20
		}

		if *intelligenceValue == -999 {
			fmt.Println("No intelligence score was provided, intelligence score has been set to 10")
			*intelligenceValue = 10
		} else if *intelligenceValue < 1 {
			fmt.Printf("Provided intelligence score (%d) is too low, intelligence score has been set to 1 instead\n", *intelligenceValue)
			*intelligenceValue = 1
		} else if *intelligenceValue > 20 {
			fmt.Printf("Provided intelligence score (%d) is too high, intelligence score has been set to 20 instead\n", *intelligenceValue)
			*intelligenceValue = 20
		}

		if *wisdomValue == -999 {
			fmt.Println("No wisdom score was provided, wisdom score has been set to 10")
			*wisdomValue = 10
		} else if *wisdomValue < 1 {
			fmt.Printf("Provided wisdom score (%d) is too low, wisdom score has been set to 1 instead\n", *wisdomValue)
			*wisdomValue = 1
		} else if *wisdomValue > 20 {
			fmt.Printf("Provided wisdom score (%d) is too high, wisdom score has been set to 20 instead\n", *wisdomValue)
			*wisdomValue = 20
		}

		if *charismaValue == -999 {
			fmt.Println("No charisma score was provided, charisma score has been set to 10")
			*charismaValue = 10
		} else if *charismaValue < 1 {
			fmt.Printf("Provided charisma score (%d) is too low, charisma score has been set to 1 instead\n", *charismaValue)
			*charismaValue = 1
		} else if *charismaValue > 20 {
			fmt.Printf("Provided charisma score (%d) is too high, charisma score has been set to 20 instead\n", *charismaValue)
			*charismaValue = 20
		}

		characterService.CreateNewCharacter(
			*characterName,
			*potentialRaceName,
			*potentialMainClassName,
			*level,
			*strengthValue,
			*dexterityValue,
			*constitutionValue,
			*intelligenceValue,
			*wisdomValue,
			*charismaValue,
		)

		fmt.Println("Character succesfully created!")

		os.Exit(0)
	case "view":
		createCmd := flag.NewFlagSet("view", flag.ExitOnError)

		characterName := createCmd.String("name", "", "character name (required)")

		err = createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *characterName == "" {
			fmt.Println("Character name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		characterService.ViewCharacter(*characterName)
	case "list":
		characterService.ListCharacters()
	case "change-level":
		createCmd := flag.NewFlagSet("change-level", flag.ExitOnError)

		characterName := createCmd.String("name", "", "character name (required)")
		level := createCmd.Int("level", -999, "main class level")

		err = createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *characterName == "" {
			fmt.Println("Character name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		if *level == -999 {
			fmt.Println("No level was provided, level has been set to 1")
			*level = 1
		} else if *level < 1 {
			fmt.Printf("Provided level (%d) is too low, level has been set to 1 instead\n", *level)
			*level = 1
		} else if *level > 20 {
			fmt.Printf("Provided level (%d) is too high, level has been set to 20 instead\n", *level)
			*level = 20
		}

		characterService.ChangeLevelOfCharacter(*characterName, *level)
	case "delete":
		createCmd := flag.NewFlagSet("delete", flag.ExitOnError)

		characterName := createCmd.String("name", "", "character name (required)")

		err = createCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		if *characterName == "" {
			fmt.Println("Character name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}

		characterService.DeleteCharacter(*characterName)
	case "equip":

	case "learn-spell":

	case "forget-spell":

	case "prepare-spell":

	default:
		usage()
		os.Exit(2)
	}
}
