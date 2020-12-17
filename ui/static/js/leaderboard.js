const bracketSelect = document.getElementById('bracket-select');
bracketSelect.addEventListener("change", (event) => {
	window.location.href = '/leaderboard/' + event.target.value;
});