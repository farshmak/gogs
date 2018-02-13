function sortTable(table, col, reverse) {
    var tb = table.tBodies[0], // use `<tbody>` to ignore `<thead>` and `<tfoot>` rows
        tr = Array.prototype.slice.call(tb.rows, 0), // put rows into array
        i;
    reverse = -((+reverse) || -1);
    tr = tr.sort(function (a, b) { // sort rows
        return reverse // `-1 *` if want opposite order
        	  * compareStrings(a.cells[col].textContent.trim(),
                             b.cells[col].textContent.trim())            
        /*
            * (a.cells[col].textContent.trim() // using `.textContent.trim()` for test
                .localeCompare(b.cells[col].textContent.trim())
               );               
        */
    });
    for(i = 0; i < tr.length; ++i) tb.appendChild(tr[i]); // append each row in order
}

function compareStrings(str1, str2) {
	var rx = /([^\d]+|\d+)/ig;
	var str1split = str1.match( rx );
    var str2split = str2.match( rx );
    if (null === str1split)
        return 1;
    if (null === str2split)
        return -1;
	for(var i=0, l=Math.min(str1split.length, str2split.length); i < l; i++) {
		var s1 = str1split[i], 
		    s2 = str2split[i];
		if (s1 === s2) continue;
		if (isNaN(+s1) || isNaN(+s2))
			return s1 > s2 ? 1 : -1;
		else
			return +s1 - s2;		
	}
	return 0;
}

function makeSortable(table) {
    var th = table.tHead, i;
    th && (th = th.rows[0]) && (th = th.cells);
    if (th) i = th.length;
    else return; // if no `<thead>` then do nothing
    while (--i >= 0) (function (i) {
        var dir = 1;
        th[i].addEventListener('click', function () {sortTable(table, i, (dir = 1 - dir))});
    }(i));
}

function makeAllSortable(parent) {
    parent = parent || document.body;
    var t = parent.getElementsByTagName('table'), i = t.length;
    while (--i >= 0) makeSortable(t[i]);
}

window.onload = function () {makeAllSortable();};
