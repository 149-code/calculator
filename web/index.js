display = "";

dAppend = (str) => {
	display += str;
	document.getElementById("display_text").innerHTML = display;
}

dSet = (str) => {
	display = str;
	document.getElementById("display_text").innerHTML = display;
}

document.addEventListener('keydown', (event) => {
	goLog(event.key);

	if (!isNaN(event.key)) {
		dAppend(event.key);
	}

	if (event.key == '+') {
		dAppend("+");
	}
	if (event.key == '-') {
		dAppend("-");
	}
	if (event.key == '*') {
		dAppend("*");
	}
	if (event.key == '/') {
		dAppend("/");
	}

	if (event.key == 'c' || event.key == 'C') {
		dSet("");
	}
	if (event.key == 'Enter') {
		goEquals(display);
	}
})
