<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'

const GRID_SIZE = 20
const TILE_COUNT = 20
const GAME_SPEED = 100

type Position = { x: number; y: number }
type Direction = 'up' | 'down' | 'left' | 'right'
type GameStatus = 'idle' | 'playing' | 'paused' | 'gameOver'

const canvasRef = ref<HTMLCanvasElement | null>(null)
const score = ref(0)
const highScore = ref(0)
const gameStatus = ref<GameStatus>('idle')
const snake = ref<Position[]>([])
const food = ref<Position>({ x: 0, y: 0 })
const direction = ref<Direction>('right')
const nextDirection = ref<Direction>('right')
let gameLoop: number | null = null

const initGame = () => {
  snake.value = [
    { x: 10, y: 10 },
    { x: 9, y: 10 },
    { x: 8, y: 10 }
  ]
  direction.value = 'right'
  nextDirection.value = 'right'
  score.value = 0
  placeFood()
}

const placeFood = () => {
  let newFood: Position
  do {
    newFood = {
      x: Math.floor(Math.random() * TILE_COUNT),
      y: Math.floor(Math.random() * TILE_COUNT)
    }
  } while (snake.value.some(segment => segment.x === newFood.x && segment.y === newFood.y))
  food.value = newFood
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
  
  if (
    head.x < 0 ||
    head.x >= TILE_COUNT ||
    head.y < 0 ||
    head.y >= TILE_COUNT ||
    snake.value.some(segment => segment.x === head.x && segment.y === head.y)
  ) {
    gameOver()
    return
  }
  
  snake.value.unshift(head)
  
  if (head.x === food.value.x && head.y === food.value.y) {
    score.value += 10
    if (score.value > highScore.value) {
      highScore.value = score.value
      localStorage.setItem('snakeHighScore', highScore.value.toString())
    }
    placeFood()
  } else {
    snake.value.pop()
  }
  
  draw()
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
  for (let i = 0; i <= TILE_COUNT; i++) {
    ctx.beginPath()
    ctx.moveTo(i * GRID_SIZE, 0)
    ctx.lineTo(i * GRID_SIZE, canvas.height)
    ctx.stroke()
    ctx.beginPath()
    ctx.moveTo(0, i * GRID_SIZE)
    ctx.lineTo(canvas.width, i * GRID_SIZE)
    ctx.stroke()
  }
  
  snake.value.forEach((segment, index) => {
    const gradient = ctx.createRadialGradient(
      segment.x * GRID_SIZE + GRID_SIZE / 2,
      segment.y * GRID_SIZE + GRID_SIZE / 2,
      0,
      segment.x * GRID_SIZE + GRID_SIZE / 2,
      segment.y * GRID_SIZE + GRID_SIZE / 2,
      GRID_SIZE / 2
    )
    
    if (index === 0) {
      gradient.addColorStop(0, '#34d399')
      gradient.addColorStop(1, '#10b981')
    } else {
      gradient.addColorStop(0, '#6ee7b7')
      gradient.addColorStop(1, '#34d399')
    }
    
    ctx.fillStyle = gradient
    ctx.beginPath()
    ctx.roundRect(
      segment.x * GRID_SIZE + 1,
      segment.y * GRID_SIZE + 1,
      GRID_SIZE - 2,
      GRID_SIZE - 2,
      4
    )
    ctx.fill()
  })
  
  const foodGradient = ctx.createRadialGradient(
    food.value.x * GRID_SIZE + GRID_SIZE / 2,
    food.value.y * GRID_SIZE + GRID_SIZE / 2,
    0,
    food.value.x * GRID_SIZE + GRID_SIZE / 2,
    food.value.y * GRID_SIZE + GRID_SIZE / 2,
    GRID_SIZE / 2
  )
  foodGradient.addColorStop(0, '#fb923c')
  foodGradient.addColorStop(1, '#f97316')
  
  ctx.fillStyle = foodGradient
  ctx.beginPath()
  ctx.arc(
    food.value.x * GRID_SIZE + GRID_SIZE / 2,
    food.value.y * GRID_SIZE + GRID_SIZE / 2,
    GRID_SIZE / 2 - 2,
    0,
    Math.PI * 2
  )
  ctx.fill()
}

const startGame = () => {
  if (gameStatus.value === 'gameOver') {
    initGame()
  }
  gameStatus.value = 'playing'
  if (gameLoop) {
    clearInterval(gameLoop)
  }
  gameLoop = window.setInterval(update, GAME_SPEED)
}

const togglePause = () => {
  if (gameStatus.value === 'playing') {
    gameStatus.value = 'paused'
    if (gameLoop) {
      clearInterval(gameLoop)
      gameLoop = null
    }
  } else if (gameStatus.value === 'paused') {
    gameStatus.value = 'playing'
    gameLoop = window.setInterval(update, GAME_SPEED)
  }
}

const gameOver = () => {
  gameStatus.value = 'gameOver'
  if (gameLoop) {
    clearInterval(gameLoop)
    gameLoop = null
  }
}

const handleKeydown = (e: KeyboardEvent) => {
  if (e.code === 'Space') {
    e.preventDefault()
    if (gameStatus.value === 'idle' || gameStatus.value === 'gameOver') {
      startGame()
    } else {
      togglePause()
    }
    return
  }
  
  if (gameStatus.value !== 'playing') return
  
  switch (e.key) {
    case 'ArrowUp':
    case 'w':
    case 'W':
      if (direction.value !== 'down') nextDirection.value = 'up'
      break
    case 'ArrowDown':
    case 's':
    case 'S':
      if (direction.value !== 'up') nextDirection.value = 'down'
      break
    case 'ArrowLeft':
    case 'a':
    case 'A':
      if (direction.value !== 'right') nextDirection.value = 'left'
      break
    case 'ArrowRight':
    case 'd':
    case 'D':
      if (direction.value !== 'left') nextDirection.value = 'right'
      break
  }
}

onMounted(() => {
  const savedHighScore = localStorage.getItem('snakeHighScore')
  if (savedHighScore) {
    highScore.value = parseInt(savedHighScore, 10)
  }
  
  initGame()
  draw()
  
  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
  if (gameLoop) {
    clearInterval(gameLoop)
  }
})
</script>

<template>
  <div class="min-h-screen bg-slate-950 flex flex-col items-center justify-center p-4 font-mono">
    <div class="text-center mb-6">
      <h1 class="text-5xl font-bold text-teal-400 mb-2 tracking-wider">SNAKE</h1>
      <p class="text-slate-400 text-sm">Use arrow keys or WASD to move</p>
    </div>
    
    <div class="flex gap-8 mb-4">
      <div class="text-center">
        <div class="text-slate-400 text-xs uppercase tracking-widest mb-1">Score</div>
        <div class="text-3xl font-bold text-teal-400">{{ score }}</div>
      </div>
      <div class="text-center">
        <div class="text-slate-400 text-xs uppercase tracking-widest mb-1">High Score</div>
        <div class="text-3xl font-bold text-orange-400">{{ highScore }}</div>
      </div>
    </div>
    
    <div class="relative">
      <canvas
        ref="canvasRef"
        :width="TILE_COUNT * GRID_SIZE"
        :height="TILE_COUNT * GRID_SIZE"
        class="rounded-lg shadow-2xl border-2 border-slate-700"
      />
      
      <div v-if="gameStatus === 'idle'" class="absolute inset-0 bg-slate-900/90 rounded-lg flex flex-col items-center justify-center">
        <h2 class="text-3xl font-bold text-teal-400 mb-4">Ready?</h2>
        <button
          @click="startGame"
          class="px-8 py-3 bg-teal-500 hover:bg-teal-400 text-slate-950 font-bold rounded-full transition-all transform hover:scale-105 animate-pulse"
        >
          Press SPACE to Start
        </button>
      </div>
      
      <div v-if="gameStatus === 'paused'" class="absolute inset-0 bg-slate-900/80 rounded-lg flex flex-col items-center justify-center">
        <h2 class="text-3xl font-bold text-yellow-400 mb-4">Paused</h2>
        <button
          @click="togglePause"
          class="px-8 py-3 bg-yellow-500 hover:bg-yellow-400 text-slate-950 font-bold rounded-full transition-all transform hover:scale-105"
        >
          Press SPACE to Resume
        </button>
      </div>
      
      <div v-if="gameStatus === 'gameOver'" class="absolute inset-0 bg-slate-900/90 rounded-lg flex flex-col items-center justify-center">
        <h2 class="text-4xl font-bold text-red-500 mb-2">Game Over!</h2>
        <p class="text-slate-300 mb-4">Final Score: <span class="text-teal-400 font-bold">{{ score }}</span></p>
        <button
          @click="startGame"
          class="px-8 py-3 bg-teal-500 hover:bg-teal-400 text-slate-950 font-bold rounded-full transition-all transform hover:scale-105"
        >
          Press SPACE to Play Again
        </button>
      </div>
    </div>
    
    <div class="mt-6 flex gap-4">
      <div class="text-center">
        <div class="flex gap-1 mb-1">
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">W</span>
        </div>
        <div class="flex gap-1">
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">A</span>
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">S</span>
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">D</span>
        </div>
      </div>
      <div class="text-slate-500 text-xs flex items-center">or</div>
      <div class="text-center">
        <div class="flex gap-1 mb-1">
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">↑</span>
        </div>
        <div class="flex gap-1">
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">←</span>
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">↓</span>
          <span class="px-3 py-1 bg-slate-800 text-slate-300 rounded text-xs">→</span>
        </div>
      </div>
    </div>
    
    <div class="mt-4 text-slate-600 text-xs">
      Press <kbd class="px-2 py-1 bg-slate-800 rounded text-slate-400">SPACE</kbd> to pause/resume
    </div>
  </div>
</template>
