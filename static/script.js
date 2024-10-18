document.addEventListener('DOMContentLoaded', () => {
    const searchForm = document.querySelector('form[action="/search"]');
    const searchInput = searchForm.querySelector('input[name="query"]');
    
    // Create a wrapper div for the search input and suggestions
    const searchWrapper = document.createElement('div');
    searchWrapper.classList.add('search-wrapper');
    
    // Move the input into the wrapper
    searchInput.parentNode.insertBefore(searchWrapper, searchInput);
    searchWrapper.appendChild(searchInput);
    
    const suggestionsContainer = document.createElement('div');
    suggestionsContainer.classList.add('suggestions-container');
    
    // Append suggestions to the wrapper instead of the form
    searchWrapper.appendChild(suggestionsContainer);
    function showSuggestions(val) {
        suggestionsContainer.innerHTML = '';
        if (val === '') {
            return;
        }
        fetch(`/suggestions?query=${encodeURIComponent(val)}`)
            .then(response => response.json())
            .then(data => {
                if (data.length === 0) {
                    suggestionsContainer.innerHTML = '<div>No suggestions found</div>';
                    return;
                }
                let list = '<ul>';
                data.forEach(item => {
                    let displayText = item.matchType === 'Member' ? item.matchedItem : item.name;
                    let suggestionText = `${displayText} - ${item.matchType}`;
                    if (item.matchType === 'Location') {
                        suggestionText = `${item.matchedItem} (${item.artistName}) - ${item.matchType}`; 
                    }
                    list += `<li class="suggestion-item" onclick="selectSuggestion('${item.artistId}', '${item.matchType}', '${item.matchedItem}', '${item.artistName}')">${suggestionText}</li>`;
                });
                
                list += '</ul>';
                suggestionsContainer.innerHTML = list;
            })
            .catch(err => {
                console.warn('Something went wrong.', err);
            });
    }
    searchInput.addEventListener('input', function() {
        const query = this.value.trim();
        showSuggestions(query);
    });
    window.selectSuggestion = function(artistId, matchType, matchedItem) {
        // Navigate to the artist page
        window.location.href = `/artist?id=${artistId}`;
    };
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