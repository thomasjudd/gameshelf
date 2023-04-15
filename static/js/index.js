function gameFilter() {
	const filterText = document.getElementById("game-filter-text").value.toLowerCase();
	console.log(filterText);
	const cards = document.querySelectorAll(".card");

	let game = null
	let cardClass = ""
	let blockMatches = false;

	for(i = 0; i < cards.length; i++) {
		blockMatches = false;
		gameLinks = document.querySelectorAll("#" + cards[i].id + " .card-body .card-text .game-link");
		for(j = 0; j < gameLinks.length; j++) {
			gameText = gameLinks[j].innerHTML.toLowerCase();
			if(gameText.match(filterText)) {
				console.log(gameText + " matches");
				blockMatches = true
			}
		}
		if(blockMatches) {
			cards[i].removeAttribute("hidden");
		} else {
			cards[i].setAttribute("hidden", !blockMatches);
		}
	}
}
