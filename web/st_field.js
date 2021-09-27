function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).innerHTML = '+2';
    } else {
         document.getElementById(bonusId).innerHTML = '+0';
    }
    updateStats()
}

function updateStats() {
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