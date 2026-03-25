import { Sphere, Capsule, Box, Plane, MeshDistortMaterial, MeshWobbleMaterial } from '@react-three/drei'
import { useFrame } from '@react-three/fiber'
import { useRef, useMemo } from 'react'
import * as THREE from 'three'

// --- 1. 人間 (Human with Clothes, Hair, and Eyes) ---
export const HumanScale = () => (
  <group position={[0, -1, 0]}>
    {/* 胴体 (シャツ) */}
    <mesh position={[0, 1.2, 0]}>
      <capsuleGeometry args={[0.3, 0.7, 4, 16]} />
      <meshStandardMaterial color="#2196f3" roughness={0.8} transparent />
    </mesh>
    {/* 腕 (肌) */}
    <mesh position={[0.45, 1.4, 0]} rotation={[0, 0, -0.1]}>
      <capsuleGeometry args={[0.07, 0.6, 4, 8]} />
      <meshStandardMaterial color="#fcd5b4" transparent />
    </mesh>
    <mesh position={[-0.45, 1.4, 0]} rotation={[0, 0, 0.1]}>
      <capsuleGeometry args={[0.07, 0.6, 4, 8]} />
      <meshStandardMaterial color="#fcd5b4" transparent />
    </mesh>
    {/* 頭部 */}
    <group position={[0, 2.1, 0]}>
      <mesh>
        <sphereGeometry args={[0.22, 32, 32]} />
        <meshStandardMaterial color="#fcd5b4" roughness={0.7} transparent />
      </mesh>
      {/* 髪の毛 */}
      <mesh position={[0, 0.1, 0]} rotation={[0.2, 0, 0]}>
        <sphereGeometry args={[0.23, 16, 16, 0, Math.PI * 2, 0, Math.PI / 2]} />
        <meshStandardMaterial color="#4e342e" transparent />
      </mesh>
      {/* 目 */}
      <mesh position={[0.08, 0, 0.18]}>
        <sphereGeometry args={[0.02, 8, 8]} />
        <meshBasicMaterial color="#000000" transparent />
      </mesh>
      <mesh position={[-0.08, 0, 0.18]}>
        <sphereGeometry args={[0.02, 8, 8]} />
        <meshBasicMaterial color="#000000" transparent />
      </mesh>
    </group>
    {/* ズボン */}
    <mesh position={[0.15, 0.45, 0]}>
      <capsuleGeometry args={[0.11, 0.9, 4, 8]} />
      <meshStandardMaterial color="#263238" transparent />
    </mesh>
    <mesh position={[-0.15, 0.45, 0]}>
      <capsuleGeometry args={[0.11, 0.9, 4, 8]} />
      <meshStandardMaterial color="#263238" transparent />
    </mesh>
  </group>
)

// --- 2. 皮膚のきめ (Microscopic Skin Surface) ---
export const SkinTextureScale = () => {
  const cells = useMemo(() => {
    const arr = []
    for (let i = 0; i < 50; i++) {
      arr.push({
        pos: [(Math.random() - 0.5) * 10, (Math.random() - 0.5) * 10, (Math.random() - 0.5) * 0.5],
        scale: 1 + Math.random()
      })
    }
    return arr
  }, [])

  return (
    <group scale={0.1}>
      <Plane args={[15, 15]} position={[0, 0, -0.2]}>
        <meshStandardMaterial color="#f5cbad" transparent />
      </Plane>
      {cells.map((c, i) => (
        <mesh key={i} position={c.pos as any} scale={c.scale} rotation={[0, 0, Math.random()]}>
          <boxGeometry args={[1.5, 1, 0.2]} />
          <MeshDistortMaterial color="#fcd5b4" speed={1} distort={0.2} transparent opacity={0.9} />
        </mesh>
      ))}
    </group>
  )
}

// --- 3. 精密な細胞 (Detailed Biological Cell) ---
export const CellScale = () => {
  return (
    <group scale={0.01}>
      {/* 細胞膜 (流動モザイクモデル風) */}
      <Sphere args={[5, 64, 64]}>
        <meshStandardMaterial color="#ffccbc" transparent opacity={0.3} roughness={0.1} />
      </Sphere>
      
      {/* 細胞核 (内部に染色体が見える演出) */}
      <group>
        <Sphere args={[1.5, 32, 32]}>
          <meshStandardMaterial color="#7e57c2" transparent opacity={0.6} />
        </Sphere>
        <mesh>
          <torusKnotGeometry args={[0.8, 0.1, 64, 8]} />
          <meshStandardMaterial color="#512da8" transparent />
        </mesh>
      </group>

      {/* オルガネラ (ミトコンドリア) */}
      {Array.from({ length: 8 }).map((_, i) => (
        <mesh key={i} position={[Math.sin(i * 1.5) * 3.5, Math.cos(i * 1.5) * 3.5, (Math.random() - 0.5) * 2]}>
          <capsuleGeometry args={[0.2, 0.6, 8, 8]} />
          <meshStandardMaterial color="#d81b60" transparent />
        </mesh>
      ))}
      
      {/* リボソーム (点) */}
      {Array.from({ length: 100 }).map((_, i) => (
        <mesh key={i} position={[(Math.random() - 0.5) * 8, (Math.random() - 0.5) * 8, (Math.random() - 0.5) * 8]}>
          <sphereGeometry args={[0.05, 4, 4]} />
          <meshBasicMaterial color="#ffffff" transparent opacity={0.5} />
        </mesh>
      ))}
    </group>
  )
}

// --- 4. DNA分子 (Nucleotide Chain) ---
export const MoleculeScale = () => {
  const dnaData = useMemo(() => {
    const steps = 120
    const points = []
    for (let i = 0; i < steps; i++) {
      const y = (i - steps / 2) * 0.4
      const angle = i * 0.4
      const r = 1.8
      points.push({ pos: [Math.cos(angle) * r, y, Math.sin(angle) * r], color: '#ffffff', type: 'backbone' })
      points.push({ pos: [Math.cos(angle + Math.PI) * r, y, Math.sin(angle + Math.PI) * r], color: '#ffffff', type: 'backbone' })
      if (i % 2 === 0) {
        const type = i % 4
        const colors = ['#ef5350', '#42a5f5', '#66bb6a', '#ffca28'] // A, T, G, C
        points.push({ pos: [Math.cos(angle) * r * 0.5, y, Math.sin(angle) * r * 0.5], color: colors[type], type: 'base' })
        points.push({ pos: [Math.cos(angle + Math.PI) * r * 0.5, y, Math.sin(angle + Math.PI) * r * 0.5], color: colors[(type + 1) % 4], type: 'base' })
      }
    }
    return points
  }, [])

  const ref = useRef<THREE.Group>(null!)
  useFrame((state) => {
    ref.current.rotation.y = state.clock.getElapsedTime() * 0.2
  })

  return (
    <group ref={ref} scale={0.0001}>
      {dnaData.map((d, i) => (
        <mesh key={i} position={d.pos as any}>
          <sphereGeometry args={[d.type === 'backbone' ? 0.12 : 0.28, 12, 12]} />
          <meshStandardMaterial color={d.color} transparent />
        </mesh>
      ))}
      {/* 水素結合 */}
      {Array.from({ length: 60 }).map((_, i) => (
        <mesh key={i} position={[0, (i * 2 - 60) * 0.4, 0]} rotation={[0, i * 0.8, Math.PI / 2]}>
          <cylinderGeometry args={[0.015, 0.015, 3.6]} />
          <meshStandardMaterial color="#ffffff" transparent opacity={0.2} />
        </mesh>
      ))}
    </group>
  )
}

// --- 5. 原子・分子結合 (C-H Bond / High Contrast) ---
export const AtomicBondScale = () => {
  return (
    <group scale={0.000005}>
      {/* 炭素原子 */}
      <Sphere args={[1, 32, 32]}>
        <meshStandardMaterial color="#222222" roughness={0.1} metalness={0.9} transparent opacity={0.8} />
      </Sphere>
      {/* 水素原子と結合 */}
      {[0, 1, 2, 3].map((i) => {
        const theta = i * Math.PI * 0.5
        const x = Math.cos(theta) * 2.5
        const y = Math.sin(theta) * 2.5
        return (
          <group key={i} rotation={[0, 0, theta]}>
            <mesh position={[1.5, 0, 0]} rotation={[0, 0, Math.PI / 2]}>
              <cylinderGeometry args={[0.1, 0.1, 1.5]} />
              <meshStandardMaterial color="#ffffff" transparent opacity={0.5} />
            </mesh>
            <mesh position={[2.5, 0, 0]}>
              <sphereGeometry args={[0.4, 16, 16]} />
              <meshStandardMaterial color="#eeeeee" transparent />
            </mesh>
          </group>
        )
      })}
    </group>
  )
}

// --- 6. 原子核と電子雲 (Carbon Nucleus & Electron Cloud) ---
export const AtomScale = () => {
  const groupRef = useRef<THREE.Group>(null!)
  const cloudPoints = useMemo(() => {
    const p = []
    for (let i = 0; i < 3000; i++) {
      const r = Math.pow(Math.random(), 0.5) * 6
      const theta = Math.random() * Math.PI * 2
      const phi = Math.acos(2 * Math.random() - 1)
      p.push(new THREE.Vector3(r * Math.sin(phi) * Math.cos(theta), r * Math.sin(phi) * Math.sin(theta), r * Math.cos(phi)))
    }
    return p
  }, [])

  useFrame((state) => {
    groupRef.current.rotation.y += 0.02
    groupRef.current.rotation.z += 0.01
  })

  return (
    <group scale={0.0000001}>
      <group ref={groupRef}>
        {cloudPoints.map((p, i) => (
          <mesh key={i} position={p}>
            <sphereGeometry args={[0.03, 4, 4]} />
            <meshBasicMaterial color="#00ffff" transparent opacity={0.15} />
          </mesh>
        ))}
      </group>
      {/* 原子核 */}
      <Sphere args={[0.5, 32, 32]}>
        <meshStandardMaterial color="#ff1744" emissive="#ff1744" emissiveIntensity={5} />
      </Sphere>
    </group>
  )
}

// --- 7. クォークの世界 (Nucleons / Quarks) ---
export const NucleusScale = () => {
  const ref = useRef<THREE.Group>(null!)
  useFrame((state) => {
    const t = state.clock.getElapsedTime() * 20
    ref.current.children.forEach((c, i) => {
      if (i < 3) {
        c.position.x += Math.sin(t + i) * 0.03
        c.position.y += Math.cos(t + i) * 0.03
      }
    })
  })

  return (
    <group ref={ref} scale={0.00000001}>
      {/* 陽子を構成するクォーク (Up, Up, Down) */}
      <group position={[0.4, 0.4, 0]}>
        <Sphere args={[0.3, 16, 16]}>
          <MeshWobbleMaterial color="#ffeb3b" speed={10} factor={1} transparent />
        </Sphere>
      </group>
      <group position={[-0.4, 0.4, 0]}>
        <Sphere args={[0.3, 16, 16]}>
          <MeshWobbleMaterial color="#ffeb3b" speed={10} factor={1} transparent />
        </Sphere>
      </group>
      <group position={[0, -0.5, 0]}>
        <Sphere args={[0.3, 16, 16]}>
          <MeshWobbleMaterial color="#2196f3" speed={10} factor={1} transparent />
        </Sphere>
      </group>
      
      {/* 強い相互作用のエネルギー体 */}
      <mesh>
        <sphereGeometry args={[1.5, 32, 32]} />
        <meshStandardMaterial color="#ffffff" wireframe transparent opacity={0.1} />
      </mesh>
    </group>
  )
}
