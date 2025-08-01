import { motion } from 'framer-motion'

export default function Timeline() {
  const events = [
    {
      time: "October 21, 10:00 AM",
      title: "Registration & Welcome",
      description: "Check-in, team formation, and opening ceremony"
    },
    {
      time: "October 21, 12:00 PM",
      title: "Hacking Begins",
      description: "Teams start working on their projects"
    },
    {
      time: "October 21, 8:00 PM",
      title: "Dinner & Networking",
      description: "Food, drinks, and networking with other participants"
    },
    {
      time: "October 22, 8:00 AM",
      title: "Breakfast & Workshops",
      description: "Morning fuel and technical workshops"
    },
    {
      time: "October 22, 12:00 PM",
      title: "Final Submissions",
      description: "Projects must be submitted for judging"
    },
    {
      time: "October 22, 6:00 PM",
      title: "Presentations & Awards",
      description: "Team presentations and winner announcements"
    }
  ]

  return (
    <section className="min-h-screen bg-gray-900 py-20">
      <div className="max-w-4xl mx-auto px-6">
        <motion.div
          initial={{ opacity: 0, y: 50 }}
          whileInView={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.8 }}
          viewport={{ once: true }}
          className="text-center mb-16"
        >
          <h2 className="text-5xl font-bold text-white mb-4">
            Event Timeline
          </h2>
          <p className="text-xl text-gray-300">
            48 hours of non-stop coding and innovation
          </p>
        </motion.div>

        <div className="relative">
          {/* Timeline line */}
          <div className="absolute left-8 top-0 bottom-0 w-0.5 bg-green-400"></div>
          
          <div className="space-y-8">
            {events.map((event, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, x: -50 }}
                whileInView={{ opacity: 1, x: 0 }}
                transition={{ duration: 0.6, delay: index * 0.1 }}
                viewport={{ once: true }}
                className="relative flex items-start pl-20"
              >
                <div className="absolute left-6 w-4 h-4 bg-green-400 rounded-full -translate-x-1/2"></div>
                <div className="bg-gray-800 p-6 rounded-lg flex-1">
                  <div className="text-green-400 font-semibold text-sm mb-2">
                    {event.time}
                  </div>
                  <h3 className="text-xl font-bold text-white mb-2">
                    {event.title}
                  </h3>
                  <p className="text-gray-300">
                    {event.description}
                  </p>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </div>
    </section>
  )
}