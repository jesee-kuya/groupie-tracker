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
    
                // Create a Set to track unique suggestions
                const seenSuggestions = new Set();
                let list = '<ul>';
                
                data.forEach(item => {
                    let suggestionText;
                    switch(item.matchType) {
                        case 'Member':
                            suggestionText = `${item.matchedItem} - ${item.matchType}`;
                            break;
                        case 'Location':
                            suggestionText = `${item.matchedItem} - ${item.matchType}`;
                            break;
                        case 'Artist':
                            suggestionText = `${item.name} - ${item.matchType}`;
                            break;
                        default:
                            suggestionText = `${item.matchedItem} - ${item.matchType}`;
                    }
    
                    // Only add if we haven't seen this exact suggestion text before
                    if (!seenSuggestions.has(suggestionText)) {
                        seenSuggestions.add(suggestionText);
                        list += `<li class="suggestion-item" onclick="selectSuggestion('${item.artistId}', '${item.matchType}', '${item.matchedItem}')">${suggestionText}</li>`;
                    }
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
        if (matchType === 'Artist') {
            // For artists, go directly to their page
            window.location.href = `/artist?id=${artistId}`;
        } else {
            // For locations, members, etc., search for all instances
            window.location.href = `/search?query=${encodeURIComponent(matchedItem)}`;
        }
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
