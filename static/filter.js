document.addEventListener('DOMContentLoaded', function() {
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

        function updateRangeTrack() {
            const minPercent = getPercent(minInput.value);
            const maxPercent = getPercent(maxInput.value);

            // Using colors that match your site's theme
            rangeTrack.style.background = `linear-gradient(
                to right,
                #e3e3e3 ${minPercent}%,
                #FF4136 ${minPercent}%,
                #FFF700 ${(minPercent + maxPercent) / 2}%,
                #FF4136 ${maxPercent}%,
                #e3e3e3 ${maxPercent}%
            )`;

            // Update value displays
            minValue.textContent = minInput.value;
            maxValue.textContent = maxInput.value;

            // Position value labels
            minValue.style.left = `${minPercent}%`;
            maxValue.style.left = `${maxPercent}%`;

            // Style the range thumbs to match the gradient
            minInput.style.setProperty('--thumb-color', '#FF4136');
            maxInput.style.setProperty('--thumb-color', '#FF4136');
        }

        minInput.addEventListener('input', function() {
            const minVal = parseInt(minInput.value);
            const maxVal = parseInt(maxInput.value);
            if (minVal > maxVal) {
                minInput.value = maxVal;
            }
            updateRangeTrack();
        });

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

    // Setup both range sliders
    const sliders = document.querySelectorAll('.double_range_slider');
    sliders.forEach(setupRangeSlider);
});