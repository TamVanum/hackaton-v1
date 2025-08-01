import { motion } from 'framer-motion'

export default function About() {
  return (
    <section className="min-h-screen bg-gray-800 flex items-center justify-center py-20">
      <div className="max-w-4xl mx-auto px-6 text-center">
        <motion.div
          initial={{ opacity: 0, y: 50 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          viewport={{ once: true }}
        >
          <h2 className="text-5xl font-bold text-white mb-8">
            About COMISTEICHON
          </h2>
          <p className="text-xl text-gray-300 leading-relaxed mb-8">
            Welcome to the ultimate hackathon where technology meets creativity. 
            Join us for an intensive coding marathon where innovation knows no bounds.
          </p>
          <div className="grid md:grid-cols-3 gap-8 mt-12">
            <div className="bg-gray-700 p-6 rounded-lg">
              <h3 className="text-2xl font-semibold text-green-400 mb-4">Innovation</h3>
              <p className="text-gray-300">
                Push the boundaries of what's possible with cutting-edge technology
              </p>
            </div>
            <div className="bg-gray-700 p-6 rounded-lg">
              <h3 className="text-2xl font-semibold text-green-400 mb-4">Collaboration</h3>
              <p className="text-gray-300">
                Work with talented developers from around the world
              </p>
            </div>
            <div className="bg-gray-700 p-6 rounded-lg">
              <h3 className="text-2xl font-semibold text-green-400 mb-4">Competition</h3>
              <p className="text-gray-300">
                Compete for amazing prizes and recognition in the tech community
              </p>
            </div>
          </div>
        </motion.div>
      </div>
    </section>
  )
}