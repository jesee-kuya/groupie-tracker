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
            
            rangeTrack.style.background = `linear-gradient(
                to right,
                #e3e3e3 ${minPercent}%,
                #4a90e2 ${minPercent}%,
                #4a90e2 ${maxPercent}%,
                #e3e3e3 ${maxPercent}%
            )`;
            
            minValue.textContent = minInput.value;
            maxValue.textContent = maxInput.value;
            
            // Position value labels
            minValue.style.left = `${minPercent}%`;
            maxValue.style.left = `${maxPercent}%`;
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