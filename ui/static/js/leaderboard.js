const bracketSelect = document.getElementById('bracket-select');
bracketSelect.addEventListener('change', (event) => {
	window.location.href = '/leaderboard/' + event.target.value;
});

const toTop = document.getElementById('to-top');
toTop.addEventListener('click', () => {
	window.scrollTo({
		top: 0, left: 0, behavior: 'smooth'
	});
});

window.addEventListener('scroll', () => {
	if (document.body.scrollTop > 300 || document.documentElement.scrollTop > 300) {
		toTop.classList.remove('hidden');
	} else {
		toTop.classList.add('hidden');
	}
});