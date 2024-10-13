import React, { useState, useEffect } from "react";

function MyComponent() {

    const [width, setWidth] = useState(window.innerWidth);
    const [height, setHeight] = useState(window.innerHeight);

    useEffect(() => {
        window.addEventListener("resize", handleResize)
        console.log("RESIZE EVENT ADD");

        return () => {
            window.removeEventListener("resize", handleResize);
            console.log("RESIZE EVENT REMOVE");
        };

    }, []);

    useEffect(() => {
        document.title = `Window: ${width} x ${height}`;
    }, [width, height]);

    const handleResize = () => {
        setWidth(window.innerWidth);
        setHeight(window.innerHeight);
    };

    return (
        <div>
            <h1>Window Width: {width}</h1>
            <h1>Window Height: {height}</h1>
        </div>
    );
}

export default MyComponent;