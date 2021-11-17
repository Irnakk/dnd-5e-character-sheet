stats = ['str', 'dex', 'con', 'int', 'wis', 'cha']

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
    for (i = 0; i < 6; i++) {
        base_i = parseInt(document.getElementById(stats[i] + "_base").value);
        bonus_i = parseInt(document.getElementById(stats[i] + "_bonus").value);

        document.getElementById(stats[i] + "_sum").value = base_i + bonus_i;
    }

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
    for (i = 0; i < 6; i++) {
        modifier_i = Math.floor((parseInt(document.getElementById(stats[i] + "_sum").value) - 10) / 2);

        if (modifier_i < 0) {
            document.getElementById(stats[i] + "_mod").value = modifier_i;
        } else {
            document.getElementById(stats[i] + "_mod").value = '+' + modifier_i.toString();
        }
    }
}

function readStats(identifier) {
    const httpRequest = new XMLHttpRequest();

    httpRequest.onreadystatechange = function () {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            const loadedStats = JSON.parse(this.responseText);

            stats_base = [
                loadedStats.StatsBase.Strength,
                loadedStats.StatsBase.Dexterity,
                loadedStats.StatsBase.Constitution,
                loadedStats.StatsBase.Intelligence,
                loadedStats.StatsBase.Wisdom,
                loadedStats.StatsBase.Charisma
            ]

            stats_bonus = [
                loadedStats.StatsBonuses.Strength,
                loadedStats.StatsBonuses.Dexterity,
                loadedStats.StatsBonuses.Constitution,
                loadedStats.StatsBonuses.Intelligence,
                loadedStats.StatsBonuses.Wisdom,
                loadedStats.StatsBonuses.Charisma
            ]

            for (i = 0; i < 6; i++) {
                document.getElementById(stats[i] + "_base").value = stats_base[i];
                document.getElementById(stats[i] + "_bonus").value = stats_bonus[i];
            }

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
