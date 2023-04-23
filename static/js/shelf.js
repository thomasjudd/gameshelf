function showAddModal() {
    const addModal = document.getElementById('add-modal');
    addModal.style.display = 'block';
}

function addGame(buttonElement) {
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
		shelfContents = document.getElementsByClassName('card-body')[0];
		game = document.createElement('div');
		game.id = 'game-' + data['game_id'];
		game.classList.add('game');
		gameText = document.createElement('div');
		gameText.classList.add('game-text');
		gameText.textContent = data['name'];
		deleteButton = document.createElement('button');
		deleteButton.classList.add('delete')
		deleteButton.onclick = function() { showDeleteConfirmModal(this) };
		deleteButton.textContent = 'Delete';

		game.appendChild(gameText);
		game.appendChild(deleteButton);

		shelfContents.append(game);
  })
  .catch(error => {
    console.error(error);
  });


  const addModal = document.getElementById('add-modal');
  addModal.style.display = 'none';
}

var toDelete = null;

function showDeleteConfirmModal(buttonElement) {
  toDelete = buttonElement.parentElement;
	
  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'block';
}

function deleteGame() {
  console.log("deleteGame");
  console.log(toDelete);

	gameId = toDelete.id.split('-')[1];

  const targetGame = {
		game_id: gameId,
		name: toDelete.firstElementChild.innerHTML,
		shelf_id: 0
	};

  const options = {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json'
    }
  };

  fetch('/game/' + targetGame['game_id'], options)
  .catch(error => {
    console.error(error);
  });

	toDelete.remove();

  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'none';
}

function cancelDelete() {
  console.log("cancelDelete");
  deleteConfirmModal = document.getElementById('delete-confirm-modal');
  deleteConfirmModal.style.display = 'none';
}
