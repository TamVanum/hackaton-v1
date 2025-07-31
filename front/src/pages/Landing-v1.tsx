import './App.css'
import Hero from '../components/landing/Hero'
import About from '../components/landing/About'
import Prizes from '../components/landing/Prizes'
import Registration from '../components/landing/Registration'

function App() {
  return (
    <div className="min-h-screen">
      <nav className="fixed top-0 w-full z-50 bg-white/90 backdrop-blur-sm border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-6 py-4 flex justify-between items-center">
          <div className="text-2xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
            Hackathon PVC
          </div>
          <div className="hidden md:flex space-x-8">
            <a href="#about" className="text-gray-600 hover:text-gray-800 transition-colors">About</a>
            <a href="#prizes" className="text-gray-600 hover:text-gray-800 transition-colors">Prizes</a>
          </div>
          <a href="#register">
            <button className="mr-10 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-6 py-2 rounded-full hover:from-blue-700 hover:to-purple-700 transition-all">
              Register
            </button>
          </a>
        </div>
      </nav>

      <Hero />  
      <section id="about">
        <About />
      </section>
      <section id="prizes">
        <Prizes />
      </section>
      <section id="register">
        <Registration />
      </section>

      <footer className="px-6 py-12 bg-gray-900">
        <div className="max-w-6xl mx-auto">
          <div className="grid md:grid-cols-3 gap-8 mb-8">
            <div>
              <h3 className="text-xl font-bold text-white mb-4">Hackathon PVC</h3>
              <p className="text-gray-400">
                The ultimate coding challenge. Build, compete, and win amazing prizes.
              </p>
            </div>
            <div>
              <h4 className="text-lg font-semibold text-white mb-4">Quick Links</h4>
              <div className="space-y-2">
                <a href="#about" className="block text-gray-400 hover:text-white transition-colors">About</a>
                <a href="#prizes" className="block text-gray-400 hover:text-white transition-colors">Prizes</a>
                <a href="#register" className="block text-gray-400 hover:text-white transition-colors">Register</a>
              </div>
            </div>
            <div>
              <h4 className="text-lg font-semibold text-white mb-4">Contact</h4>
              <div className="space-y-2 text-gray-400">
                <p>info@hackathonpvc.com</p>
                <p>+1 (555) 123-4567</p>
                <div className="flex space-x-4 mt-4">
                  <a href="#" className="text-gray-400 hover:text-white transition-colors">Twitter</a>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors">Discord</a>
                  <a href="#" className="text-gray-400 hover:text-white transition-colors">LinkedIn</a>
                </div>
              </div>
            </div>
          </div>
          <div className="border-t border-gray-800 pt-8 text-center">
            <p className="text-gray-400">© 2024 Hackathon PVC. All rights reserved.</p>
          </div>
        </div>
      </footer>
    </div>
  )
}

export default App
