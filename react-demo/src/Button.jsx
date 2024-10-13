

function Button() {

    const handleClick = (e) => {
        e.target.textContent = "OUCH! You clicked me!"
    }

    return (
        <button onDoubleClick={(e) => handleClick(e)} >Click me</button>
    );
}

export default Button;