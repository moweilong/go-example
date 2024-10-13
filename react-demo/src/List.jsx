import PropTypes from 'prop-types';

function List(props) {
    // fruits.sort((a, b) => a.name.localeCompare(b.name)); // Sort the fruits alphabetically by name
    // fruits.sort((a, b) => b.name.localeCompare(a.name)); // Sort the fruits alphabetically by name (reverse)
    // fruits.sort((a, b) => a.calories - b.calories); // Sort the fruits by calories in ascending order
    // fruits.sort((a, b) => b.calories - a.calories); // Sort the fruits by calories in descending order

    // const lowCalFruits = fruits.filter((fruit) => fruit.calories < 100);
    // const highCalFruits = fruits.filter((fruit) => fruit.calories >= 100);

    const category = props.category;
    const itemList = props.items;

    const listItems = itemList.map((item) => <li key={item.id}>{item.name} : &nbsp;<b>{item.calories}</b></li>);

    return (
        <>
            <h2 className="list-category">{category}</h2>
            <ul className="list-items">{listItems}</ul>
            <ol className="list-items">{listItems}</ol>
        </>
    );
}

List.propTypes = {
    category: PropTypes.string,
    items: PropTypes.arrayOf(
        PropTypes.shape({
            id: PropTypes.number,
            name: PropTypes.string,
            calories: PropTypes.number,
        }),
    ),
}

List.defaultProps = {
    category: 'Category',
    items: [],
};


export default List;