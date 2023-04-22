var targetGame = null;

function showAddModal() {
	const addModal = document.getElementById('add-modal');
	addModal.style.display = "block";
}

function addGame() {
	// make request to insert game
	const gameName = document.getElementById('add-game-input').value;
	const shelfId = window.location.pathname.split('/').pop();
	shelfIdInt = parseInt(shelfId);
	console.log(shelfIdInt);
	const payload = {
		name: gameName,
		shelf_id: shelfIdInt
	}

  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
  };

  fetch('/game/new', options)
	.then(response => response.json())
  .then(data => {
    console.log(data);
  })
  .catch(error => {
    console.error(error);
  });

	const addModal = document.getElementById('add-modal');
	addModal.style.display = 'none';
}


function deleteConfirmModal(elem) {
	targetGame = elem.parentElement
  gameName = targetGame.firstChild.innerHTML;

  deleteConfirmModal = document.getElementById('delete-confirm-modal');
	deleteConfirmModal.style.display.block;

}


function deleteGame() {
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: targetGame.innerHTML
  };

	fetch('/game/delete', options)
	.then(response =>  response.json())
	.then(data => {
		console.log(data);
	})
	.catch(error  => {
		console.error(error);
	});

	deleteConfirmModal = document.getElementById('delete-confirm-modal');
	deleteConfirmModal.style.display = 'none';
}

function cancelDelete() {
	targetGame = null;
	deleteConfirmModal = document.getElementById('delete-confirm-modal');
	deleteConfirmModal.style.display = 'none';
}
