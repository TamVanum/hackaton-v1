import './App.css'
import Hero from './components/landing/Hero'
import About from './components/landing/About'
import Timeline from './components/landing/Timeline'


function App() {
  return (
    <div className="bg-gray-900">
      <Hero />
      <About />
      <Timeline />
    </div>
  )
}

export default App
