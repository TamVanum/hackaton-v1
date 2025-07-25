function About() {
    return (
        <section className="py-20 bg-white">
            <div className="container mx-auto px-6">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                        What is Hackathon PVC?
                    </h2>
                    <p className="text-xl text-gray-600 max-w-3xl mx-auto">
                        A 48-hour intensive coding marathon where developers, designers, and innovators 
                        come together to build groundbreaking solutions and compete for amazing prizes.
                    </p>
                </div>

                <div className="grid md:grid-cols-3 gap-8 mb-16">
                    <div className="text-center p-8 rounded-2xl bg-gradient-to-br from-blue-50 to-cyan-50 border border-blue-100">
                        <div className="w-16 h-16 bg-blue-500 rounded-full flex items-center justify-center mx-auto mb-6">
                            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z" />
                            </svg>
                        </div>
                        <h3 className="text-2xl font-bold text-gray-900 mb-4">Innovation</h3>
                        <p className="text-gray-600">
                            Transform ideas into reality. Build solutions that can change the world, 
                            one line of code at a time.
                        </p>
                    </div>

                    <div className="text-center p-8 rounded-2xl bg-gradient-to-br from-purple-50 to-pink-50 border border-purple-100">
                        <div className="w-16 h-16 bg-purple-500 rounded-full flex items-center justify-center mx-auto mb-6">
                            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                            </svg>
                        </div>
                        <h3 className="text-2xl font-bold text-gray-900 mb-4">Collaboration</h3>
                        <p className="text-gray-600">
                            Team up with brilliant minds. Learn from each other and create 
                            something extraordinary together.
                        </p>
                    </div>

                    <div className="text-center p-8 rounded-2xl bg-gradient-to-br from-green-50 to-emerald-50 border border-green-100">
                        <div className="w-16 h-16 bg-green-500 rounded-full flex items-center justify-center mx-auto mb-6">
                            <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                            </svg>
                        </div>
                        <h3 className="text-2xl font-bold text-gray-900 mb-4">Competition</h3>
                        <p className="text-gray-600">
                            Compete with the best developers. Push your limits and 
                            showcase your skills on a global stage.
                        </p>
                    </div>
                </div>

                <div className="bg-gradient-to-r from-gray-900 to-black rounded-3xl p-12 text-white text-center">
                    <h3 className="text-3xl md:text-4xl font-bold mb-6">
                        Who Can Participate?
                    </h3>
                    <div className="grid md:grid-cols-4 gap-6 text-center">
                        <div>
                            <p className="text-2xl font-bold text-cyan-400 mb-2">Students</p>
                            <p className="text-gray-300">All skill levels welcome</p>
                        </div>
                        <div>
                            <p className="text-2xl font-bold text-purple-400 mb-2">Professionals</p>
                            <p className="text-gray-300">Industry experts</p>
                        </div>
                        <div>
                            <p className="text-2xl font-bold text-green-400 mb-2">Designers</p>
                            <p className="text-gray-300">UI/UX specialists</p>
                        </div>
                        <div>
                            <p className="text-2xl font-bold text-yellow-400 mb-2">Entrepreneurs</p>
                            <p className="text-gray-300">Startup founders</p>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default About 