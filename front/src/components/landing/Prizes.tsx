function Prizes() {
    return (
        <section className="py-20 bg-gradient-to-br from-purple-900 via-blue-900 to-indigo-900 text-white">
            <div className="container mx-auto px-6">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-bold mb-6">
                        Amazing Prizes Await
                    </h2>
                    <p className="text-xl text-blue-200 max-w-3xl mx-auto">
                        Compete for a total prize pool of $50,000 and incredible opportunities 
                        to kickstart your career or startup.
                    </p>
                </div>

                {/* Main Prizes */}
                <div className="grid md:grid-cols-3 gap-8 mb-16">
                    {/* 1st Place */}
                    <div className="relative group">
                        <div className="absolute inset-0 bg-gradient-to-r from-yellow-400 to-orange-500 rounded-3xl blur opacity-25 group-hover:opacity-40 transition duration-300"></div>
                        <div className="relative bg-white/10 backdrop-blur-sm border border-white/20 rounded-3xl p-8 text-center hover:transform hover:scale-105 transition-all duration-300">
                            <div className="text-6xl mb-4">🥇</div>
                            <h3 className="text-3xl font-bold mb-4 text-yellow-400">1st Place</h3>
                            <p className="text-4xl font-bold mb-4">$20,000</p>
                            <ul className="text-left space-y-2 text-gray-300">
                                <li>• Cash prize</li>
                                <li>• Mentorship program</li>
                                <li>• Investor introductions</li>
                                <li>• Tech equipment package</li>
                            </ul>
                        </div>
                    </div>

                    {/* 2nd Place */}
                    <div className="relative group">
                        <div className="absolute inset-0 bg-gradient-to-r from-gray-300 to-gray-500 rounded-3xl blur opacity-25 group-hover:opacity-40 transition duration-300"></div>
                        <div className="relative bg-white/10 backdrop-blur-sm border border-white/20 rounded-3xl p-8 text-center hover:transform hover:scale-105 transition-all duration-300">
                            <div className="text-6xl mb-4">🥈</div>
                            <h3 className="text-3xl font-bold mb-4 text-gray-300">2nd Place</h3>
                            <p className="text-4xl font-bold mb-4">$12,000</p>
                            <ul className="text-left space-y-2 text-gray-300">
                                <li>• Cash prize</li>
                                <li>• Cloud credits</li>
                                <li>• Development tools</li>
                                <li>• Conference tickets</li>
                            </ul>
                        </div>
                    </div>

                    {/* 3rd Place */}
                    <div className="relative group">
                        <div className="absolute inset-0 bg-gradient-to-r from-amber-600 to-yellow-600 rounded-3xl blur opacity-25 group-hover:opacity-40 transition duration-300"></div>
                        <div className="relative bg-white/10 backdrop-blur-sm border border-white/20 rounded-3xl p-8 text-center hover:transform hover:scale-105 transition-all duration-300">
                            <div className="text-6xl mb-4">🥉</div>
                            <h3 className="text-3xl font-bold mb-4 text-amber-400">3rd Place</h3>
                            <p className="text-4xl font-bold mb-4">$8,000</p>
                            <ul className="text-left space-y-2 text-gray-300">
                                <li>• Cash prize</li>
                                <li>• Software licenses</li>
                                <li>• Online courses</li>
                                <li>• Tech swag</li>
                            </ul>
                        </div>
                    </div>
                </div>

                {/* Special Awards */}
                <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
                    <div className="bg-white/5 backdrop-blur-sm border border-white/10 rounded-2xl p-6 text-center">
                        <div className="text-3xl mb-3">🎨</div>
                        <h4 className="text-lg font-bold mb-2 text-cyan-400">Best Design</h4>
                        <p className="text-sm text-gray-300">$2,000</p>
                    </div>
                    
                    <div className="bg-white/5 backdrop-blur-sm border border-white/10 rounded-2xl p-6 text-center">
                        <div className="text-3xl mb-3">🚀</div>
                        <h4 className="text-lg font-bold mb-2 text-green-400">Most Innovative</h4>
                        <p className="text-sm text-gray-300">$2,000</p>
                    </div>
                    
                    <div className="bg-white/5 backdrop-blur-sm border border-white/10 rounded-2xl p-6 text-center">
                        <div className="text-3xl mb-3">🌍</div>
                        <h4 className="text-lg font-bold mb-2 text-blue-400">Social Impact</h4>
                        <p className="text-sm text-gray-300">$2,000</p>
                    </div>
                    
                    <div className="bg-white/5 backdrop-blur-sm border border-white/10 rounded-2xl p-6 text-center">
                        <div className="text-3xl mb-3">👥</div>
                        <h4 className="text-lg font-bold mb-2 text-purple-400">People's Choice</h4>
                        <p className="text-sm text-gray-300">$2,000</p>
                    </div>
                </div>

                {/* Additional Benefits */}
                <div className="mt-16 text-center">
                    <h3 className="text-2xl font-bold mb-8">All Participants Receive</h3>
                    <div className="grid md:grid-cols-3 gap-6">
                        <div className="flex items-center justify-center space-x-3">
                            <div className="text-2xl">🎓</div>
                            <span className="text-lg">Certificate of Participation</span>
                        </div>
                        <div className="flex items-center justify-center space-x-3">
                            <div className="text-2xl">🍕</div>
                            <span className="text-lg">Free Meals & Snacks</span>
                        </div>
                        <div className="flex items-center justify-center space-x-3">
                            <div className="text-2xl">🎁</div>
                            <span className="text-lg">Exclusive Swag Kit</span>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Prizes 