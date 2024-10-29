const creationDateMin = document.getElementById("creationDateMin");
const creationDateMax = document.getElementById("creationDateMax");
const minDisplay = document.getElementById("creationDateMinValue");
const maxDisplay = document.getElementById("creationDateMaxValue");
const albumDateMin = document.getElementById("albumDateMin");
const albumDateMax = document.getElementById("albumDateMax");
const albuminDisplay = document.getElementById("albumDateMinValue");
const albummaxDisplay = document.getElementById("albumDateMaxValue");
const items = document.querySelectorAll(".item");

// Update displayed values
function updateDisplay() {
    minDisplay.textContent = creationDateMin.value;
    maxDisplay.textContent = creationDateMax.value;
    albuminDisplay.textContent = albumDateMin.value;
    albummaxDisplay.textContent = albumDateMax.value;
    filterItems();
}

// Filter items based on selected range
function filterItems() {
    const minYear = parseInt(creationDateMin.value);
    const maxYear = parseInt(creationDateMax.value);
    const albminYear = parseInt(albumDateMin.value);
    const albmaxYear = parseInt(albumDateMax.value);

    items.forEach(item => {
        const itemYear = parseInt(item.getAttribute("data-year"));
        
        // Check if itemYear is within both ranges
        if (itemYear >= minYear && itemYear <= maxYear && itemYear >= albminYear && itemYear <= albmaxYear) {
            item.style.display = "block";
        } else {
            item.style.display = "none";
        }
    });
}

// Event listeners for both sliders
creationDateMin.addEventListener("input", updateDisplay);
creationDateMax.addEventListener("input", updateDisplay);
albumDateMin.addEventListener("input", updateDisplay);
albumDateMax.addEventListener("input", updateDisplay);

// Initial filter on load
filterItems();