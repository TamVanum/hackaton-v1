import { useState } from 'react'

function Registration() {
    const [formData, setFormData] = useState({
        name: '',
        nickname: '',
        projectIdea: '',
        teammate: '',
        role: ''
    })

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault()
        // Handle form submission here
        console.log('Registration data:', formData)
        alert('Registration submitted! Check your email for confirmation.')
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
        const { name, value, type } = e.target
        setFormData(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? (e.target as HTMLInputElement).checked : value
        }))
    }

    return (
        <section className="py-20 bg-gradient-to-br from-cyan-50 to-blue-100">
            <div className="container mx-auto px-6">
                <div className="max-w-4xl mx-auto">
                    <div className="text-center mb-12">
                        <h2 className="text-4xl md:text-5xl font-bold text-gray-900 mb-6">
                            Ready to Join the Challenge?
                        </h2>
                        <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
                            Registration is free and spots are limited. Secure your place in 
                            the most exciting hackathon of the year!
                        </p>
                        
                        {/* Urgency indicators */}
                        <div className="flex justify-center space-x-8 mb-8">
                            <div className="text-center">
                                <p className="text-3xl font-bold text-red-500">47</p>
                                <p className="text-sm text-gray-600">Days Left</p>
                            </div>
                            <div className="text-center">
                                <p className="text-3xl font-bold text-orange-500">350</p>
                                <p className="text-sm text-gray-600">Spots Available</p>
                            </div>
                            <div className="text-center">
                                <p className="text-3xl font-bold text-green-500">150</p>
                                <p className="text-sm text-gray-600">Already Registered</p>
                            </div>
                        </div>
                    </div>

                    <div className="bg-white rounded-3xl shadow-2xl p-8 md:p-12">
                        <form onSubmit={handleSubmit} className="space-y-6">
                            <div className="grid md:grid-cols-2 gap-6">
                                <div>
                                    <label className="block text-sm font-semibold text-gray-700 mb-2">
                                        Full Name *
                                    </label>
                                    <input
                                        type="text"
                                        name="name"
                                        required
                                        value={formData.name}
                                        onChange={handleChange}
                                        className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                                        placeholder="Enter your full name"
                                    />
                                </div>
                                
                                <div>
                                    <label className="block text-sm font-semibold text-gray-700 mb-2">
                                        Nickname *
                                    </label>
                                    <input
                                        type="text"
                                        name="nickname"
                                        required
                                        value={formData.nickname}
                                        onChange={handleChange}
                                        className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                                        placeholder="Your nickname"
                                    />
                                </div>
                            </div>

                            <div>
                                <label className="block text-sm font-semibold text-gray-700 mb-2">
                                    Project Idea *
                                </label>
                                <textarea
                                    name="projectIdea"
                                    required
                                    value={formData.projectIdea}
                                    onChange={handleChange}
                                    rows={4}
                                    className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all resize-none"
                                    placeholder="Describe your project idea or what you'd like to build..."
                                />
                            </div>

                            <div className="grid md:grid-cols-2 gap-6">
                                <div>
                                    <label className="block text-sm font-semibold text-gray-700 mb-2">
                                        Who would you like to enter the hackathon with?
                                    </label>
                                    <input
                                        type="text"
                                        name="teammate"
                                        value={formData.teammate}
                                        onChange={handleChange}
                                        className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                                        placeholder="Person's name (optional)"
                                    />
                                </div>
                                
                                <div>
                                    <label className="block text-sm font-semibold text-gray-700 mb-2">
                                        Which role fits you better? *
                                    </label>
                                    <select
                                        name="role"
                                        required
                                        value={formData.role}
                                        onChange={handleChange}
                                        className="w-full px-4 py-3 border border-gray-300 rounded-xl focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
                                    >
                                        <option value="">Select your role</option>
                                        <option value="frontend">Frontend Developer</option>
                                        <option value="designer">Designer</option>
                                        <option value="backend">Backend Developer</option>
                                        <option value="fullstack">Fullstack Developer</option>
                                    </select>
                                </div>
                            </div>

                            <div className="pt-6">
                                <button
                                    type="submit"
                                    className="w-full bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-bold py-4 px-8 rounded-xl text-lg transform hover:scale-105 transition-all duration-300 shadow-lg"
                                >
                                    🚀 Register for Hackathon PVC
                                </button>
                                
                                <p className="text-center text-sm text-gray-500 mt-4">
                                    By registering, you agree to our terms and conditions. 
                                    Registration is completely free.
                                </p>
                            </div>
                        </form>
                    </div>

                    {/* Additional CTAs */}
                    <div className="mt-12 text-center">
                        <p className="text-lg text-gray-600 mb-6">
                            Have questions? Need help forming a team?
                        </p>
                        <div className="flex flex-col sm:flex-row gap-4 justify-center">
                            <button className="bg-white border-2 border-blue-600 text-blue-600 hover:bg-blue-600 hover:text-white px-8 py-3 rounded-xl font-semibold transition-all duration-300">
                                Join Discord Community
                            </button>
                            <button className="bg-white border-2 border-purple-600 text-purple-600 hover:bg-purple-600 hover:text-white px-8 py-3 rounded-xl font-semibold transition-all duration-300">
                                Contact Organizers
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}

export default Registration 