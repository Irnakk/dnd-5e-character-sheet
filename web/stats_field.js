function loadStats() {
    const httpRequest = new XMLHttpRequest(); // No idea why "const" here; could be also "var"

    httpRequest.onreadystatechange = function () {
        // telling the request object what to do once the server replies
        // processing the server response here

        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const rollResult = JSON.parse(this.responseText);
            document.getElementById("str_base").innerHTML = rollResult.Strength;
            document.getElementById("dex_base").innerHTML = rollResult.Dexterity;
            document.getElementById("con_base").innerHTML = rollResult.Constitution;
            document.getElementById("int_base").innerHTML = rollResult.Intelligence;
            document.getElementById("wis_base").innerHTML = rollResult.Wisdom;
            document.getElementById("cha_base").innerHTML = rollResult.Charisma;

            updateStats();
        }
    }

    httpRequest.open("GET", "roll-stats");
    httpRequest.send();
}

function updateStats() {
    const strSum = parseInt(document.getElementById("str_base").innerHTML) + parseInt(document.getElementById("str_bonus").innerHTML);
    const dexSum = parseInt(document.getElementById("dex_base").innerHTML) + parseInt(document.getElementById("dex_bonus").innerHTML);
    const conSum = parseInt(document.getElementById("con_base").innerHTML) + parseInt(document.getElementById("con_bonus").innerHTML);
    const intSum = parseInt(document.getElementById("int_base").innerHTML) + parseInt(document.getElementById("int_bonus").innerHTML);
    const wisSum = parseInt(document.getElementById("wis_base").innerHTML) + parseInt(document.getElementById("wis_bonus").innerHTML);
    const chaSum = parseInt(document.getElementById("cha_base").innerHTML) + parseInt(document.getElementById("cha_bonus").innerHTML);
    
    document.getElementById("str_sum").innerHTML = strSum;
    document.getElementById("dex_sum").innerHTML = dexSum;
    document.getElementById("con_sum").innerHTML = conSum;
    document.getElementById("int_sum").innerHTML = intSum;
    document.getElementById("wis_sum").innerHTML = wisSum;
    document.getElementById("cha_sum").innerHTML = chaSum;

    countModifiers();
}

function statInc(stat) {
    const current_bonus = parseInt(document.getElementById(stat).innerHTML);

    if (current_bonus <= 99)
    {
        document.getElementById(stat).innerHTML = current_bonus + 1;
    }

    updateStats();
}

function statDec(stat) {
    const current_bonus = parseInt(document.getElementById(stat).innerHTML);

    if (current_bonus > -9)
    {
        document.getElementById(stat).innerHTML = current_bonus - 1;
    }

    updateStats();
}

function countModifiers(stat) {
    const strMod = Math.floor((parseInt(document.getElementById("str_sum").innerHTML) - 10) / 2);
    const dexMod = Math.floor((parseInt(document.getElementById("dex_sum").innerHTML) - 10) / 2);
    const conMod = Math.floor((parseInt(document.getElementById("con_sum").innerHTML) - 10) / 2);
    const intMod = Math.floor((parseInt(document.getElementById("int_sum").innerHTML) - 10) / 2);
    const wisMod = Math.floor((parseInt(document.getElementById("wis_sum").innerHTML) - 10) / 2);
    const chaMod = Math.floor((parseInt(document.getElementById("cha_sum").innerHTML) - 10) / 2);

    if (strMod < 0) {
        document.getElementById("str_mod").innerHTML = strMod;
    } else {
        document.getElementById("str_mod").innerHTML = '+' + strMod.toString();
    }
    if (dexMod < 0) {
        document.getElementById("dex_mod").innerHTML = dexMod;
    } else {
        document.getElementById("dex_mod").innerHTML = '+' + dexMod.toString();
    }
    if (conMod < 0) {
        document.getElementById("con_mod").innerHTML = conMod;
    } else {
        document.getElementById("con_mod").innerHTML = '+' + conMod.toString();
    }
    if (intMod < 0) {
        document.getElementById("int_mod").innerHTML = intMod;
    } else {
        document.getElementById("int_mod").innerHTML = '+' + intMod.toString();
    }
    if (wisMod < 0) {
        document.getElementById("wis_mod").innerHTML = wisMod;
    } else {
        document.getElementById("wis_mod").innerHTML = '+' + wisMod.toString();
    }
    if (chaMod < 0) {
        document.getElementById("cha_mod").innerHTML = chaMod;
    } else {
        document.getElementById("cha_mod").innerHTML = '+' + chaMod.toString();
    }
}

function readStats(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);
            document.getElementById("str_base").innerHTML = loadedStats.StatsBase.Strength;
            document.getElementById("dex_base").innerHTML = loadedStats.StatsBase.Dexterity;
            document.getElementById("con_base").innerHTML = loadedStats.StatsBase.Constitution;
            document.getElementById("int_base").innerHTML = loadedStats.StatsBase.Intelligence;
            document.getElementById("wis_base").innerHTML = loadedStats.StatsBase.Wisdom;
            document.getElementById("cha_base").innerHTML = loadedStats.StatsBase.Charisma;

            document.getElementById("str_bonus").innerHTML = loadedStats.StatsBonuses.Strength;
            document.getElementById("dex_bonus").innerHTML = loadedStats.StatsBonuses.Dexterity;
            document.getElementById("con_bonus").innerHTML = loadedStats.StatsBonuses.Constitution;
            document.getElementById("int_bonus").innerHTML = loadedStats.StatsBonuses.Intelligence;
            document.getElementById("wis_bonus").innerHTML = loadedStats.StatsBonuses.Wisdom;
            document.getElementById("cha_bonus").innerHTML = loadedStats.StatsBonuses.Charisma;

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
            Strength:       parseInt(document.getElementById("str_base").innerHTML),
            Dexterity:      parseInt(document.getElementById("dex_base").innerHTML),
            Constitution:   parseInt(document.getElementById("con_base").innerHTML),
            Intelligence:   parseInt(document.getElementById("int_base").innerHTML),
            Wisdom:         parseInt(document.getElementById("wis_base").innerHTML),
            Charisma:       parseInt(document.getElementById("cha_base").innerHTML)
        },
        StatsBonuses: {
            Strength:       parseInt(document.getElementById("str_bonus").innerHTML),
            Dexterity:      parseInt(document.getElementById("dex_bonus").innerHTML),
            Constitution:   parseInt(document.getElementById("con_bonus").innerHTML),
            Intelligence:   parseInt(document.getElementById("int_bonus").innerHTML),
            Wisdom:         parseInt(document.getElementById("wis_bonus").innerHTML),
            Charisma:       parseInt(document.getElementById("cha_bonus").innerHTML)
        }
    }

    httpRequest.open("POST", "write-sheet"); // Does not work wit GET
    httpRequest.setRequestHeader("Content-Type", "application/json"); // It works without it, though
    httpRequest.send(JSON.stringify(reply));
    console.log(JSON.stringify(reply))
}
