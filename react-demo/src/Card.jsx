import profilePic from './assets/profile.png'
function Card() {
    return (
        <div className="card">
            <img className="card-image" src={profilePic} alt="profile picture" />
            <h2 className='card-title'>WeiLong Mo</h2>
            <p className='card-text'>This is a card</p>
        </div>
    )
}

export default Card;