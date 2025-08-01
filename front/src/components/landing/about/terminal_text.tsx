
import React from 'react';

interface TerminalEditorProps {
  commands: string;
}

function TerminalEditor({ commands }: TerminalEditorProps) {
  return (
    <div className="relative bg-black text-green-400 font-mono">
      <pre className="whitespace-pre-wrap p-0 m-0 text-sm leading-relaxed">
        {commands}
      </pre>
    </div>
  );
}

const terminalContent = `comisteichon@hackathon:~$ cat ./comisteichon-hack-2025/info.md

██████╗ ██████╗ ███╗   ███╗██╗███████╗████████╗███████╗██╗ ██████╗██╗  ██╗ ██████╗ ███╗   ██╗
██╔════╝██╔═══██╗████╗ ████║██║██╔════╝╚══██╔══╝██╔════╝██║██╔════╝██║  ██║██╔═══██╗████╗  ██║
██║     ██║   ██║██╔████╔██║██║███████╗   ██║   █████╗  ██║██║     ███████║██║   ██║██╔██╗ ██║
██║     ██║   ██║██║╚██╔╝██║██║╚════██║   ██║   ██╔══╝  ██║██║     ██╔══██║██║   ██║██║╚██╗██║
╚██████╗╚██████╔╝██║ ╚═╝ ██║██║███████║   ██║   ███████╗██║╚██████╗██║  ██║╚██████╔╝██║ ╚████║
╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚═╝╚══════╝   ╚═╝   ╚══════╝╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝

# COMISTEICHON HACK 2025 🌭
> Una iniciativa para fomentar el desarrollo personal y profesional

## CONFIGURACIÓN DEL EVENTO
───────────────────────────────────────────────────────────────────────────────
🔥 EQUIPOS: 2 personas por equipo
⏰ DURACIÓN: 48 horas intensas  
☕ RECURSOS: Cafeína y condiciones cuestionables
🚀 OBJETIVO: Crear sin límites y crecer como desarrolladores

## PROCESO
───────────────────────────────────────────────────────────────────────────────
1. Los proyectos se inscriben al comienzo
2. Se sortean y revelan al iniciar el evento  
3. 48 horas de desarrollo intenso
4. Cada equipo presenta su creación
5. Premios, aprendizaje y mucha diversión

comisteichon@hackathon:~$ ./start-hackathon.sh
[INFO] Inicializando COMISTEICHON HACK 2025...
[INFO] Cargando equipos participantes...
[SUCCESS] ¡Que comience la magia! ✨
[PROMPT] ¿Estás listo para el desafío? (y/N): _`;

const TerminalText: React.FC = () => {
  return (
    <div className="p-4 h-96 overflow-y-auto min-w-[600px] bg-black">
      <TerminalEditor commands={terminalContent} />
    </div>
  );
};

export default TerminalText;