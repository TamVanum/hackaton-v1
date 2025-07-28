
function Hero() {
    return (
        <section className="min-h-screen bg-gradient-to-br from-purple-900 via-blue-900 to-indigo-900 text-white relative overflow-hidden">
            
            {/* Hero content */}
            <div className="relative z-10 container mx-auto px-6 py-20 flex flex-col items-center justify-center min-h-screen text-center">
                <div className="mb-8">
                    <h1 className="text-6xl md:text-8xl font-bold mb-6 bg-gradient-to-r from-cyan-400 to-purple-400 bg-clip-text text-transparent">
                        Hackathon PVC
                    </h1>
                    <p className="text-2xl md:text-3xl font-light mb-4 text-blue-200">
                        Code • Create • Compete
                    </p>
                    <p className="text-lg md:text-xl text-gray-300 mb-8 max-w-2xl">
                        Join the most exciting 48-hour coding challenge. Build innovative solutions, 
                        meet amazing developers, and compete for incredible prizes.
                    </p>
                </div>

                <div className="mb-8 text-center">
                    <div className="inline-block bg-white/10 backdrop-blur-sm rounded-2xl p-6 border border-white/20">
                        <p className="text-sm uppercase tracking-wider text-cyan-300 mb-2">Event Date</p>
                        <p className="text-2xl md:text-3xl font-bold text-white">March 15-17, 2024</p>
                        <p className="text-lg text-gray-300 mt-2">48 Hours of Pure Innovation</p>
                    </div>
                </div>

                <div className="flex flex-col sm:flex-row gap-6 justify-center">
                    <a href="#register">
                        <button className="bg-gradient-to-r from-cyan-500 to-purple-600 hover:from-cyan-600 hover:to-purple-700 px-12 py-4 rounded-full text-xl font-semibold transform hover:scale-105 transition-all duration-300 shadow-2xl">
                            Register Now
                        </button>
                    </a>
                    <a href="#about">
                        <button className="border-2 border-white/30 hover:border-white/50 px-12 py-4 rounded-full text-xl font-semibold backdrop-blur-sm hover:bg-white/10 transition-all duration-300">
                            Learn More
                        </button>
                    </a>
                </div>

                {/* Stats */}
                <div className="grid grid-cols-3 gap-8 mt-16 text-center">
                    <div>
                        <p className="text-3xl md:text-4xl font-bold text-cyan-400">$50K</p>
                        <p className="text-sm text-gray-400">Prize Pool</p>
                    </div>
                    <div>
                        <p className="text-3xl md:text-4xl font-bold text-purple-400">500+</p>
                        <p className="text-sm text-gray-400">Participants</p>
                    </div>
                    <div>
                        <p className="text-3xl md:text-4xl font-bold text-blue-400">48h</p>
                        <p className="text-sm text-gray-400">To Code</p>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Hero 
