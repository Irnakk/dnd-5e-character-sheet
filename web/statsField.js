function loadStats() {
    const httpRequest = new XMLHttpRequest(); // No idea why "const" here; could be also "var"

    httpRequest.onreadystatechange = function () {
        // telling the request object what to do once the server replies
        // processing the server response here
        const rollResult = JSON.parse(this.responseText);
        document.getElementById("str_base").innerHTML = rollResult.Strength;
        document.getElementById("dex_base").innerHTML = rollResult.Dexterity;
        document.getElementById("con_base").innerHTML = rollResult.Constitution;
        document.getElementById("int_base").innerHTML = rollResult.Intelligence;
        document.getElementById("wis_base").innerHTML = rollResult.Wisdom;
        document.getElementById("cha_base").innerHTML = rollResult.Charisma;

        const strSum = parseInt(document.getElementById("str_base").innerHTML) + parseInt(document.getElementById("str_bonus").innerHTML);
        const dexSum = parseInt(rollResult.Dexterity) + parseInt(document.getElementById("dex_bonus").innerHTML);
        const conSum = parseInt(rollResult.Constitution) + parseInt(document.getElementById("con_bonus").innerHTML);
        const intSum = parseInt(rollResult.Intelligence) + parseInt(document.getElementById("int_bonus").innerHTML);
        const wisSum = parseInt(rollResult.Wisdom) + parseInt(document.getElementById("wis_bonus").innerHTML);
        const chaSum = parseInt(rollResult.Charisma) + parseInt(document.getElementById("cha_bonus").innerHTML); 
        

        document.getElementById("str_sum").innerHTML = strSum;

        document.getElementById("dex_sum").innerHTML = dexSum;
        document.getElementById("con_sum").innerHTML = conSum;
        document.getElementById("int_sum").innerHTML = intSum;
        document.getElementById("wis_sum").innerHTML = wisSum;
        document.getElementById("cha_sum").innerHTML = chaSum;

    }

        httpRequest.open("GET", "roll-stats");
        httpRequest.send();
}