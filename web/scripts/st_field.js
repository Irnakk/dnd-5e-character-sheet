function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).value = '+2';
    } else {
         document.getElementById(bonusId).value = '+0';
    }
}

function updateStats() {
    bonusUpdate('str_st','str_bonus');
    bonusUpdate('dex_st','dex_bonus');
    bonusUpdate('con_st','con_bonus');
    bonusUpdate('int_st','int_bonus');
    bonusUpdate('wis_st','wis_bonus');
    bonusUpdate('cha_st','cha_bonus');

    const strSum = parseInt(document.getElementById("str_base").value) + parseInt(document.getElementById("str_bonus").value);
    const dexSum = parseInt(document.getElementById("dex_base").value) + parseInt(document.getElementById("dex_bonus").value);
    const conSum = parseInt(document.getElementById("con_base").value) + parseInt(document.getElementById("con_bonus").value);
    const intSum = parseInt(document.getElementById("int_base").value) + parseInt(document.getElementById("int_bonus").value);
    const wisSum = parseInt(document.getElementById("wis_base").value) + parseInt(document.getElementById("wis_bonus").value);
    const chaSum = parseInt(document.getElementById("cha_base").value) + parseInt(document.getElementById("cha_bonus").value);
    

    if (strSum >= 0) {
        document.getElementById("str_sum").value = "+" + strSum;
    } else {
        document.getElementById("str_sum").value = strSum;
    }
    
    if (dexSum >= 0) {
        document.getElementById("dex_sum").value = "+" + dexSum;
    } else {
        document.getElementById("dex_sum").value = dexSum;
    }

    if (conSum >= 0) {
        document.getElementById("con_sum").value = "+" + conSum;
    } else {
        document.getElementById("con_sum").value = conSum;
    }

    if (intSum >= 0) {
        document.getElementById("int_sum").value = "+" + intSum;
    } else {
        document.getElementById("int_sum").value = intSum;
    }
    
    if (wisSum >= 0) {
        document.getElementById("wis_sum").value = "+" + wisSum;
    } else {
        document.getElementById("wis_sum").value = wisSum;
    }

    if (chaSum >= 0) {
        document.getElementById("cha_sum").value = "+" + chaSum;
    } else {
        document.getElementById("cha_sum").value = chaSum;
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
