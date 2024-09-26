// Function to listen for 'Alt + h' keydown event and fetch homepage content
export function fetchHomeOnKeydown() {
    document.addEventListener('keydown', (event) => {
        if (event.key === 'h' && event.altKey) {
            fetch('/home')
                .then(response => response.text())
                .then(data => {
                    document.body.innerHTML = data;
                })
                .catch(error => {
                    console.error('Error fetching homepage:', error);
                });
        }
    });
}
