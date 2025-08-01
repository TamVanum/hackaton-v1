import { motion } from 'framer-motion'
import { Prism as SyntaxHighlighter } from 'react-syntax-highlighter'
import { vscDarkPlus } from 'react-syntax-highlighter/dist/esm/styles/prism'
import { useState, useRef, useEffect } from 'react'


function CodeEditor({ code }: { code: string }) {

  return (
    <div className="relative">
      <SyntaxHighlighter
        language="javascript"
        style={vscDarkPlus}
        customStyle={{
          background: 'transparent',
          padding: 0,
          margin: 0,
          fontSize: '14px',
          fontFamily: 'Fira Code, Consolas, Monaco, monospace',
          lineHeight: '1.5'
        }}
        showLineNumbers={true}
        lineNumberStyle={{
          color: '#6B7280',
          fontSize: '12px',
          minWidth: '3em',
          paddingRight: '1em'
        }}
      >
        {code || ' '}
      </SyntaxHighlighter>
    </div>
  )
}

function InteractiveTerminal({
  onCommandSuccess,
  failedAttempts,
  setFailedAttempts,
  setIsTyping
}: {
  onCommandSuccess: () => void
  failedAttempts: number
  setFailedAttempts: (attempts: number) => void
  setIsTyping: (typing: boolean) => void
}) {
  const [command, setCommand] = useState('')
  const [commandHistory, setCommandHistory] = useState<string[]>([
    '$ tree',
    'comisteichon',
    '└── hack.js',
    '',
    '1 directory, 1 file'
  ])
  const inputRef = useRef<HTMLInputElement>(null)
  const typingTimeoutRef = useRef<number | null>(null)

  const handleTyping = () => {
    setIsTyping(true)
    
    // Clear existing timeout
    if (typingTimeoutRef.current) {
      clearTimeout(typingTimeoutRef.current)
    }
    
    // Set typing back to false after 100ms
    typingTimeoutRef.current = setTimeout(() => {
      setIsTyping(false)
    }, 100)
  }

  const validCommands = [
    'code comisteichon/hack.js',
    'code comisteichon/hack',
    'code hack.js',
    'code hack',
    'vim comisteichon/hack.js',
    'vim comisteichon/hack',
    'vim hack.js',
    'vim hack',
    'cursor comisteichon/hack.js',
    'cursor comisteichon/hack',
    'cursor hack.js',
    'cursor hack',
    'nano comisteichon/hack.js',
    'nano comisteichon/hack',
    'nano hack.js',
    'nano hack',
  ]


  useEffect(() => {
    // Cleanup timeout on unmount
    return () => {
      if (typingTimeoutRef.current) {
        clearTimeout(typingTimeoutRef.current)
      }
    }
  }, [])

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    if (command.trim()) {
      const newHistory = [...commandHistory, `$ ${command}`]

      if (validCommands.includes(command.trim())) {
        newHistory.push('Opening comisteichon/hack.js...')
        setCommandHistory(newHistory)
        setCommand('')
        setFailedAttempts(0) // Reset attempts on success
        setTimeout(() => {
          onCommandSuccess()
        }, 500)
      } else {
        const newFailedAttempts = failedAttempts + 1
        setFailedAttempts(newFailedAttempts)

        newHistory.push(`Command not found: ${command}`)

        // Progressive hints based on number of failed attempts
        if (newFailedAttempts === 1) {
          newHistory.push('pista: intenta abrir el archivo hack.js')
        } else if (newFailedAttempts === 2) {
          newHistory.push('pista: intenta usando el comando code, vim, nano')
        } else if (newFailedAttempts >= 3) {
          newHistory.push('pista: el wn perkin, escribe esto code comisteichon/hack')
        }

        setCommandHistory(newHistory)
        setCommand('')
      }
    }
  }

  return (
    <div className="text-green-400 font-mono text-sm">
      <div className="mb-4">
        <div className="text-green-300 mb-2">comisteichon@hackathon:~$</div>
        <div className="text-gray-400 text-xs mb-4">
          Enter a command to open the hackathon file...
        </div>
      </div>

      {commandHistory.map((line, index) => (
        <div key={index} className="mb-1">
          {line}
        </div>
      ))}

      <form onSubmit={handleSubmit} className="flex items-center">
        <span className="text-green-300 mr-2">$</span>
        <input
          ref={inputRef}
          type="text"
          value={command}
          onChange={(e) => {
            setCommand(e.target.value)
            handleTyping()
          }}
          className="bg-transparent outline-none flex-1 text-green-400 caret-green-400"
          placeholder="type your command here..."
        />
      </form>
    </div>
  )
}

export default function About() {
  const [showCodeEditor, setShowCodeEditor] = useState(false)
  const [failedAttempts, setFailedAttempts] = useState(0)
  const [isTyping, setIsTyping] = useState(false)

  // Function to get the current image based on state
  const getCurrentImage = () => {
    if (isTyping) {
      return '/about/transition.png'
    }
    
    if (showCodeEditor) {
      return '/about/succes.png'
    }

    switch (failedAttempts) {
      case 0:
        return '/about/about.png'
      case 1:
        return '/about/fail-1.png'
      case 2:
        return '/about/fail-2.png'
      default:
        return '/about/fail-3.png'
    }
  }

  const codeContent = `// COMISTEICHON HACK 2025 🌭
// ============================

const hackathonInfo = {
  name: "Comisteichon Hack",
  description: \`
    Una iniciativa para fomentar el desarrollo personal 
    y profesional a través de la creación de software e ideas.
    
    🔥 Equipos de 2 personas
    ⏰ 48 horas intensas
    ☕ Cafeína y condiciones cuestionables
    
    El objetivo: crear sin límites y crecer como desarrolladores.
    Los proyectos se inscriben al comienzo, se sortean 
    y se revelan al iniciar el evento.
    
    Al finalizar, cada equipo presentará su creación.
    Habrá premios, aprendizaje, y mucha diversión.
    
    ¿Estás listo para el desafío? 🚀
  \`,
  startHackathon: () => {
    console.log("¡Que comience la magia! ✨");
  }
};`

  return (
    <section className="min-h-screen bg-gray-900 flex">
      <div className="w-full flex">
        <motion.div
          initial={{ opacity: 0, x: -50 }}
          whileInView={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.8 }}
          viewport={{ once: true }}
          className="w-1/2 flex items-center justify-center px-8  md:ml-40 lg:ml-40"
        >
          <div className="w-full max-w-4xl">
            {!showCodeEditor ? (
              <div>
                <h3 className="text-2xl font-mono text-green-400 mb-6">
                  COMISTERMINAL 🌭
                </h3>
                <div className="bg-[#1e1e1e] rounded-lg shadow-2xl overflow-hidden h-[550px] flex flex-col">
                  <div className="bg-[#2d2d30] px-4 py-2 flex items-center space-x-2 flex-shrink-0">
                    <div className="flex space-x-2">
                      <div className="w-3 h-3 bg-red-500 rounded-full"></div>
                      <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
                      <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                    </div>
                    <span className="text-gray-400 text-sm ml-4">terminal</span>
                  </div>

                  <div className="p-4 overflow-y-auto min-w-[600px] flex-1">
                    <InteractiveTerminal
                      onCommandSuccess={() => setShowCodeEditor(true)}
                      failedAttempts={failedAttempts}
                      setFailedAttempts={setFailedAttempts}
                      setIsTyping={setIsTyping}
                    />
                  </div>
                </div>
              </div>
            ) : (
              <div>
                <div className="flex justify-between items-center mb-6">
                  <h3 className="text-2xl font-mono text-green-400">
                    comisteichon/hack.js
                  </h3>
                  <button
                    onClick={() => {
                      setShowCodeEditor(false)
                      setFailedAttempts(0) // Reset failed attempts when going back to terminal
                    }}
                    className="text-gray-400 hover:text-white transition-colors text-2xl font-mono"
                  >
                    ← Back to terminal
                  </button>
                </div>

                <div className="bg-[#1e1e1e] rounded-lg shadow-2xl overflow-hidden h-[550px] flex flex-col">
                  <div className="bg-[#2d2d30] px-4 py-2 flex items-center space-x-2 flex-shrink-0">
                    <div className="flex space-x-2">
                      <div className="w-3 h-3 bg-red-500 rounded-full"></div>
                      <div className="w-3 h-3 bg-yellow-500 rounded-full"></div>
                      <div className="w-3 h-3 bg-green-500 rounded-full"></div>
                    </div>
                    <span className="text-gray-400 text-sm ml-4">hack.js</span>
                  </div>

                  <div className="p-4 overflow-y-auto min-w-[600px] flex-1">
                    <CodeEditor code={codeContent} />
                  </div>
                </div>
              </div>
            )}
          </div>
        </motion.div>

        {/* Right Half - Image */}
        <motion.div
          initial={{ opacity: 0, x: 50 }}
          whileInView={{ opacity: 1, x: 0 }}
          transition={{ duration: 0.8, delay: 0.2 }}
          viewport={{ once: true }}
          className="w-1/2 flex items-center justify-center px-8 md:mr-40 lg:mr-40"
        >
          <div className="w-[550px] h-[550px] flex items-center justify-center">
            <img
              key={getCurrentImage()}
              src={getCurrentImage()}
              alt="COMISTEICHON Event"
              className="w-full h-full object-cover rounded-lg transition-all duration-500 ease-in-out"
              style={{
                maskImage: 'linear-gradient(to bottom, rgba(0,0,0,1) 0%, rgba(0,0,0,1) 40%, rgba(0,0,0,0) 100%)',
                WebkitMaskImage: 'linear-gradient(to bottom, rgba(0,0,0,1) 0%, rgba(0,0,0,1) 70%, rgba(0,0,0,0) 100%)'
              }}
            />
          </div>
        </motion.div>
      </div>
    </section>
  )
}