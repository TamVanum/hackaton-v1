import '../../css/parallax.css'

function Parallax() {
    return (
        <div className="parallax-container">
            <div className="parallax-layer parallax-background">
                <img src="/parallax/background.png" alt="Background" />
            </div>
            <div className="parallax-layer parallax-clouds">
                <img src="/parallax/clouds.png" alt="Clouds" />
            </div>
            <div className="parallax-layer parallax-mountain">
                <img src="/parallax/mountain.png" alt="Mountain" />
            </div>
            <div className="parallax-layer parallax-pipes">
                <img src="/parallax/pipes.png" alt="Pipes" />
            </div>
        </div>
    )
}

export default Parallax; 