function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).innerHTML = '+2';
    } else {
         document.getElementById(bonusId).innerHTML = '+0';
    }
}

function updateStats() {
    bonusUpdate('str_st','str_bonus');
    bonusUpdate('dex_st','dex_bonus');
    bonusUpdate('con_st','con_bonus');
    bonusUpdate('int_st','int_bonus');
    bonusUpdate('wis_st','wis_bonus');
    bonusUpdate('cha_st','cha_bonus');

    const strSum = parseInt(document.getElementById("str_base").innerHTML) + parseInt(document.getElementById("str_bonus").innerHTML);
    const dexSum = parseInt(document.getElementById("dex_base").innerHTML) + parseInt(document.getElementById("dex_bonus").innerHTML);
    const conSum = parseInt(document.getElementById("con_base").innerHTML) + parseInt(document.getElementById("con_bonus").innerHTML);
    const intSum = parseInt(document.getElementById("int_base").innerHTML) + parseInt(document.getElementById("int_bonus").innerHTML);
    const wisSum = parseInt(document.getElementById("wis_base").innerHTML) + parseInt(document.getElementById("wis_bonus").innerHTML);
    const chaSum = parseInt(document.getElementById("cha_base").innerHTML) + parseInt(document.getElementById("cha_bonus").innerHTML);
    

    if (strSum >= 0) {
        document.getElementById("str_sum").innerHTML = "+" + strSum;
    } else {
        document.getElementById("str_sum").innerHTML = strSum;
    }
    
    if (dexSum >= 0) {
        document.getElementById("dex_sum").innerHTML = "+" + dexSum;
    } else {
        document.getElementById("dex_sum").innerHTML = dexSum;
    }

    if (conSum >= 0) {
        document.getElementById("con_sum").innerHTML = "+" + conSum;
    } else {
        document.getElementById("con_sum").innerHTML = conSum;
    }

    if (intSum >= 0) {
        document.getElementById("int_sum").innerHTML = "+" + intSum;
    } else {
        document.getElementById("int_sum").innerHTML = intSum;
    }
    
    if (wisSum >= 0) {
        document.getElementById("wis_sum").innerHTML = "+" + wisSum;
    } else {
        document.getElementById("wis_sum").innerHTML = wisSum;
    }

    if (chaSum >= 0) {
        document.getElementById("cha_sum").innerHTML = "+" + chaSum;
    } else {
        document.getElementById("cha_sum").innerHTML = chaSum;
    }

}

function readST(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);
            str_mod = loadedStats.STModifiers.Strength;
            dex_mod = loadedStats.STModifiers.Dexterity;
            con_mod = loadedStats.STModifiers.Constitution;
            int_mod = loadedStats.STModifiers.Intelligence;
            wis_mod = loadedStats.STModifiers.Wisdom;
            cha_mod = loadedStats.STModifiers.Charisma;

            if (str_mod < 0) {
                document.getElementById("str_base").innerHTML = str_mod;
            } else {
                document.getElementById("str_base").innerHTML = '+' + str_mod;
            }

            if (dex_mod < 0) {
                document.getElementById("dex_base").innerHTML = dex_mod;
            } else {
                document.getElementById("dex_base").innerHTML = '+' + dex_mod;
            }

            if (con_mod < 0) {
                document.getElementById("con_base").innerHTML = con_mod;
            } else {
                document.getElementById("con_base").innerHTML = '+' + con_mod;
            }

            if (int_mod < 0) {
                document.getElementById("int_base").innerHTML = int_mod;
            } else {
                document.getElementById("int_base").innerHTML = '+' + int_mod;
            }

            if (wis_mod < 0) {
                document.getElementById("wis_base").innerHTML = wis_mod;
            } else {
                document.getElementById("wis_base").innerHTML = '+' + wis_mod;
            }

            if (cha_mod < 0) {
                document.getElementById("cha_base").innerHTML = cha_mod;
            } else {
                document.getElementById("cha_base").innerHTML = '+' + cha_mod;
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
