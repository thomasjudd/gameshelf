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

function deleteGame() {
	const deleteGame = document.getElementById('delete-game');
	const deleteModal = document.getElementById('delete-confirm-modal');

	payload = {
		"name": 
	}
  const options = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
  }; 

  fetch('/game/delete', options)
	.then(response => response.json())
  .then(data => {
    console.log(data);
  })
  .catch(error => {
    console.error(error);
  });


	cancelDelete.style.display = 'none';
}

function cancelDelete() {
	const deleteModal= document.getElementById('delete-confirm-modal');
	deleteModal.style.display = 'none';

}

function deleteConfirmModal() {
	const deleteModal = document.getElementById('delete-confirm-modal');
	deleteModal.style.display = 'block';
}
