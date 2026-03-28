<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'

const GRID_SIZE = 20
const CELL_SIZE = 20
const GAME_SPEED = 150

type Direction = 'up' | 'down' | 'left' | 'right'
type GameState = 'idle' | 'playing' | 'paused' | 'gameOver'

interface Position {
  x: number
  y: number
}

const canvasRef = ref<HTMLCanvasElement | null>(null)
const gameState = ref<GameState>('idle')
const score = ref(0)
const highScore = ref(0)
const snake = ref<Position[]>([{ x: 10, y: 10 }])
const direction = ref<Direction>('right')
const nextDirection = ref<Direction>('right')
const food = ref<Position>({ x: 15, y: 10 })
let gameLoop: number | null = null

const loadHighScore = () => {
  const saved = localStorage.getItem('snakeHighScore')
  if (saved) {
    highScore.value = parseInt(saved, 10)
  }
}

const saveHighScore = () => {
  if (score.value > highScore.value) {
    highScore.value = score.value
    localStorage.setItem('snakeHighScore', highScore.value.toString())
  }
}

const generateFood = () => {
  let newFood: Position
  do {
    newFood = {
      x: Math.floor(Math.random() * GRID_SIZE),
      y: Math.floor(Math.random() * GRID_SIZE)
    }
  } while (snake.value.some(segment => segment.x === newFood.x && segment.y === newFood.y))
  food.value = newFood
}

const draw = () => {
  const canvas = canvasRef.value
  if (!canvas) return

  const ctx = canvas.getContext('2d')
  if (!ctx) return

  ctx.fillStyle = '#0f172a'
  ctx.fillRect(0, 0, canvas.width, canvas.height)

  ctx.strokeStyle = '#1e293b'
  ctx.lineWidth = 1
  for (let i = 0; i <= GRID_SIZE; i++) {
    ctx.beginPath()
    ctx.moveTo(i * CELL_SIZE, 0)
    ctx.lineTo(i * CELL_SIZE, canvas.height)
    ctx.stroke()
    ctx.beginPath()
    ctx.moveTo(0, i * CELL_SIZE)
    ctx.lineTo(canvas.width, i * CELL_SIZE)
    ctx.stroke()
  }

  const foodX = food.value.x * CELL_SIZE
  const foodY = food.value.y * CELL_SIZE
  const pulseScale = 1 + Math.sin(Date.now() / 200) * 0.1
  const foodRadius = (CELL_SIZE / 2 - 2) * pulseScale
  
  const gradient = ctx.createRadialGradient(
    foodX + CELL_SIZE / 2, foodY + CELL_SIZE / 2, 0,
    foodX + CELL_SIZE / 2, foodY + CELL_SIZE / 2, foodRadius
  )
  gradient.addColorStop(0, '#fb923c')
  gradient.addColorStop(1, '#f97316')
  ctx.fillStyle = gradient
  ctx.beginPath()
  ctx.arc(foodX + CELL_SIZE / 2, foodY + CELL_SIZE / 2, foodRadius, 0, Math.PI * 2)
  ctx.fill()

  snake.value.forEach((segment, index) => {
    const x = segment.x * CELL_SIZE
    const y = segment.y * CELL_SIZE
    const isHead = index === 0
    
    const segmentGradient = ctx.createLinearGradient(x, y, x + CELL_SIZE, y + CELL_SIZE)
    if (isHead) {
      segmentGradient.addColorStop(0, '#2dd4bf')
      segmentGradient.addColorStop(1, '#0d9488')
    } else {
      const alpha = 1 - (index / snake.value.length) * 0.5
      segmentGradient.addColorStop(0, `rgba(45, 212, 191, ${alpha})`)
      segmentGradient.addColorStop(1, `rgba(13, 148, 136, ${alpha})`)
    }
    
    ctx.fillStyle = segmentGradient
    ctx.fillRect(x + 1, y + 1, CELL_SIZE - 2, CELL_SIZE - 2)
    
    if (isHead) {
      ctx.fillStyle = '#ffffff'
      const eyeSize = 3
      const eyeOffset = 5
      
      let eye1X, eye1Y, eye2X, eye2Y
      switch (direction.value) {
        case 'up':
          eye1X = x + eyeOffset
          eye1Y = y + eyeOffset
          eye2X = x + CELL_SIZE - eyeOffset - eyeSize
          eye2Y = y + eyeOffset
          break
        case 'down':
          eye1X = x + eyeOffset
          eye1Y = y + CELL_SIZE - eyeOffset - eyeSize
          eye2X = x + CELL_SIZE - eyeOffset - eyeSize
          eye2Y = y + CELL_SIZE - eyeOffset - eyeSize
          break
        case 'left':
          eye1X = x + eyeOffset
          eye1Y = y + eyeOffset
          eye2X = x + eyeOffset
          eye2Y = y + CELL_SIZE - eyeOffset - eyeSize
          break
        case 'right':
          eye1X = x + CELL_SIZE - eyeOffset - eyeSize
          eye1Y = y + eyeOffset
          eye2X = x + CELL_SIZE - eyeOffset - eyeSize
          eye2Y = y + CELL_SIZE - eyeOffset - eyeSize
          break
      }
      
      ctx.beginPath()
      ctx.arc(eye1X + eyeSize / 2, eye1Y + eyeSize / 2, eyeSize, 0, Math.PI * 2)
      ctx.fill()
      ctx.beginPath()
      ctx.arc(eye2X + eyeSize / 2, eye2Y + eyeSize / 2, eyeSize, 0, Math.PI * 2)
      ctx.fill()
    }
  })
}

const update = () => {
  direction.value = nextDirection.value
  
  const head = { ...snake.value[0] }
  
  switch (direction.value) {
    case 'up':
      head.y -= 1
      break
    case 'down':
      head.y += 1
      break
    case 'left':
      head.x -= 1
      break
    case 'right':
      head.x += 1
      break
  }
  
  if (head.x < 0 || head.x >= GRID_SIZE || head.y < 0 || head.y >= GRID_SIZE) {
    gameOver()
    return
  }
  
  if (snake.value.some(segment => segment.x === head.x && segment.y === head.y)) {
    gameOver()
    return
  }
  
  snake.value.unshift(head)
  
  if (head.x === food.value.x && head.y === food.value.y) {
    score.value += 10
    generateFood()
  } else {
    snake.value.pop()
  }
  
  draw()
}

const gameOver = () => {
  gameState.value = 'gameOver'
  saveHighScore()
  if (gameLoop) {
    cancelAnimationFrame(gameLoop)
    gameLoop = null
  }
}

const startGame = () => {
  snake.value = [{ x: 10, y: 10 }]
  direction.value = 'right'
  nextDirection.value = 'right'
  score.value = 0
  generateFood()
  gameState.value = 'playing'
  
  let lastTime = 0
  const loop = (currentTime: number) => {
    if (gameState.value !== 'playing') return
    
    if (currentTime - lastTime >= GAME_SPEED) {
      update()
      lastTime = currentTime
    }
    
    gameLoop = requestAnimationFrame(loop)
  }
  
  gameLoop = requestAnimationFrame(loop)
}

const togglePause = () => {
  if (gameState.value === 'playing') {
    gameState.value = 'paused'
    if (gameLoop) {
      cancelAnimationFrame(gameLoop)
      gameLoop = null
    }
  } else if (gameState.value === 'paused') {
    gameState.value = 'playing'
    let lastTime = performance.now()
    const loop = (currentTime: number) => {
      if (gameState.value !== 'playing') return
      
      if (currentTime - lastTime >= GAME_SPEED) {
        update()
        lastTime = currentTime
      }
      
      gameLoop = requestAnimationFrame(loop)
    }
    gameLoop = requestAnimationFrame(loop)
  }
}

const changeDirection = (newDir: Direction) => {
  const opposites: Record<Direction, Direction> = {
    up: 'down',
    down: 'up',
    left: 'right',
    right: 'left'
  }
  
  if (opposites[newDir] !== direction.value) {
    nextDirection.value = newDir
  }
}

const handleKeyDown = (e: KeyboardEvent) => {
  if (gameState.value === 'idle' || gameState.value === 'gameOver') {
    if (e.code === 'Space' || e.code === 'Enter') {
      startGame()
    }
    return
  }
  
  if (e.code === 'Space') {
    togglePause()
    return
  }
  
  if (gameState.value !== 'playing') return
  
  switch (e.code) {
    case 'ArrowUp':
    case 'KeyW':
      changeDirection('up')
      break
    case 'ArrowDown':
    case 'KeyS':
      changeDirection('down')
      break
    case 'ArrowLeft':
    case 'KeyA':
      changeDirection('left')
      break
    case 'ArrowRight':
    case 'KeyD':
      changeDirection('right')
      break
  }
}

const animateBackground = () => {
  if (gameState.value === 'playing') {
    draw()
  }
  requestAnimationFrame(animateBackground)
}

onMounted(() => {
  loadHighScore()
  draw()
  animateBackground()
  window.addEventListener('keydown', handleKeyDown)
})

onUnmounted(() => {
  if (gameLoop) {
    cancelAnimationFrame(gameLoop)
  }
  window.removeEventListener('keydown', handleKeyDown)
})
</script>

<template>
  <div class="min-h-screen bg-slate-900 flex flex-col items-center justify-center p-4">
    <div class="text-center mb-6">
      <h1 class="text-5xl font-bold bg-gradient-to-r from-teal-400 to-cyan-500 bg-clip-text text-transparent mb-2">
        Snake Game
      </h1>
      <p class="text-slate-400 text-sm">Use arrow keys or WASD to control</p>
    </div>
    
    <div class="flex gap-8 mb-6">
      <div class="bg-slate-800 rounded-xl px-6 py-3 border border-slate-700">
        <p class="text-slate-400 text-xs uppercase tracking-wider mb-1">Score</p>
        <p class="text-3xl font-bold text-teal-400">{{ score }}</p>
      </div>
      <div class="bg-slate-800 rounded-xl px-6 py-3 border border-slate-700">
        <p class="text-slate-400 text-xs uppercase tracking-wider mb-1">High Score</p>
        <p class="text-3xl font-bold text-orange-400">{{ highScore }}</p>
      </div>
    </div>
    
    <div class="relative">
      <canvas
        ref="canvasRef"
        :width="GRID_SIZE * CELL_SIZE"
        :height="GRID_SIZE * CELL_SIZE"
        class="rounded-xl border-4 border-slate-700 shadow-2xl"
      />
      
      <div v-if="gameState === 'idle'" class="absolute inset-0 bg-slate-900/90 rounded-xl flex flex-col items-center justify-center">
        <div class="text-center">
          <h2 class="text-3xl font-bold text-white mb-4">Ready to Play?</h2>
          <p class="text-slate-400 mb-6">Eat the food, grow the snake, don't hit walls!</p>
          <button
            @click="startGame"
            class="px-8 py-3 bg-gradient-to-r from-teal-500 to-cyan-500 text-white font-bold rounded-lg hover:from-teal-600 hover:to-cyan-600 transition-all transform hover:scale-105 shadow-lg"
          >
            Start Game
          </button>
        </div>
      </div>
      
      <div v-if="gameState === 'paused'" class="absolute inset-0 bg-slate-900/90 rounded-xl flex flex-col items-center justify-center">
        <div class="text-center">
          <h2 class="text-3xl font-bold text-white mb-4">Paused</h2>
          <button
            @click="togglePause"
            class="px-8 py-3 bg-gradient-to-r from-teal-500 to-cyan-500 text-white font-bold rounded-lg hover:from-teal-600 hover:to-cyan-600 transition-all transform hover:scale-105 shadow-lg"
          >
            Resume
          </button>
        </div>
      </div>
      
      <div v-if="gameState === 'gameOver'" class="absolute inset-0 bg-slate-900/90 rounded-xl flex flex-col items-center justify-center">
        <div class="text-center">
          <h2 class="text-4xl font-bold text-red-400 mb-4">Game Over!</h2>
          <p class="text-2xl text-white mb-2">Final Score: <span class="text-teal-400">{{ score }}</span></p>
          <p v-if="score === highScore" class="text-orange-400 mb-6 font-bold">🎉 New High Score! 🎉</p>
          <p v-else class="text-slate-400 mb-6">Try again to beat your high score!</p>
          <button
            @click="startGame"
            class="px-8 py-3 bg-gradient-to-r from-teal-500 to-cyan-500 text-white font-bold rounded-lg hover:from-teal-600 hover:to-cyan-600 transition-all transform hover:scale-105 shadow-lg"
          >
            Play Again
          </button>
        </div>
      </div>
    </div>
    
    <div class="mt-6 grid grid-cols-3 gap-2 md:hidden">
      <div></div>
      <button
        @click="changeDirection('up')"
        class="w-16 h-16 bg-slate-800 rounded-xl border border-slate-700 flex items-center justify-center text-2xl text-teal-400 active:bg-slate-700"
      >
        ↑
      </button>
      <div></div>
      <button
        @click="changeDirection('left')"
        class="w-16 h-16 bg-slate-800 rounded-xl border border-slate-700 flex items-center justify-center text-2xl text-teal-400 active:bg-slate-700"
      >
        ←
      </button>
      <button
        @click="gameState === 'playing' ? togglePause() : startGame()"
        class="w-16 h-16 bg-slate-800 rounded-xl border border-slate-700 flex items-center justify-center text-sm text-teal-400 active:bg-slate-700"
      >
        {{ gameState === 'playing' ? 'Pause' : 'Play' }}
      </button>
      <button
        @click="changeDirection('right')"
        class="w-16 h-16 bg-slate-800 rounded-xl border border-slate-700 flex items-center justify-center text-2xl text-teal-400 active:bg-slate-700"
      >
        →
      </button>
      <div></div>
      <button
        @click="changeDirection('down')"
        class="w-16 h-16 bg-slate-800 rounded-xl border border-slate-700 flex items-center justify-center text-2xl text-teal-400 active:bg-slate-700"
      >
        ↓
      </button>
      <div></div>
    </div>
    
    <div class="mt-8 text-slate-500 text-sm">
      <p>Press <kbd class="px-2 py-1 bg-slate-800 rounded border border-slate-700">Space</kbd> to start/pause</p>
    </div>
  </div>
</template>
