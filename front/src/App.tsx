import './App.css'
import { motion, AnimatePresence } from 'framer-motion'
import { useEffect, useState } from 'react'
import BackgroundCode from './components/BackgroundCode'

function App() {
  const [messages, setMessages] = useState<{ msg: string }[]>([])
  const [showHack, setShowHack] = useState(false)
  const [showTextRelocated, setShowTextRelocated] = useState(false)
  const [showInfo, setShowInfo] = useState(false)
  const fallBackData = [{ "msg": "Bienvenido hijo de tu pta verga, a la vrga prro (╯°□°)╯︵ ┻━┻" }, { "msg": "Me relajo prro, me vine muy arriba, es que esto esta muy brigido ┬─┻ノ( º _ ºノ)" }, { "msg": "Ya con esto me retiro, como un campeon, valio todo la pena 😎" }, { "msg": "Se cierra el estadio señores 🏟️" }, { "msg": "A mimir 😴" }]

  useEffect(() => {
    fetch('https://api.comisteichon.cl/')
      .then(response => response.json())
      .then(data => {
        setMessages(data);
      })
      .catch(error => {
        setMessages(fallBackData)
        console.error('Error:', error);
      });
  }, [])

  // Show HACK text after all animations complete
  useEffect(() => {
    const timer = setTimeout(() => {
      setShowHack(true)
      setShowTextRelocated(true)
    }, 4000)

    return () => clearTimeout(timer)
  }, [])

  useEffect(() => {
    const timer = setTimeout(() => {
      setShowInfo(true)
    }, 4500)

    return () => clearTimeout(timer)
  }, [])

  return (
    <div className="relative isolate overflow-hidden bg-gray-900 h-screen">

      <AnimatePresence mode="wait">
        {showHack && (
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 2, ease: "easeInOut" }}
          >
            <BackgroundCode />
          </motion.div>
        )}
      </AnimatePresence>

      <div
        aria-hidden="true"
        className="absolute top-10 left-[calc(50%-4rem)] -z-10 transform-gpu blur-3xl sm:left-[calc(50%-18rem)] lg:top-[calc(50%-30rem)] lg:left-48 xl:left-[calc(50%-24rem)]"
      >
        <div
          style={{
            clipPath:
              'polygon(73.6% 51.7%, 91.7% 11.8%, 100% 46.4%, 97.4% 82.2%, 92.5% 84.9%, 75.7% 64%, 55.3% 47.5%, 46.5% 49.4%, 45% 62.9%, 50.3% 87.2%, 21.3% 64.1%, 0.1% 100%, 5.4% 51.1%, 21.4% 63.9%, 58.9% 0.2%, 73.6% 51.7%)',
          }}
          className="aspect-1108/632 w-[69.25rem] bg-linear-to-r from-[#13b16c] to-[#01f7ffd3] opacity-20"
        />
      </div>

      <div className="relative h-screen">
        <div className='flex items-end justify-center h-screen'>
          <img
            className="relative z-10 mb-0 w-full max-w-[300px] sm:max-w-[400px] md:max-w-[600px] lg:max-w-[800px] object-contain"
            src={"/hasbik2.png"}
            width={800}
            height={800}
            alt="Imagen de ejemplo"
          />
        </div>

        <div className="absolute inset-0 flex items-center justify-center z-0 mb-20 sm:mb-40 md:mb-60 lg:mb-95">
          <div className='flex flex-col px-4'>

            <motion.h1
              className="text-4xl sm:text-7xl md:text-8xl lg:text-8xl xl:text-9xl text-gray-200 font-extrabold leading-none text-center"
              initial={{ opacity: 0, y: 100 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{
                duration: 1.2,
                ease: "easeOut",
                delay: 0.3
              }}
            >
              COMISTEICHON
            </motion.h1>

            <AnimatePresence mode="wait">
              {!showTextRelocated ? (
                <motion.p
                  key="subtitle-text"
                  className="mt-2 text-gray-400 text-sm sm:text-base md:text-lg lg:text-xl text-right relative z-20"
                  initial={{ opacity: 0, x: -50 }}
                  animate={{ opacity: 1, x: 0 }}
                  exit={{
                    opacity: 0,
                    y: -10,
                    transition: { duration: 0.3 }
                  }}
                  transition={{
                    duration: 0.8,
                    ease: "easeOut",
                    delay: 1.2
                  }}
                >
                  El arte de comer comida
                </motion.p>
              ) : (
                <motion.div
                  key="empty-space"
                  initial={{ opacity: 0 }}
                  animate={{ opacity: 1 }}
                  className="mt-2"
                />
              )}
            </AnimatePresence>

            <motion.div
              className="mt-4 space-y-1"
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 2 }}
            >
              <AnimatePresence mode="wait">
                {!showHack ? (
                  <motion.div key="messages">
                    {messages.map((message, index) => (
                      <motion.div
                        key={`message-${index}`}
                        className="text-green-400 text-xs sm:text-sm font-mono opacity-60 text-right"
                        initial={{ opacity: 0, y: 10 }}
                        animate={{ opacity: 0.6, y: 0 }}
                        exit={{
                          opacity: 0,
                          y: -10,
                          transition: { duration: 0.3 }
                        }}
                        transition={{
                          delay: 2 + index * 0.1,
                          duration: 0.8
                        }}
                      >
                        {message.msg}
                      </motion.div>
                    ))}
                  </motion.div>
                ) : (
                  <motion.div
                    key="hack"
                    initial={{ opacity: 0, y: 10 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ delay: 0.1 }}
                  >
                    <h2 className="text-6xl md:text-8xl lg:text-8xl xl:text-9xl text-gray-200 font-extrabold text-right tracking-wider">
                      HACK
                    </h2>
                  </motion.div>
                )}
                <AnimatePresence mode="wait">
                  {showInfo && (
                    <motion.div
                      key="hack"
                      initial={{ opacity: 0, y: 10 }}
                      animate={{ opacity: 1, y: 0 }}
                      transition={{ delay: 0.1 }}
                    >
                      <h2 className="mr-1 text-lg md:text-xl lg:text-xl xl:text-xl text-green-400 font-extrabold text-right tracking-wider">
                      (╯°□°)╯︵ ┻━┻ 21-22.oct
                      </h2>
                    </motion.div>
                  )}
                </AnimatePresence>
              </AnimatePresence>
            </motion.div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
