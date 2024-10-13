import React, { useState } from 'react';

function ColorPicker() {
    const [color, setColor] = useState('#FFFFFF');

    const handleColorChange = (e) => {
        setColor(e.target.value);
    };

    return (
        <div className='color-picker-container'>
            <h2>Color Picker</h2>
            <div className='color-display' style={{ backgroundColor: color }}>
                <p>Selected Color: {color}</p>
            </div>
            <label >Select a Color:</label>
            <input type='color' value={color} onChange={handleColorChange} />
        </div>
    );
}

export default ColorPicker;