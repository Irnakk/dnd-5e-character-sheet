skills = ['acr', 'anHand', 'arc', 'athl', 'dece', 
'his', 'ins', 'inti', 'invs', 
'med', 'nat', 'perce', 'perf', 
'pers', 'relig', 'SoH', 'stl', 'surv']

proficiency_bonus = 2

function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).value = '+' + proficiency_bonus;
    } else {
         document.getElementById(bonusId).value = '+0';
    }
}

function updateStats() {
    for (i = 0; i < 18; i++) {
        bonusUpdate(skills[i] + '_st', skills[i] + '_bonus')
    }

    for (i = 0; i < 18; i++) {
        skill = skills[i]

        skill_sum = parseInt(document.getElementById(skill + "_base").value) + parseInt(document.getElementById(skill + "_bonus").value);

        if (skill_sum >= 0) {
            document.getElementById(skill + "_sum").value = "+" + skill_sum;
        } else {
            document.getElementById(skill + "_sum").value = skill_sum;
        }
    }
}

function readSkills(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);

            skills_modifiers = [
                loadedStats.SkillsModifiers.Acrobatics,
                loadedStats.SkillsModifiers.AnimalHandling,
                loadedStats.SkillsModifiers.Arcana,
                loadedStats.SkillsModifiers.Athletics,
                loadedStats.SkillsModifiers.Deception,
                loadedStats.SkillsModifiers.History,
                loadedStats.SkillsModifiers.Insight,
                loadedStats.SkillsModifiers.Intimidation,
                loadedStats.SkillsModifiers.Investigation,
                loadedStats.SkillsModifiers.Medicine,
                loadedStats.SkillsModifiers.Nature,
                loadedStats.SkillsModifiers.Perception,
                loadedStats.SkillsModifiers.Performance,
                loadedStats.SkillsModifiers.Persuasion,
                loadedStats.SkillsModifiers.Religion,
                loadedStats.SkillsModifiers.SleightOfHand,
                loadedStats.SkillsModifiers.Stealth,
                loadedStats.SkillsModifiers.Survival
            ]
            
            skills_proficiency = [
                loadedStats.SkillsProficiency.Acrobatics,
                loadedStats.SkillsProficiency.AnimalHandling,
                loadedStats.SkillsProficiency.Arcana,
                loadedStats.SkillsProficiency.Athletics,
                loadedStats.SkillsProficiency.Deception,
                loadedStats.SkillsProficiency.History,
                loadedStats.SkillsProficiency.Insight,
                loadedStats.SkillsProficiency.Intimidation,
                loadedStats.SkillsProficiency.Investigation,
                loadedStats.SkillsProficiency.Medicine,
                loadedStats.SkillsProficiency.Nature,
                loadedStats.SkillsProficiency.Perception,
                loadedStats.SkillsProficiency.Performance,
                loadedStats.SkillsProficiency.Persuasion,
                loadedStats.SkillsProficiency.Religion,
                loadedStats.SkillsProficiency.SleightOfHand,
                loadedStats.SkillsProficiency.Stealth,
                loadedStats.SkillsProficiency.Survival
            ]

            proficiency_bonus = loadedStats.ProficiencyBonus; // Changes the global value

            for (i = 0; i < 18; i++) {
                if (skills_modifiers[i] < 0) {
                    document.getElementById(skills[i] + "_base").value = skills_modifiers[i];
                } else {
                    document.getElementById(skills[i] + "_base").value = '+' + skills_modifiers[i];
                }

                document.getElementById(skills[i] + "_st").checked = skills_proficiency[i];
            }

            updateStats();
        }
    }

    const reply = {Identifier: identifier}

    httpRequest.open("POST", "read-sheet"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
}


function writeSkills(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            alert("Saved skills to file");
        }
    }

    const reply = {
        Identifier: identifier,

        SkillsProficiency: {
            Acrobatics:     document.getElementById("acr_st").checked,
            AnimalHandling: document.getElementById("anHand_st").checked,
            Arcana:         document.getElementById("arc_st").checked,
            Athletics:      document.getElementById("athl_st").checked,
            Deception:      document.getElementById("dece_st").checked,
            History:        document.getElementById("his_st").checked,
            Insight:        document.getElementById("ins_st").checked,
            Intimidation:   document.getElementById("inti_st").checked,
            Investigation:  document.getElementById("invs_st").checked,
            Medicine:       document.getElementById("med_st").checked,
            Nature:         document.getElementById("nat_st").checked,
            Perception:     document.getElementById("perce_st").checked,
            Performance:    document.getElementById("perf_st").checked,
            Persuasion:     document.getElementById("pers_st").checked,
            Religion:       document.getElementById("relig_st").checked,
            SleightOfHand:  document.getElementById("SoH_st").checked,
            Stealth:        document.getElementById("stl_st").checked,
            Survival:       document.getElementById("surv_st").checked
        }
    }

    httpRequest.open("POST", "write-skills"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
    console.log(JSON.stringify(reply))
}
