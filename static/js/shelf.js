function showAddModal() {
    const addModal = document.getElementById('add-modal');
    addModal.style.display = 'block';
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

var targetGame = null;

function showDeleteConfirmModal(buttonElement) {
  const toDelete = buttonElement.parentElement;
  const targetGame = {
		game_id: 0,
		name: toDelete.firstElementChild.innerHTML,
		shelf_id: 0
	};	
	console.log(targetGame);
  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'block';
}

function deleteGame() {
  console.log("deleteGame");
  console.log(targetGame);

  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(targetGame)
  };

  fetch('/game/delete', options)
  .then(response => {
		console.log(response);
	})
  .catch(error => {
    console.error(error);
  });

  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'none';
}

function cancelDelete() {
  console.log("cancelDelete");
  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'none';
}
