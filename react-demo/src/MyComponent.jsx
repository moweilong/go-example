import React, { useState } from "react";

function MyComponent() {

    const [car, setCar] = useState({ Year: "2015", Make: "Tesla", Model: "Model S" });

    const handleYearChange = (e) => {
        setCar(c => ({ ...c, Year: e.target.value }));
    };

    const handleMakeChange = (e) => {
        setCar(c => ({ ...c, Make: e.target.value }));
    };

    const handleModelChange = (e) => {
        setCar(c => ({ ...c, Model: e.target.value }));
    };

    return (
        <div>
            <p>Your favorite car is: {car.Year} {car.Make} {car.Model}</p>

            <input type="number" onChange={handleYearChange} value={car.Year} /><br />
            <input type="text" onChange={handleMakeChange} value={car.Make} /><br />
            <input type="text" onChange={handleModelChange} value={car.Model} />
        </div>
    );
}

export default MyComponent;