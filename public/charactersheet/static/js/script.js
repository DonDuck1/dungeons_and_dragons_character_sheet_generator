/**
 * Format number with sign
 *
 * @param {number} number
 */
function formatNumberWithSign(number) {
    return (number >= 0 ? "+" : "") + number;
};

/**
 * Sets up spellcasting info text based on the class spellcasting info of the character
 *
 * @param {character.Class.ClassSpellcastingInfo} classSpellcastingInfo
 *
 * @returns {string}
 */
function getSpellcastingInfoTextForNormalSpellcaster(classSpellcastingInfo) {
    spellcastingInfoText = "Spellcasting info:\n";
    spellcastingInfoText += `  Max known cantrips: ${classSpellcastingInfo.MaxKnownCantrips}\n`;
    if (classSpellcastingInfo.MaxKnownSpells != null) {
        spellcastingInfoText += `  Max known spells: ${classSpellcastingInfo.MaxKnownSpells}\n`;
    };
    if (classSpellcastingInfo.MaxPreparedSpells != null) {
        spellcastingInfoText += `  Max prepared spells: ${classSpellcastingInfo.MaxPreparedSpells}\n`;
    };

    spellcastingInfoText += "  Spell slots:\n";
    for (const [i, spellSlotLevelAmount] of classSpellcastingInfo.SpellSlotAmount.entries()) {
        spellcastingInfoText += `    Level ${i + 1}: ${spellSlotLevelAmount}\n`;
    };

    if (classSpellcastingInfo.SpellcastingAbility != null) {
        spellcastingInfoText += `  Spellcasting ability: ${classSpellcastingInfo.SpellcastingAbility.Name}\n`;
    };

    spellcastingInfoText += `  Spell save DC: ${classSpellcastingInfo.SpellSaveDC}\n`;
    spellcastingInfoText += `  Spell attack bonus: ${formatNumberWithSign(classSpellcastingInfo.SpellAttackBonus)}\n\n`;
}


/**
 * Sets up spellcasting info text based on the class warlock casting info of the character
 *
 * @param {character.Class.ClassWarlockCastingInfo} classWarlockCastingInfo
 *
 * @returns {string}
 */
function getSpellcastingInfoTextForWarlock(classWarlockCastingInfo) {
    spellcastingInfoText = "Warlock casting info:\n";
    spellcastingInfoText += `  Max known cantrips: ${classWarlockCastingInfo.MaxKnownCantrips}\n`;
    spellcastingInfoText += `  Max known spells: ${classWarlockCastingInfo.MaxKnownSpells}\n`;

    spellcastingInfoText += "  Spell slots:\n";
    spellcastingInfoText += `    Level ${classWarlockCastingInfo.SpellSlotLevel}: ${classWarlockCastingInfo.SpellSlotAmount}\n`;

    if (classWarlockCastingInfo.SpellcastingAbility != null) {
        spellcastingInfoText += `  Spellcasting ability: ${classWarlockCastingInfo.SpellcastingAbility.Name}\n`;
    };

    spellcastingInfoText += `  Spell save DC: ${classWarlockCastingInfo.SpellSaveDC}\n`;
    spellcastingInfoText += `  Spell attack bonus: ${formatNumberWithSign(classWarlockCastingInfo.SpellAttackBonus)}\n\n`;

    return spellcastingInfoText
}

/**
 * Sets up spell list text based on the normal spells of the character
 *
 * @param {character.Class.ClassSpellcastingInfo.SpellList.Spells} spells
 *
 * @returns {string}
 */
function getSpellListTextForNormalSpellcaster(spells) {
    spellsListText = "Spells:\n";

    for (const spell of spells) {
        if (spell.Prepared) {
            spellListText += `${spell.Name}\n`;
            spellListText += `  Level: ${spell.Level}\n`;
            spellListText += `  School: ${spell.School}\n`;
            spellListText += `  Range: ${spell.SpellRange}\n\n`;
        };
    };

    return spellListText
}

/**
 * Sets up spell list text based on the warlock spells of the character
 *
 * @param {character.Class.ClassWarlockCastingInfo.SpellList.Spells} spells
 *
 * @returns {string}
 */
function getSpellListTextForWarlock(spells) {
    spellListText = "Warlock spells:\n";

    for (const spell of spells) {
        if (spell.Prepared) {
            spellListText += `${spell.Name}\n`;
            spellListText += `  Level: ${spell.Level}\n`;
            spellListText += `  School: ${spell.School}\n`;
            spellListText += `  Range: ${spell.SpellRange}\n\n`;
        };
    };

    return spellListText
}

/**
 * Sets up equipment list text based on the inventory of the character
 *
 * @param {character.Inventory} inventory
 *
 * @returns {string}
 */
function getEquipmentListText(inventory) {
    let equipmentListText = "";

    if (inventory.WeaponSlots.MainHand != null) {
        equipmentListText += `Main hand: ${inventory.WeaponSlots.MainHand.Name}`;
        if (inventory.WeaponSlots.MainHand.TwoHanded) {
            equipmentListText += " (two-handed)\n";
        } else {
            equipmentListText += "\n";
        };
        equipmentListText += `  Category: ${inventory.WeaponSlots.MainHand.WeaponCategory}\n`;
        equipmentListText += `  Normal range: ${inventory.WeaponSlots.MainHand.NormalRange} feet\n\n`;
    };

    if (inventory.WeaponSlots.OffHand != null) {
        equipmentListText += `Off hand: ${inventory.WeaponSlots.OffHand.Name}`;
        if (inventory.WeaponSlots.OffHand.TwoHanded) {
            equipmentListText += " (two-handed)\n";
        } else {
            equipmentListText += "\n";
        };
        equipmentListText += `  Category: ${inventory.WeaponSlots.OffHand.WeaponCategory}\n`;
        equipmentListText += `  Normal range: ${inventory.WeaponSlots.OffHand.NormalRange} feet\n\n`;
    };

    if (inventory.Armor != null) {
        equipmentListText += `Armor: ${inventory.Armor.Name}\n\n`;
    };

    if (inventory.Shield != null) {
        equipmentListText += `Shield: ${inventory.Shield.Name}\n\n`;
    };

    return equipmentListText
}

document.addEventListener('DOMContentLoaded', () => {
    const select = document.getElementById('characterSelect');

    const characterNameField = document.getElementById('charName');
    const classLevelField = document.getElementById('classLevel');
    const backgroundField = document.getElementById('background');
    const raceField = document.getElementById('race');

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

    const proficiencyBonusField = document.getElementById('proficiencyBonus');

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

    const maxHpField = document.getElementById('maxHp');

    const equipmentListTextArea = document.getElementById('equipmentList');

    const featuresTextArea = document.getElementById('features');

    let characters = [];

    fetch('/api/characters')
        .then(response => response.json())
        .then(data => {
            characters = data;

            for (const character of characters) {
                const option = document.createElement('option');
                option.value = character.Name;

                if (character.Race.SubRace == null) {
                    option.textContent = `${character.Name}, Lv${character.Class.Level} ${character.Class.Name}, ${character.Race.Name}, ${character.Background.Name}`;
                } else {
                    option.textContent = `${character.Name}, Lv${character.Class.Level} ${character.Class.Name}, ${character.Race.SubRace.Name}, ${character.Background.Name}`;
                };
                
                select.appendChild(option);
            };
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

            if (character.Race.SubRace == null) {
                raceField.value = character.Race.Name;
            } else {
                raceField.value = character.Race.SubRace.Name;
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
            let spellListText = "";

            if (character.Class.ClassSpellcastingInfo != null) {
                spellcastingInfoText += getSpellcastingInfoTextForNormalSpellcaster(character.Class.ClassSpellcastingInfo)
                spellListText += getSpellListTextForNormalSpellcaster(character.Class.ClassSpellcastingInfo.SpellList.Spells)
            }

            if (character.Class.ClassWarlockCastingInfo != null) {
                spellcastingInfoText += getSpellcastingInfoTextForWarlock(character.Class.ClassWarlockCastingInfo)
                spellListText += getSpellListTextForWarlock(character.Class.ClassWarlockCastingInfo.SpellList.Spells)
            }

            otherProfsTextArea.value = spellcastingInfoText;

            featuresTextArea.value = spellListText;

            const equipmentListText = getEquipmentListText(character.Inventory);

            equipmentListTextArea.value = equipmentListText;
        };
    });
});