document.addEventListener('DOMContentLoaded', function() {
    // Function to handle the range slider functionality
    function setupRangeSlider(container) {
        const rangeTrack = container.querySelector('.range_track');
        const minInput = container.querySelector('input.min');
        const maxInput = container.querySelector('input.max');
        const minValue = container.querySelector('.minvalue');
        const maxValue = container.querySelector('.maxvalue');
        
        // Calculate percentage for positioning
        function getPercent(value) {
            const min = parseInt(minInput.min);
            const max = parseInt(minInput.max);
            return ((value - min) / (max - min)) * 100;
        }

        // Update the range track and value displays
        function updateRangeTrack() {
            const minPercent = getPercent(minInput.value);
            const maxPercent = getPercent(maxInput.value);

            rangeTrack.style.background = `linear-gradient(
                to right,
                #e3e3e3 ${minPercent}%,
                #FF4136 ${minPercent}%,
                #FFF700 ${(minPercent + maxPercent) / 2}%,
                #FF4136 ${maxPercent}%,
                #e3e3e3 ${maxPercent}%
            )`;

            minValue.textContent = minInput.value;
            maxValue.textContent = maxInput.value;
            minValue.style.left = `${minPercent}%`;
            maxValue.style.left = `${maxPercent}%`;

            // Style the range thumbs
            minInput.style.setProperty('--thumb-color', '#FF4136');
            maxInput.style.setProperty('--thumb-color', '#FF4136');
        }

        // Ensure min input is not greater than max input
        minInput.addEventListener('input', function() {
            const minVal = parseInt(minInput.value);
            const maxVal = parseInt(maxInput.value);
            if (minVal > maxVal) {
                minInput.value = maxVal;
            }
            updateRangeTrack();
        });

        // Ensure max input is not less than min input
        maxInput.addEventListener('input', function() {
            const minVal = parseInt(minInput.value);
            const maxVal = parseInt(maxInput.value);
            if (maxVal < minVal) {
                maxInput.value = minVal;
            }
            updateRangeTrack();
        });

        // Initial update
        updateRangeTrack();
    }

    // Set up all range sliders
    const sliders = document.querySelectorAll('.double_range_slider');
    sliders.forEach(setupRangeSlider);

    // Handle the reset button functionality
    const resetBtn = document.querySelector('.reset-filter');
    resetBtn.addEventListener('click', function(event) {
        event.preventDefault();

        // Reset Range Sliders
        const minInputs = document.querySelectorAll('input.min');
        const maxInputs = document.querySelectorAll('input.max');
        const minValues = document.querySelectorAll('.minvalue');
        const maxValues = document.querySelectorAll('.maxvalue');
        minInputs.forEach(input => input.value = 1900);
        maxInputs.forEach(input => input.value = 2024);
        minValues.forEach(value => value.textContent = '1900');
        maxValues.forEach(value => value.textContent = '2024');
        
        // Update the range tracks to reflect reset values
        sliders.forEach(setupRangeSlider);

        // Reset Checkboxes
        const checkboxes = document.querySelectorAll('input[type="checkbox"]');
        checkboxes.forEach(checkbox => checkbox.checked = false);

        // Reset Text Inputs
        const textInputs = document.querySelectorAll('input[type="text"]');
        textInputs.forEach(input => input.value = '');
    });
});
