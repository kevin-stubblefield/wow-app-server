const expanders = document.getElementsByClassName('expander');
for (let expander of expanders) {
	expander.addEventListener('click', () => {
		let breakdown = expander.nextElementSibling;
		breakdown.classList.toggle('hidden');
		expander.lastChild.textContent = expander.lastChild.textContent == '+' ? '-' : '+';
	});
}