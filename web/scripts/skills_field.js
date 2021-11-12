stats = ['acr', 'anHand', 'arc', 'athl', 'dece', 
'his', 'ins', 'inti', 'invs', 
'med', 'nat', 'perce', 'perf', 
'pers', 'relig', 'SoH', 'stl', 'surv']

function bonusUpdate(checkBoxId, bonusId) {
    if (document.getElementById(checkBoxId).checked) {
        document.getElementById(bonusId).value = '+2';
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
