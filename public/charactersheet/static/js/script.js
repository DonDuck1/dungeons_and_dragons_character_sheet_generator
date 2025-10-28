function formatNumberWithSign(number) {
    return (number >= 0 ? "+" : "") + number;
};

document.addEventListener('DOMContentLoaded', () => {
    const select = document.getElementById('characterSelect');

    const characterNameField = document.getElementById('charName');
    const classLevelField = document.getElementById('classLevel');
    const backgroundField = document.getElementById('background');
    const playerNameField = document.getElementById('playerName');
    const raceField = document.getElementById('race');
    const alignmentField = document.getElementById('alignment');
    const experiencePointsField = document.getElementById('experiencePoints');

    const strengthScoreField = document.getElementById('strengthScore');
    const strengthModField = document.getElementById('strengthMod');
    const dexterityScoreField = document.getElementById('dexterityScore');
    const dexterityModField = document.getElementById('dexterityMod');
    const constitutionScoreField = document.getElementById('constitutionScore');
    const constitutionModField = document.getElementById('constitutionMod');
    const wisdomScoreField = document.getElementById('wisdomScore');
    const wisdomModField = document.getElementById('wisdomMod');
    const intelligenceScoreField = document.getElementById('intelligenceScore');
    const intelligenceModField = document.getElementById('intelligenceMod');
    const charismaScoreField = document.getElementById('charismaScore');
    const charismaModField = document.getElementById('charismaMod');

    const inspirationCheckbox = document.getElementById('inspiration');
    const proficiencyBonusField = document.getElementById('proficiencyBonus');

    const strengthSaveField = document.getElementById('strengthSave');
    const strengthSaveProfCheckbox = document.getElementById('strengthSaveProf');
    const dexteritySaveField = document.getElementById('dexteritySave');
    const dexteritySaveProfCheckbox = document.getElementById('dexteritySaveProf');
    const constitutionSaveField = document.getElementById('constitutionSave');
    const constitutionSaveProfCheckbox = document.getElementById('constitutionSaveProf');
    const wisdomSaveField = document.getElementById('wisdomSave');
    const wisdomSaveProfCheckbox = document.getElementById('wisdomSaveProf');
    const intelligenceSaveField = document.getElementById('intelligenceSave');
    const intelligenceSaveProfCheckbox = document.getElementById('intelligenceSaveProf');
    const charismaSaveField = document.getElementById('charismaSave');
    const charismaSaveProf = document.getElementById('charismaSaveProf');

    const acrobaticsField = document.getElementById('acrobatics');
    const acrobaticsProfCheckbox = document.getElementById('acrobaticsProf');
    const animalHandlingField = document.getElementById('animalHandling');
    const animalHandlingProfCheckbox = document.getElementById('animalHandlingProf');
    const arcanaField = document.getElementById('arcana');
    const arcanaProfCheckbox = document.getElementById('arcanaProf');
    const athleticsField = document.getElementById('athletics');
    const athleticsProfCheckbox = document.getElementById('athleticsProf');
    const deceptionField = document.getElementById('deception');
    const deceptionProfCheckbox = document.getElementById('deceptionProf');
    const historyField = document.getElementById('history');
    const historyProfCheckbox = document.getElementById('historyProf');
    const insightField = document.getElementById('insight');
    const insightProfCheckbox = document.getElementById('insightProf');
    const intimidationField = document.getElementById('intimidation');
    const intimidationProfCheckbox = document.getElementById('intimidationProf');
    const investigationField = document.getElementById('investigation');
    const investigationProfCheckbox = document.getElementById('investigationProf');
    const medicineField = document.getElementById('medicine');
    const medicineProfCheckbox = document.getElementById('medicineProf');
    const natureField = document.getElementById('nature');
    const natureProfCheckbox = document.getElementById('natureProf');
    const perceptionField = document.getElementById('perception');
    const perceptionProfCheckbox = document.getElementById('perceptionProf');
    const performanceField = document.getElementById('performance');
    const performanceProfCheckbox = document.getElementById('performanceProf');
    const persuasionField = document.getElementById('persuasion');
    const persuasionProfCheckbox = document.getElementById('persuasionProf');
    const religionField = document.getElementById('religion');
    const religionProfCheckbox = document.getElementById('religionProf');
    const sleightOfHandField = document.getElementById('sleightOfHand');
    const sleightOfHandProfCheckbox = document.getElementById('sleightOfHandProf');
    const stealthField = document.getElementById('stealth');
    const stealthProfCheckbox = document.getElementById('stealthProf');
    const survivalField = document.getElementById('survival');
    const survivalProfCheckbox = document.getElementById('survivalProf');

    const passivePerceptionField = document.getElementById('passivePerception');

    const otherProfsTextArea = document.getElementById('otherProfs');

    const armorClassField = document.getElementById('ac');
    const initiativeField = document.getElementById('initiative');
    const speedField = document.getElementById('speed');

    const maxHpField = document.getElementById('maxHp');
    const currentHpField = document.getElementById('currentHp');
    const tempHpField = document.getElementById('tempHp');

    const totalHdField = document.getElementById('totalHd');
    const remainingHdField = document.getElementById('remainingHd');

    const deathSuccess1Checkbox = document.getElementById('deathSuccess1');
    const deathSuccess2Checkbox = document.getElementById('deathSuccess2');
    const deathSuccess3Checkbox = document.getElementById('deathSuccess3');
    const deathFail1Checkbox = document.getElementById('deathFail1');
    const deathFail2Checkbox = document.getElementById('deathFail2');
    const deathFail3Checkbox = document.getElementById('deathFail3');

    const atkName1Field = document.getElementById('atkName1');
    const atkBonus1Field = document.getElementById('atkBonus1');
    const atkDamage1Field = document.getElementById('atkDamage1');
    const atkName2Field = document.getElementById('atkName2');
    const atkBonus2Field = document.getElementById('atkBonus2');
    const atkDamage2Field = document.getElementById('atkDamage2');
    const atkName3Field = document.getElementById('atkName3');
    const atkBonus3Field = document.getElementById('atkBonus3');
    const atkDamage3Field = document.getElementById('atkDamage3');
    const otherAttacksAndSpellcastingTextArea = document.getElementById('otherAttacksAndSpellcasting');

    const copperPiecesField = document.getElementById('cp');
    const silverPiecesField = document.getElementById('sp');
    const electrumPiecesField = document.getElementById('ep');
    const goldPiecesField = document.getElementById('gp');
    const platinumPiecesField = document.getElementById('pp');
    const equipmentListTextArea = document.getElementById('equipmentList');

    const personalityTextArea = document.getElementById('personality');
    const idealsTextArea = document.getElementById('ideals');
    const bondsTextArea = document.getElementById('bonds');
    const flawsTextArea = document.getElementById('flaws');
    const featuresTextArea = document.getElementById('features');

    let characters = [];

    fetch('/api/characters')
        .then(response => response.json())
        .then(data => {
            characters = data;

            characters.forEach((character, i) => {
                const option = document.createElement('option');
                option.value = character.Name;
                option.textContent = `${character.Name}, Lv${character.Class.Level} ${character.Class.Name}, ${character.Race.Name}, ${character.Background.Name}`;
                select.appendChild(option);
            });
        });

    select.addEventListener('change', () => {
        const characterName = select.value;

        if (characterName === "") {
            characterNameField.value = "";
            classLevelField.value = "";
            backgroundField.value = "";
            raceField.value = "";

            strengthScoreField.value = "";
            strengthModField.value = "";

            dexterityScoreField.value = "";
            dexterityModField.value = "";

            constitutionScoreField.value = "";
            constitutionModField.value = "";

            wisdomScoreField.value = "";
            wisdomModField.value = "";

            intelligenceScoreField.value = "";
            intelligenceModField.value = "";

            charismaScoreField.value = "";
            charismaModField.value = "";

            proficiencyBonusField.value = "";

            acrobaticsField.value = "";
            acrobaticsProfCheckbox.checked = false;

            animalHandlingField.value = "";
            animalHandlingProfCheckbox.checked = false;

            arcanaField.value = "";
            arcanaProfCheckbox.checked = false;

            athleticsField.value = "";
            athleticsProfCheckbox.checked = false;

            deceptionField.value = "";
            deceptionProfCheckbox.checked = false;

            historyField.value = "";
            historyProfCheckbox.checked = false;

            insightField.value = "";
            insightProfCheckbox.checked = false;

            intimidationField.value = "";
            intimidationProfCheckbox.checked = false;

            investigationField.value = "";
            investigationProfCheckbox.checked = false;

            medicineField.value = "";
            medicineProfCheckbox.checked = false;

            natureField.value = "";
            natureProfCheckbox.checked = false;

            perceptionField.value = "";
            perceptionProfCheckbox.checked = false;

            performanceField.value = "";
            performanceProfCheckbox.checked = false;

            persuasionField.value = "";
            persuasionProfCheckbox.checked = false;

            religionField.value = "";
            religionProfCheckbox.checked = false;

            sleightOfHandField.value = "";
            sleightOfHandProfCheckbox.checked = false;

            stealthField.value = "";
            stealthProfCheckbox.checked = false;

            survivalField.value = "";
            survivalProfCheckbox.checked = false;

            passivePerceptionField.value = "";

            armorClassField.value = "";
            initiativeField.value = "";

            maxHpField.value = "";

            otherProfsTextArea.value = "";

            featuresTextArea.value = ""

            equipmentListTextArea.value = "";
        } else {
            const character = characters.find(character => character.Name === characterName);

            characterNameField.value = character.Name;
            classLevelField.value = `Lv${character.Class.Level} ${character.Class.Name}`;
            backgroundField.value = character.Background.Name;

            if (character.Race.SubRace != null) {
                raceField.value = character.Race.SubRace.Name;
            } else {
                raceField.value = character.Race.Name;
            };

            strengthScoreField.value = character.AbilityScoreList.Strength.FinalValue;
            strengthModField.value = formatNumberWithSign(character.AbilityScoreList.Strength.Modifier);

            dexterityScoreField.value = character.AbilityScoreList.Dexterity.FinalValue;
            dexterityModField.value = formatNumberWithSign(character.AbilityScoreList.Dexterity.Modifier);

            constitutionScoreField.value = character.AbilityScoreList.Constitution.FinalValue;
            constitutionModField.value = formatNumberWithSign(character.AbilityScoreList.Constitution.Modifier);

            wisdomScoreField.value = character.AbilityScoreList.Wisdom.FinalValue;
            wisdomModField.value = formatNumberWithSign(character.AbilityScoreList.Wisdom.Modifier);

            intelligenceScoreField.value = character.AbilityScoreList.Intelligence.FinalValue;
            intelligenceModField.value = formatNumberWithSign(character.AbilityScoreList.Intelligence.Modifier);

            charismaScoreField.value = character.AbilityScoreList.Charisma.FinalValue;
            charismaModField.value = formatNumberWithSign(character.AbilityScoreList.Charisma.Modifier);

            proficiencyBonusField.value = formatNumberWithSign(character.ProficiencyBonus)

            acrobaticsField.value = formatNumberWithSign(character.SkillProficiencyList.Acrobatics.Modifier);
            acrobaticsProfCheckbox.checked = character.SkillProficiencyList.Acrobatics.Proficient;

            animalHandlingField.value = formatNumberWithSign(character.SkillProficiencyList.AnimalHandling.Modifier);
            animalHandlingProfCheckbox.checked = character.SkillProficiencyList.AnimalHandling.Proficient;

            arcanaField.value = formatNumberWithSign(character.SkillProficiencyList.Arcana.Modifier);
            arcanaProfCheckbox.checked = character.SkillProficiencyList.Arcana.Proficient;

            athleticsField.value = formatNumberWithSign(character.SkillProficiencyList.Athletics.Modifier);
            athleticsProfCheckbox.checked = character.SkillProficiencyList.Athletics.Proficient;

            deceptionField.value = formatNumberWithSign(character.SkillProficiencyList.Deception.Modifier);
            deceptionProfCheckbox.checked = character.SkillProficiencyList.Deception.Proficient;

            historyField.value = formatNumberWithSign(character.SkillProficiencyList.History.Modifier);
            historyProfCheckbox.checked = character.SkillProficiencyList.History.Proficient;

            insightField.value = formatNumberWithSign(character.SkillProficiencyList.Insight.Modifier);
            insightProfCheckbox.checked = character.SkillProficiencyList.Insight.Proficient;

            intimidationField.value = formatNumberWithSign(character.SkillProficiencyList.Intimidation.Modifier);
            intimidationProfCheckbox.checked = character.SkillProficiencyList.Intimidation.Proficient;

            investigationField.value = formatNumberWithSign(character.SkillProficiencyList.Investigation.Modifier);
            investigationProfCheckbox.checked = character.SkillProficiencyList.Investigation.Proficient;

            medicineField.value = formatNumberWithSign(character.SkillProficiencyList.Medicine.Modifier);
            medicineProfCheckbox.checked = character.SkillProficiencyList.Medicine.Proficient;

            natureField.value = formatNumberWithSign(character.SkillProficiencyList.Nature.Modifier);
            natureProfCheckbox.checked = character.SkillProficiencyList.Nature.Proficient;

            perceptionField.value = formatNumberWithSign(character.SkillProficiencyList.Perception.Modifier);
            perceptionProfCheckbox.checked = character.SkillProficiencyList.Perception.Proficient;

            performanceField.value = formatNumberWithSign(character.SkillProficiencyList.Performance.Modifier);
            performanceProfCheckbox.checked = character.SkillProficiencyList.Performance.Proficient;

            persuasionField.value = formatNumberWithSign(character.SkillProficiencyList.Persuasion.Modifier);
            persuasionProfCheckbox.checked = character.SkillProficiencyList.Persuasion.Proficient;

            religionField.value = formatNumberWithSign(character.SkillProficiencyList.Religion.Modifier);
            religionProfCheckbox.checked = character.SkillProficiencyList.Religion.Proficient;

            sleightOfHandField.value = formatNumberWithSign(character.SkillProficiencyList.SleightOfHand.Modifier);
            sleightOfHandProfCheckbox.checked = character.SkillProficiencyList.SleightOfHand.Proficient;

            stealthField.value = formatNumberWithSign(character.SkillProficiencyList.Stealth.Modifier);
            stealthProfCheckbox.checked = character.SkillProficiencyList.Stealth.Proficient;

            survivalField.value = formatNumberWithSign(character.SkillProficiencyList.Survival.Modifier);
            survivalProfCheckbox.checked = character.SkillProficiencyList.Survival.Proficient;

            passivePerceptionField.value = character.PassivePerception;

            armorClassField.value = character.ArmorClass;
            initiativeField.value = formatNumberWithSign(character.Initiative);

            maxHpField.value = character.MaxHitPoints;

            let spellcastingInfoText = "";
            let spellsListText = "";

            if (character.Class.ClassSpellcastingInfo != null) {
                spellcastingInfoText += "Spellcasting info:\n";
                spellcastingInfoText += `  Max known cantrips: ${character.Class.ClassSpellcastingInfo.MaxKnownCantrips}\n`;
                if (character.Class.ClassSpellcastingInfo.MaxKnownSpells != null) {
                    spellcastingInfoText += `  Max known spells: ${character.Class.ClassSpellcastingInfo.MaxKnownSpells}\n`;
                };
                if (character.Class.ClassSpellcastingInfo.MaxPreparedSpells != null) {
                    spellcastingInfoText += `  Max prepared spells: ${character.Class.ClassSpellcastingInfo.MaxPreparedSpells}\n`;
                };

                spellcastingInfoText += "  Spell slots:\n";
                character.Class.ClassSpellcastingInfo.SpellSlotAmount.forEach((spellSlotLevelAmount, i) => {
                    spellcastingInfoText += `    Level ${i + 1}: ${spellSlotLevelAmount}\n`;
                });

                if (character.Class.ClassSpellcastingInfo.SpellcastingAbility != null) {
                    spellcastingInfoText += `  Spellcasting ability: ${character.Class.ClassSpellcastingInfo.SpellcastingAbility.Name}\n`;
                };

                spellcastingInfoText += `  Spell save DC: ${character.Class.ClassSpellcastingInfo.SpellSaveDC}\n`;
                spellcastingInfoText += `  Spell attack bonus: ${formatNumberWithSign(character.Class.ClassSpellcastingInfo.SpellAttackBonus)}\n\n`;
            
                spellsListText += "Spells:\n";
                character.Class.ClassSpellcastingInfo.SpellList.Spells.forEach((spell, i) => {
                    if (spell.Prepared) {
                        spellsListText += `${spell.Name}\n`;
                        spellsListText += `  Level: ${spell.Level}\n`;
                        spellsListText += `  School: ${spell.School}\n`;
                        spellsListText += `  Range: ${spell.SpellRange}\n\n`;
                    };
                });
            }

            if (character.Class.ClassWarlockCastingInfo != null) {
                spellcastingInfoText += "Warlock casting info:\n";
                spellcastingInfoText += `  Max known cantrips: ${character.Class.ClassWarlockCastingInfo.MaxKnownCantrips}\n`;
                spellcastingInfoText += `  Max known spells: ${character.Class.ClassWarlockCastingInfo.MaxKnownSpells}\n`;

                spellcastingInfoText += "  Spell slots:\n";
                spellcastingInfoText += `    Level ${character.Class.ClassWarlockCastingInfo.SpellSlotLevel}: ${character.Class.ClassWarlockCastingInfo.SpellSlotAmount}\n`;

                if (character.Class.ClassWarlockCastingInfo.SpellcastingAbility != null) {
                    spellcastingInfoText += `  Spellcasting ability: ${character.Class.ClassWarlockCastingInfo.SpellcastingAbility.Name}\n`;
                };

                spellcastingInfoText += `  Spell save DC: ${character.Class.ClassWarlockCastingInfo.SpellSaveDC}\n`;
                spellcastingInfoText += `  Spell attack bonus: ${formatNumberWithSign(character.Class.ClassWarlockCastingInfo.SpellAttackBonus)}\n\n`;
            
                spellsListText += "Warlock spells:\n";
                character.Class.ClassWarlockCastingInfo.SpellList.Spells.forEach((spell, i) => {
                    if (spell.Prepared) {
                        spellsListText += `${spell.Name}\n`;
                        spellsListText += `  Level: ${spell.Level}\n`;
                        spellsListText += `  School: ${spell.School}\n`;
                        spellsListText += `  Range: ${spell.SpellRange}\n\n`;
                    };
                });
            }

            otherProfsTextArea.value = spellcastingInfoText;

            featuresTextArea.value = spellsListText;

            let equipmentListText = "";

            if (character.Inventory.WeaponSlots.MainHand != null) {
                equipmentListText += `Main hand: ${character.Inventory.WeaponSlots.MainHand.Name}`;
                if (character.Inventory.WeaponSlots.MainHand.TwoHanded) {
                    equipmentListText += " (two-handed)\n";
                } else {
                    equipmentListText += "\n";
                };
                equipmentListText += `  Category: ${character.Inventory.WeaponSlots.MainHand.WeaponCategory}\n`;
                equipmentListText += `  Normal range: ${character.Inventory.WeaponSlots.MainHand.NormalRange} feet\n\n`;
            };

            if (character.Inventory.WeaponSlots.OffHand != null) {
                equipmentListText += `Off hand: ${character.Inventory.WeaponSlots.OffHand.Name}`;
                if (character.Inventory.WeaponSlots.OffHand.TwoHanded) {
                    equipmentListText += " (two-handed)\n";
                } else {
                    equipmentListText += "\n";
                };
                equipmentListText += `  Category: ${character.Inventory.WeaponSlots.OffHand.WeaponCategory}\n`;
                equipmentListText += `  Normal range: ${character.Inventory.WeaponSlots.OffHand.NormalRange} feet\n\n`;
            };

            if (character.Inventory.Armor != null) {
                equipmentListText += `Armor: ${character.Inventory.Armor.Name}\n\n`;
            };

            if (character.Inventory.Shield != null) {
                equipmentListText += `Shield: ${character.Inventory.Shield.Name}\n\n`;
            };

            equipmentListTextArea.value = equipmentListText;
        };
    });
});