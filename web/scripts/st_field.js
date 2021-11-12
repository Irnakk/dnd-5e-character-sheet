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
            str_mod = loadedStats.StatsModifiers.Strength;
            dex_mod = loadedStats.StatsModifiers.Dexterity;
            con_mod = loadedStats.StatsModifiers.Constitution;
            int_mod = loadedStats.StatsModifiers.Intelligence;
            wis_mod = loadedStats.StatsModifiers.Wisdom;
            cha_mod = loadedStats.StatsModifiers.Charisma;

            proficiency_bonus = loadedStats.ProficiencyBonus;

            if (str_mod < 0) {
                document.getElementById("str_base").value = str_mod;
            } else {
                document.getElementById("str_base").value = '+' + str_mod;
            }

            if (dex_mod < 0) {
                document.getElementById("dex_base").value = dex_mod;
            } else {
                document.getElementById("dex_base").value = '+' + dex_mod;
            }

            if (con_mod < 0) {
                document.getElementById("con_base").value = con_mod;
            } else {
                document.getElementById("con_base").value = '+' + con_mod;
            }

            if (int_mod < 0) {
                document.getElementById("int_base").value = int_mod;
            } else {
                document.getElementById("int_base").value = '+' + int_mod;
            }

            if (wis_mod < 0) {
                document.getElementById("wis_base").value = wis_mod;
            } else {
                document.getElementById("wis_base").value = '+' + wis_mod;
            }

            if (cha_mod < 0) {
                document.getElementById("cha_base").value = cha_mod;
            } else {
                document.getElementById("cha_base").value = '+' + cha_mod;
            }

            document.getElementById("str_st").checked = loadedStats.STProficiency.Strength;
            document.getElementById("dex_st").checked = loadedStats.STProficiency.Dexterity;
            document.getElementById("con_st").checked = loadedStats.STProficiency.Constitution;
            document.getElementById("int_st").checked = loadedStats.STProficiency.Intelligence;
            document.getElementById("wis_st").checked = loadedStats.STProficiency.Wisdom;
            document.getElementById("cha_st").checked = loadedStats.STProficiency.Charisma;

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
