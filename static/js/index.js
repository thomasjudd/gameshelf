function gameFilter() {
	const filterText = document.getElementById("game-filter-text").value.toLowerCase();
	console.log(filterText);
	const shelves = document.querySelectorAll(".shelf");

	let game = null
	let shelfClass = ""
	let blockMatches = false;

	for(i = 0; i < shelves.length; i++) {
		blockMatches = false;
		gameLinks = document.querySelectorAll("#" + shelves[i].id + " .shelf-body .shelf-text .game-link");
		for(j = 0; j < gameLinks.length; j++) {
			gameText = gameLinks[j].innerHTML.toLowerCase();
			if(gameText.match(filterText)) {
				console.log(gameText + " matches");
				blockMatches = true
			}
		}
		if(blockMatches) {
			shelves[i].removeAttribute("hidden");
		} else {
			shelves[i].setAttribute("hidden", !blockMatches);
		}
	}
}
