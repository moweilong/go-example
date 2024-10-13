import PropTypes from 'prop-types';

function UserGreeting(props) {
    // 写法一
    // if (props.isLoggedIn) {
    //     return <h2>Welcome, {props.username}</h2>;
    // }
    // return <h2>Please login</h2>;

    // 写法二
    // return (
    //     props.isLoggedIn ? <h2 className="welcome-message">Welcome, {props.username}</h2> : <h2 className="login-prompt">Please login</h2>
    // );

    // 写法三
    const welcomeMessage = <h2 className="welcome-message">Welcome, {props.username}</h2>;
    const loginPrompt = <h2 className="login-prompt">Please log in to continue</h2>;
    return props.isLoggedIn ? welcomeMessage : loginPrompt;

}

UserGreeting.propTypes = {
    isLoggedIn: PropTypes.bool,
    username: PropTypes.string,
}

UserGreeting.defaultProps = {
    isLoggedIn: false,
    username: 'Guest',
}

export default UserGreeting;
