stats = ['str', 'dex', 'con', 'int', 'wis', 'cha']

proficiency_bonus = 2   

function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).value = '+' + proficiency_bonus;
    } else {
         document.getElementById(bonusId).value = '+0';
    }
}

function updateStats() {
    for (i = 0; i < 6; i++) {
        bonusUpdate(stats[i] + '_st', stats[i] + '_bonus')
    }

    for (i = 0; i < 6; i++) {
        stat = stats[i]

        st_sum = parseInt(document.getElementById(stat + "_base").value) + parseInt(document.getElementById(stat + "_bonus").value);

        if (st_sum >= 0) {
            document.getElementById(stat + "_sum").value = "+" + st_sum;
        } else {
            document.getElementById(stat + "_sum").value = st_sum;
        }
    }

}

function readST(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);

            modifiers = [
                loadedStats.StatsModifiers.Strength,
                loadedStats.StatsModifiers.Dexterity,
                loadedStats.StatsModifiers.Constitution,
                loadedStats.StatsModifiers.Intelligence,
                loadedStats.StatsModifiers.Wisdom,
                loadedStats.StatsModifiers.Charisma
            ]
            
            st_proficiency = [
                loadedStats.STProficiency.Strength,
                loadedStats.STProficiency.Dexterity,
                loadedStats.STProficiency.Constitution,
                loadedStats.STProficiency.Intelligence,
                loadedStats.STProficiency.Wisdom,
                loadedStats.STProficiency.Charisma
            ]

            proficiency_bonus = loadedStats.ProficiencyBonus; // Changes the global value

            for (i = 0; i < 6; i++) {
                if (modifiers[i] < 0) {
                    document.getElementById(stats[i] + "_base").value = modifiers[i];
                } else {
                    document.getElementById(stats[i] + "_base").value = '+' + modifiers[i];
                }

                document.getElementById(stats[i] + "_st").checked = st_proficiency[i];
            }

            updateStats();
        }
    }

    const reply = {Identifier: identifier}

    httpRequest.open("POST", "read-sheet"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
}

function writeST(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            alert("Saved saving throws to file");
        }
    }

    const reply = {
        Identifier: identifier,

        STProficiency: {
            Strength:       document.getElementById("str_st").checked,
            Dexterity:      document.getElementById("dex_st").checked,
            Constitution:   document.getElementById("con_st").checked,
            Intelligence:   document.getElementById("int_st").checked,
            Wisdom:         document.getElementById("wis_st").checked,
            Charisma:       document.getElementById("cha_st").checked
        }
    }

    httpRequest.open("POST", "write-st"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
    console.log(JSON.stringify(reply))
}
