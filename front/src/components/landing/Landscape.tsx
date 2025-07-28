function Landscape() {
    return (
          <div 
            className="relative w-full mb-6 overflow-hidden"
            style={{
                height: '450px',
                backgroundImage: 'url(/landscape/frutiger_background.webp)',
                backgroundSize: 'cover',
                backgroundPosition: 'center',
                backgroundRepeat: 'no-repeat'
            }}
          >
            <img 
              src="/landscape/running_man.gif" 
              alt="Running man animation" 
              className="absolute w-40 h-40 object-contain running-man-animation"
              style={{
                bottom: '10px',
                left: '10px'
              }}
            />  
          </div>
    )
}

export default Landscape;
