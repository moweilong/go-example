
function Food() {
    const Food1 = "Chicken";
    const Food2 = "Beef";

    return (
        <ul>
            <li>Apple</li>
            <li>{Food1}</li>
            <li>{Food2.toUpperCase()}</li>
        </ul>
    )
}

export default Food;