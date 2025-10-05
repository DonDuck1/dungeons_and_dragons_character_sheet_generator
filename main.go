package main

import (
	"dungeons_and_dragons_character_sheet_generator/domain"
	"dungeons_and_dragons_character_sheet_generator/infrastructure"
	"flag"
	"fmt"
	"log"
	"math"
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

	switch cmd {
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
		dndApiGateway := infrastructure.NewDndApiGateway("https://www.dnd5eapi.co")

		jsonBackgroundRepository, err := infrastructure.NewJsonBackgroundRepository("./data/backgrounds.json")
		if err != nil {
			log.Fatal(err)
		}

		jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
		if err != nil {
			log.Fatal(err)
		}

		jsonRaceRepository, err := infrastructure.NewJsonRaceRepository("./data/races.json")
		if err != nil {
			log.Fatal(err)
		}

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
		if !jsonCharacterRepository.IsCharacterNameUnique(*characterName) {
			fmt.Println("Another character with the same name already exists")
			os.Exit(2)
		}

		if *potentialRaceName == "" {
			fmt.Println("Race name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}
		race, err := jsonRaceRepository.GetByName(*potentialRaceName)
		if err != nil {
			log.Fatal(err)
		}

		if *potentialMainClassName == "" {
			fmt.Println("Class name is required")
			fmt.Println("")
			createCmd.Usage()
			os.Exit(2)
		}
		mainClassName, err := domain.ClassNameFromUntypedPotentialClassName(*potentialMainClassName)
		if err != nil {
			log.Fatal(err)
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

		proficiencyBonus := int(math.Ceil(float64(*level)/4)) + 1

		abilityScoreImprovements := race.GetChosenAbilityScoreImprovements()
		abilityScoreImprovementList := domain.NewAbilityScoreImprovementList(abilityScoreImprovements)
		abilityScoreValueList := domain.NewAbilityScoreValueList(*strengthValue, *dexterityValue, *constitutionValue, *intelligenceValue, *wisdomValue, *charismaValue)
		abilityScoreList := domain.NewAbilityScoreList(abilityScoreValueList, abilityScoreImprovementList)

		mainClass := CreateClassUsingApi(mainClassName, *level, proficiencyBonus, abilityScoreList, dndApiGateway)

		background, err := jsonBackgroundRepository.GetRandom()
		if err != nil {
			log.Fatal(err)
		}

		skillProficiencies := []domain.SkillProficiencyName{}
		skillProficiencies = append(skillProficiencies, mainClass.SkillProficiencies...)
		skillProficiencies = append(skillProficiencies, background.SkillProficiencies...)
		skillProficiencyList := domain.NewSkillProficiencyList(&abilityScoreList, skillProficiencies, proficiencyBonus)

		inventory := domain.NewEmptyInventory(race.NumberOfHandSlots)

		armorClass := inventory.GetArmorClass(abilityScoreList.Dexterity.Modifier)

		initiative := abilityScoreList.Dexterity.Modifier

		passivePerception := 10 + skillProficiencyList.Perception.Modifier

		character := domain.NewCharacter(
			*characterName,
			*race,
			mainClass,
			*background,
			proficiencyBonus,
			abilityScoreList,
			skillProficiencyList,
			armorClass,
			initiative,
			passivePerception,
			inventory,
		)

		jsonCharacterRepository.AddCharacter(character)
		err = jsonCharacterRepository.SaveCharacterList()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Character succesfully created!")

		os.Exit(0)
	case "view":
		jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
		if err != nil {
			log.Fatal(err)
		}

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

		character, err := jsonCharacterRepository.GetByName(*characterName)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(character)
		os.Exit(0)
	case "list":
		jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
		if err != nil {
			log.Fatal(err)
		}

		characters := jsonCharacterRepository.GetAll()
		if len(*characters) <= 0 {
			fmt.Println("There are no characters yet!")
			os.Exit(0)
		}

		fmt.Println("All characters:")
		for _, character := range *characters {
			fmt.Printf("%s, Lv%d %s, %s, %s\n", character.Name, character.MainClass.Level, character.MainClass.Name, character.Race.Name, character.Background.Name)
		}
		os.Exit(0)
	case "change-level":
		dndApiGateway := infrastructure.NewDndApiGateway("https://www.dnd5eapi.co")

		jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
		if err != nil {
			log.Fatal(err)
		}

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

		character, err := jsonCharacterRepository.GetByName(*characterName)
		if err != nil {
			log.Fatal(err)
		}

		proficiencyBonus := int(math.Ceil(float64(*level)/4)) + 1
		character.ProficiencyBonus = proficiencyBonus

		EditClassUsingApi(&character.MainClass, *level, proficiencyBonus, &character.AbilityScoreList, dndApiGateway)

		character.SkillProficiencyList.UpdateSkillProficiencies(proficiencyBonus)

		character.PassivePerception = 10 + character.SkillProficiencyList.Perception.Modifier

		err = jsonCharacterRepository.SaveCharacterList()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Character succesfully updated to level %d!\n", *level)
		os.Exit(0)
	case "delete":
		jsonCharacterRepository, err := infrastructure.NewJsonCharacterRepository("./data/characters.json")
		if err != nil {
			log.Fatal(err)
		}

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

		err = jsonCharacterRepository.DeleteCharacter(*characterName)
		if err != nil {
			log.Fatal(err)
		}

		err = jsonCharacterRepository.SaveCharacterList()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Character succesfully deleted!")
		os.Exit(0)
	case "equip":

	case "learn-spell":

	case "forget-spell":

	case "prepare-spell":

	default:
		usage()
		os.Exit(2)
	}
}
