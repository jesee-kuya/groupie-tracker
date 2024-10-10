document.addEventListener('DOMContentLoaded', () => {
    const searchForm = document.querySelector('form[action="/search"]');
    const searchInput = searchForm.querySelector('input[name="query"]');
    const suggestionsContainer = document.createElement('div'); // Create a container for suggestions
    suggestionsContainer.classList.add('suggestions-container');
    searchForm.appendChild(suggestionsContainer);

    // Function to show suggestions
    function showSuggestions(val) {
        suggestionsContainer.innerHTML = ''; // Clear previous suggestions
        if (val === '') {
            return; // Don't fetch suggestions if the input is empty
        }

        fetch(`/suggestions?query=${encodeURIComponent(val)}`)
            .then(response => response.json())
            .then(data => {
                if (data.length === 0) {
                    suggestionsContainer.innerHTML = '<div>No suggestions found</div>';
                    return;
                }

                // Create a list for suggestions
                let list = '<ul>';
                data.forEach(item => {
                    list += `<li class="suggestion-item" onclick="selectSuggestion('${item.name}')">${item.name}</li>`;
                });
                list += '</ul>';
                suggestionsContainer.innerHTML = list;
            })
            .catch(err => {
                console.warn('Something went wrong.', err);
            });
    }

    // Event listener for input
    searchInput.addEventListener('input', function() {
        const query = this.value.trim();
        showSuggestions(query); // Call the function to show suggestions
    });

    // Suggestion click handler
    window.selectSuggestion = function(name) {
        searchInput.value = name; // Set the input value to the selected suggestion
        suggestionsContainer.innerHTML = ''; // Clear suggestions
    };

    // Existing form submission handler
    searchForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const query = searchInput.value.trim();
        if (query) {
            window.location.href = `/search?query=${encodeURIComponent(query)}`;
        } else {
            alert('Please enter a search query.');
        }
    });
});
