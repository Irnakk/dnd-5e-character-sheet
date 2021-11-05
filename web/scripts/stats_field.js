function loadStats() {
    // Perform an AJAX-request to the server in order for it to roll stats;
    // then receive the rolled stats and write them into corresponding fields
    const httpRequest = new XMLHttpRequest(); // No idea why "const" here; could be also "var"

    httpRequest.onreadystatechange = function () {
        // telling the request object what to do once the server replies
        // processing the server response here

        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const rollResult = JSON.parse(this.responseText);
            document.getElementById("str_base").value = rollResult.Strength;
            document.getElementById("dex_base").value = rollResult.Dexterity;
            document.getElementById("con_base").value = rollResult.Constitution;
            document.getElementById("int_base").value = rollResult.Intelligence;
            document.getElementById("wis_base").value = rollResult.Wisdom;
            document.getElementById("cha_base").value = rollResult.Charisma;

            updateStats();
        }
    }

    httpRequest.open("GET", "roll-stats");
    httpRequest.send();
}

function updateStats() {
    const strSum = parseInt(document.getElementById("str_base").value) + parseInt(document.getElementById("str_bonus").value);
    const dexSum = parseInt(document.getElementById("dex_base").value) + parseInt(document.getElementById("dex_bonus").value);
    const conSum = parseInt(document.getElementById("con_base").value) + parseInt(document.getElementById("con_bonus").value);
    const intSum = parseInt(document.getElementById("int_base").value) + parseInt(document.getElementById("int_bonus").value);
    const wisSum = parseInt(document.getElementById("wis_base").value) + parseInt(document.getElementById("wis_bonus").value);
    const chaSum = parseInt(document.getElementById("cha_base").value) + parseInt(document.getElementById("cha_bonus").value);
    
    document.getElementById("str_sum").value = strSum;
    document.getElementById("dex_sum").value = dexSum;
    document.getElementById("con_sum").value = conSum;
    document.getElementById("int_sum").value = intSum;
    document.getElementById("wis_sum").value = wisSum;
    document.getElementById("cha_sum").value = chaSum;

    countModifiers();
}

function statInc(stat) {
    const current_bonus = parseInt(document.getElementById(stat).value);

    if (current_bonus < 99)
    {
        document.getElementById(stat).value = current_bonus + 1;
    }

    updateStats();
}

function statDec(stat) {
    const current_bonus = parseInt(document.getElementById(stat).value);

    if (current_bonus > -99)
    {
        document.getElementById(stat).value = current_bonus - 1;
    }

    updateStats();
}

function countModifiers(stat) {
    const strMod = Math.floor((parseInt(document.getElementById("str_sum").value) - 10) / 2);
    const dexMod = Math.floor((parseInt(document.getElementById("dex_sum").value) - 10) / 2);
    const conMod = Math.floor((parseInt(document.getElementById("con_sum").value) - 10) / 2);
    const intMod = Math.floor((parseInt(document.getElementById("int_sum").value) - 10) / 2);
    const wisMod = Math.floor((parseInt(document.getElementById("wis_sum").value) - 10) / 2);
    const chaMod = Math.floor((parseInt(document.getElementById("cha_sum").value) - 10) / 2);

    if (strMod < 0) {
        document.getElementById("str_mod").value = strMod;
    } else {
        document.getElementById("str_mod").value = '+' + strMod.toString();
    }
    if (dexMod < 0) {
        document.getElementById("dex_mod").value = dexMod;
    } else {
        document.getElementById("dex_mod").value = '+' + dexMod.toString();
    }
    if (conMod < 0) {
        document.getElementById("con_mod").value = conMod;
    } else {
        document.getElementById("con_mod").value = '+' + conMod.toString();
    }
    if (intMod < 0) {
        document.getElementById("int_mod").value = intMod;
    } else {
        document.getElementById("int_mod").value = '+' + intMod.toString();
    }
    if (wisMod < 0) {
        document.getElementById("wis_mod").value = wisMod;
    } else {
        document.getElementById("wis_mod").value = '+' + wisMod.toString();
    }
    if (chaMod < 0) {
        document.getElementById("cha_mod").value = chaMod;
    } else {
        document.getElementById("cha_mod").value = '+' + chaMod.toString();
    }
}

function readStats(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);
            document.getElementById("str_base").value = loadedStats.StatsBase.Strength;
            document.getElementById("dex_base").value = loadedStats.StatsBase.Dexterity;
            document.getElementById("con_base").value = loadedStats.StatsBase.Constitution;
            document.getElementById("int_base").value = loadedStats.StatsBase.Intelligence;
            document.getElementById("wis_base").value = loadedStats.StatsBase.Wisdom;
            document.getElementById("cha_base").value = loadedStats.StatsBase.Charisma;

            document.getElementById("str_bonus").value = loadedStats.StatsBonuses.Strength;
            document.getElementById("dex_bonus").value = loadedStats.StatsBonuses.Dexterity;
            document.getElementById("con_bonus").value = loadedStats.StatsBonuses.Constitution;
            document.getElementById("int_bonus").value = loadedStats.StatsBonuses.Intelligence;
            document.getElementById("wis_bonus").value = loadedStats.StatsBonuses.Wisdom;
            document.getElementById("cha_bonus").value = loadedStats.StatsBonuses.Charisma;

            updateStats();
        }
    }

    const reply = {Identifier: identifier}

    httpRequest.open("POST", "read-sheet"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
}

function writeStats(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            alert("Saved stats to file");
        }
    }

    const reply = {
        Identifier: identifier,
        StatsBase: {
            Strength:       parseInt(document.getElementById("str_base").value),
            Dexterity:      parseInt(document.getElementById("dex_base").value),
            Constitution:   parseInt(document.getElementById("con_base").value),
            Intelligence:   parseInt(document.getElementById("int_base").value),
            Wisdom:         parseInt(document.getElementById("wis_base").value),
            Charisma:       parseInt(document.getElementById("cha_base").value)
        },
        StatsBonuses: {
            Strength:       parseInt(document.getElementById("str_bonus").value),
            Dexterity:      parseInt(document.getElementById("dex_bonus").value),
            Constitution:   parseInt(document.getElementById("con_bonus").value),
            Intelligence:   parseInt(document.getElementById("int_bonus").value),
            Wisdom:         parseInt(document.getElementById("wis_bonus").value),
            Charisma:       parseInt(document.getElementById("cha_bonus").value)
        }
    }

    httpRequest.open("POST", "write-stats"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
    console.log(JSON.stringify(reply))
}
