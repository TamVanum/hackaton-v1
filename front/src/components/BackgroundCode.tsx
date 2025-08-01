import { useEffect, useRef } from 'react'

interface BackgroundCodeProps {
  className?: string
}

export default function BackgroundCode({ className = '' }: BackgroundCodeProps) {
  const containerRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    const container = containerRef.current
    if (!container) return

    // Sample code snippets to use
    const codeSnippets = [
      'const hackathon = () => { console.log("Coding is life"); };',
      'function createMagic() { return Math.random() * 100; }',
      'if (developer.isAwesome()) { return "COMISTEICHON"; }',
      'let dreams = ["code", "coffee", "repeat"]; dreams.forEach(dream => live(dream));',
      'class Hacker { constructor() { this.skills = "legendary"; } }',
      'async function solve(problem) { await think(); return solution; }',
      'const matrix = Array.from({length: 1000}, () => Math.random());',
      'while(true) { code(); debug(); celebrate(); }',
      'export default function Innovation() { return <Future />; }',
      'git commit -m "Another day, another breakthrough";',
      
      // More Go
      'go func() { for range time.Tick(1*time.Second) { fmt.Println("hack") } }()',
      '(╯°□°)╯︵ ┻━┻',

      
      // More Python
      'data = [x**2 for x in range(100) if x % 2 == 0]',
      '@decorator def hack_function(): return "legendary"',
      '(╯°□°)╯︵ ┻━┻',
      
      // C for Arduino
      'void setup() { pinMode(LED_BUILTIN, OUTPUT); Serial.begin(9600); }',
      'digitalWrite(13, HIGH); delay(1000); digitalWrite(13, LOW);',
      'int sensor = analogRead(A0); Serial.println(sensor);',
      
      // Frontend
      'const [state, setState] = useState("hacking"); useEffect(() => { hack(); }, []);',
      'fetch("/api/hack").then(res => res.json()).then(data => console.log(data));',
      'document.querySelector(".hack").addEventListener("click", () => { animate(); });'
    ]

    // Create the repeating code line
    const createCodeLine = () => {
      const line = codeSnippets.join(' ').repeat(10)
      return line
    }

    const codeLine = createCodeLine()
    
    // Create multiple lines to fill 98% of the screen
    const lines: HTMLDivElement[] = []
    const numberOfLines = Math.ceil((window.innerHeight * 0.98) / 25) + 5 // 25px line height with extra lines

    // Opacity levels for variety
    const opacityLevels = ['text-green-900/20', 'text-green-900/30', 'text-green-900/50', 'text-green-900/70', 'text-green-900/80']
    
    for (let i = 0; i < numberOfLines; i++) {
      const lineElement = document.createElement('div')
      
      // Initial random opacity
      const randomOpacity = opacityLevels[Math.floor(Math.random() * opacityLevels.length)]
      
      lineElement.className = `absolute whitespace-nowrap ${randomOpacity} font-mono text-sm select-none pointer-events-none transition-all duration-3000 ease-in-out`
      lineElement.style.top = `${i * 25}px`
      lineElement.style.left = '0'
      lineElement.style.transform = `translateX(${(i % 2 === 0 ? 0 : -200)}px)`
      lineElement.textContent = codeLine
      
      // Add animation
      lineElement.style.animation = `scrollCode ${100 + (i % 5) * 20}s linear infinite`
      
      lines.push(lineElement)
      container.appendChild(lineElement)
    }

    // Add 4 fade-out lines at the bottom for smooth disappearing effect
    const fadeOpacities = ['text-green-900/60', 'text-green-900/40', 'text-green-900/20', 'text-green-900/5']
    
    for (let i = 0; i < 4; i++) {
      const fadeLineElement = document.createElement('div')
      
      fadeLineElement.className = `absolute whitespace-nowrap ${fadeOpacities[i]} font-mono text-sm select-none pointer-events-none transition-all duration-3000 ease-in-out`
      fadeLineElement.style.top = `${(numberOfLines + i) * 25}px`
      fadeLineElement.style.left = '0'
      fadeLineElement.style.transform = `translateX(${((numberOfLines + i) % 2 === 0 ? 0 : -200)}px)`
      fadeLineElement.textContent = codeLine
      
      // Add animation
      fadeLineElement.style.animation = `scrollCode ${100 + ((numberOfLines + i) % 5) * 20}s linear infinite`
      
      lines.push(fadeLineElement)
      container.appendChild(fadeLineElement)
    }

    // Function to randomly change opacity of main lines (excluding fade-out lines)
    const changeOpacity = () => {
      // Only change opacity for the main lines (not the last 4 fade-out lines)
      const mainLines = lines.slice(0, numberOfLines)
      mainLines.forEach(line => {
        // Remove all current opacity classes
        opacityLevels.forEach(opacity => {
          line.classList.remove(opacity)
        })
        
        // Add new random opacity class
        const newOpacity = opacityLevels[Math.floor(Math.random() * opacityLevels.length)]
        line.classList.add(newOpacity)
      })
    }

    // Change opacity every 8 seconds
    const opacityInterval = setInterval(changeOpacity, 8000)

    // Add CSS animation
    const style = document.createElement('style')
    style.textContent = `
      @keyframes scrollCode {
        0% {
          transform: translateX(-50%);
        }
        100% {
          transform: translateX(0%);
        }
      }
    `
    document.head.appendChild(style)

    // Cleanup
    return () => {
      clearInterval(opacityInterval)
      lines.forEach(line => {
        if (container.contains(line)) {
          container.removeChild(line)
        }
      })
      if (document.head.contains(style)) {
        document.head.removeChild(style)
      }
    }
  }, [])

  return (
    <div
      ref={containerRef}
      className={`absolute top-0 left-0 right-0 overflow-hidden pointer-events-none ${className}`}
      style={{ 
        zIndex: -10,
        height: '98%'
      }}
    />
  )
}