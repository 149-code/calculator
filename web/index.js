display = "";

dAppend = (str) => {
	display += str;
	document.getElementById("display_text").innerHTML = display;
}

dSet = (str) => {
	display = str;
	document.getElementById("display_text").innerHTML = display;
}
