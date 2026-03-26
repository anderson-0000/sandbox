import { Sphere, Torus, MeshDistortMaterial, MeshWobbleMaterial } from '@react-three/drei'
import { useFrame } from '@react-three/fiber'
import { useRef, useMemo } from 'react'
import * as THREE from 'three'

// --- ユーティリティ: 範囲内にランダム配置するコンポーネント ---
const DistributedObjects = ({ count, area, children }: { count: number, area: number, children: (i: number) => React.ReactNode }) => {
  return (
    <group>
      {Array.from({ length: count }).map((_, i) => (
        <group key={i} position={[(Math.random() - 0.5) * area, (Math.random() - 0.5) * area, (Math.random() - 0.5) * area]}>
          {children(i)}
        </group>
      ))}
    </group>
  )
}

// --- 1. アニメスタイル人間 (1.7m) ---
export const HumanScale = () => {
  const skinMat = <meshStandardMaterial color="#ffe0bd" roughness={0.3} transparent />
  const hairMat = <meshStandardMaterial color="#ff7043" roughness={0.5} transparent />
  const eyeMat = <meshStandardMaterial color="#222222" transparent />
  const clothMat = <meshStandardMaterial color="#ffffff" roughness={0.8} transparent />
  const vestMat = <meshStandardMaterial color="#5c6bc0" roughness={0.8} transparent />

  return (
    <group position={[0, -2.8, 0]} scale={1.3}>
      <mesh position={[0, 2.3, 0]}>
        <capsuleGeometry args={[0.22, 0.7, 4, 16]} />
        {vestMat}
      </mesh>
      <mesh position={[0, 2.3, 0.02]} scale={[1.05, 1, 1.05]}>
        <capsuleGeometry args={[0.2, 0.6, 4, 16]} />
        {clothMat}
      </mesh>
      <group position={[0, 3.5, 0]}>
        <Sphere args={[0.28, 32, 32]}>{skinMat}</Sphere>
        <group position={[0, 0, 0.22]}>
          <group position={[0.1, 0, 0]}>
            <Sphere args={[0.06, 16, 16]} scale={[1.2, 1.5, 0.5]}>{eyeMat}</Sphere>
            <Sphere args={[0.02, 8, 8]} position={[0.02, 0.03, 0.03]}>
              <meshBasicMaterial color="white" transparent />
            </Sphere>
          </group>
          <group position={[-0.1, 0, 0]}>
            <Sphere args={[0.06, 16, 16]} scale={[1.2, 1.5, 0.5]}>{eyeMat}</Sphere>
            <Sphere args={[0.02, 8, 8]} position={[-0.02, 0.03, 0.03]}>
              <meshBasicMaterial color="white" transparent />
            </Sphere>
          </group>
        </group>
        <group position={[0, 0.1, 0]}>
          <Sphere args={[0.3, 16, 16, 0, Math.PI * 2, 0, Math.PI / 1.8]}>{hairMat}</Sphere>
          <group position={[0, -0.05, 0.2]}>
            {[ -0.2, 0, 0.2 ].map((x, i) => (
              <mesh key={i} position={[x, 0, 0]} rotation={[0.4, 0, x * 2]}>
                <capsuleGeometry args={[0.04, 0.2, 4, 8]} />
                {hairMat}
              </mesh>
            ))}
          </group>
        </group>
      </group>
      <mesh position={[0.4, 2.7, 0]} rotation={[0, 0, -0.15]}>
        <capsuleGeometry args={[0.06, 0.9, 4, 8]} />
        {skinMat}
      </mesh>
      <mesh position={[-0.4, 2.7, 0]} rotation={[0, 0, 0.15]}>
        <capsuleGeometry args={[0.06, 0.9, 4, 8]} />
        {skinMat}
      </mesh>
      <mesh position={[0.12, 1.0, 0]}>
        <capsuleGeometry args={[0.09, 1.8, 4, 8]} />
        <meshStandardMaterial color="#37474f" transparent />
      </mesh>
      <mesh position={[-0.12, 1.0, 0]}>
        <capsuleGeometry args={[0.09, 1.8, 4, 8]} />
        <meshStandardMaterial color="#37474f" transparent />
      </mesh>
    </group>
  )
}

// --- 2. 皮膚組織 (1mm) ---
export const SkinTextureScale = () => {
  return (
    <group>
      <DistributedObjects count={80} area={30}>
        {() => (
          <mesh rotation={[Math.random(), Math.random(), 0]}>
            <boxGeometry args={[1.6, 1.1, 0.35]} />
            <MeshDistortMaterial color="#ffe0bd" speed={1} distort={0.2} transparent opacity={0.8} />
          </mesh>
        )}
      </DistributedObjects>
    </group>
  )
}

// --- 3. 赤血球とリンパ球 (10μm) ---
export const BloodCellScale = () => {
  return (
    <group>
      <DistributedObjects count={60} area={40}>
        {(idx) => (
          idx % 15 === 0 ? (
            <Sphere args={[1.3, 16, 16]}>
              <MeshDistortMaterial color="#ffffff" speed={2} distort={0.2} transparent opacity={0.9} />
            </Sphere>
          ) : (
            <group rotation={[Math.random(), Math.random(), 0]}>
              <Torus args={[0.8, 0.4, 12, 24]}>
                <meshStandardMaterial color="#d32f2f" transparent opacity={0.8} />
              </Torus>
              <Sphere args={[0.6, 12, 12]} scale={[1, 1, 0.3]}>
                <meshStandardMaterial color="#d32f2f" transparent opacity={0.8} />
              </Sphere>
            </group>
          )
        )}
      </DistributedObjects>
    </group>
  )
}

// --- 4. 無数のDNA二重螺旋 (100nm) ---
export const DNAScale = () => {
  const DNAChain = () => {
    const steps = 30
    const spiralRadius = 0.5
    const stepHeight = 0.4
    const ref = useRef<THREE.Group>(null!)
    useFrame(() => { if (ref.current) ref.current.rotation.y += 0.012 })
    return (
      <group ref={ref}>
        {Array.from({ length: steps }).map((_, i) => {
          const angle = i * 0.5
          const y = (i - steps / 2) * stepHeight
          return (
            <group key={i} position={[0, y, 0]}>
              <Sphere args={[0.12, 8, 8]} position={[Math.cos(angle) * spiralRadius, 0, Math.sin(angle) * spiralRadius]}>
                <meshStandardMaterial color="#64b5f6" transparent />
              </Sphere>
              <Sphere args={[0.12, 8, 8]} position={[Math.cos(angle + Math.PI) * spiralRadius, 0, Math.sin(angle + Math.PI) * spiralRadius]}>
                <meshStandardMaterial color="#64b5f6" transparent />
              </Sphere>
              <mesh rotation={[0, angle, 0]}>
                <boxGeometry args={[spiralRadius * 2, 0.06, 0.06]} />
                <meshStandardMaterial color={i % 2 === 0 ? "#ef5350" : "#66bb6a"} transparent opacity={0.7} />
              </mesh>
            </group>
          )
        })}
      </group>
    )
  }
  return (
    <group>
      <DistributedObjects count={30} area={50}>
        {() => (
          <group rotation={[Math.random(), Math.random(), 0]} scale={0.5}>
            <DNAChain />
          </group>
        )}
      </DistributedObjects>
    </group>
  )
}

// --- 5. 動的な分子の世界 (1nm) - 結合と解離 ---
export const Molecules1nmScale = () => {
  const ReactiveMolecule = ({ offset }: { offset: number }) => {
    const groupRef = useRef<THREE.Group>(null!)
    const h1Ref = useRef<THREE.Mesh>(null!)
    const h2Ref = useRef<THREE.Mesh>(null!)
    const bond1Ref = useRef<THREE.Mesh>(null!)
    const bond2Ref = useRef<THREE.Mesh>(null!)

    useFrame((state) => {
      const t = state.clock.getElapsedTime() + offset
      
      // 全体の熱運動（ランダムな振動）
      groupRef.current.position.x += Math.sin(t * 10) * 0.01
      groupRef.current.position.y += Math.cos(t * 12) * 0.01
      groupRef.current.rotation.z += Math.sin(t * 2) * 0.01

      // 結合と解離のサイクル (6秒周期)
      // cycle: 0-2s (approach), 2-4s (bonded), 4-6s (leave)
      const cycle = t % 6
      let dist = 3 // 離れている時
      let bondOpacity = 0

      if (cycle < 2) {
        dist = 3 - (cycle / 2) * 2.3 // 3 -> 0.7
        bondOpacity = (cycle / 2)
      } else if (cycle < 4) {
        dist = 0.7 + Math.sin(t * 20) * 0.05 // 結合中（振動）
        bondOpacity = 1
      } else {
        dist = 0.7 + ((cycle - 4) / 2) * 2.3 // 0.7 -> 3
        bondOpacity = 1 - ((cycle - 4) / 2)
      }

      // 水素原子の位置更新 (V字型)
      const angle = 0.9 // 104.5度に近い角度
      h1Ref.current.position.set(Math.cos(angle) * dist, Math.sin(angle) * dist, 0)
      h2Ref.current.position.set(Math.cos(-angle) * dist, Math.sin(-angle) * dist, 0)

      // 結合ラインの更新
      bond1Ref.current.scale.set(1, dist, 1)
      bond1Ref.current.position.set(Math.cos(angle) * dist * 0.5, Math.sin(angle) * dist * 0.5, 0)
      bond1Ref.current.rotation.z = angle - Math.PI / 2
      
      bond2Ref.current.scale.set(1, dist, 1)
      bond2Ref.current.position.set(Math.cos(-angle) * dist * 0.5, Math.sin(-angle) * dist * 0.5, 0)
      bond2Ref.current.rotation.z = -angle - Math.PI / 2

      const mat1 = bond1Ref.current.material as THREE.MeshStandardMaterial
      const mat2 = bond2Ref.current.material as THREE.MeshStandardMaterial
      mat1.opacity = bondOpacity * 0.5
      mat2.opacity = bondOpacity * 0.5
    })

    return (
      <group ref={groupRef}>
        {/* 酸素原子 (中心) */}
        <Sphere args={[0.5, 12, 12]}>
          <meshStandardMaterial color="#ef5350" transparent />
        </Sphere>
        
        {/* 水素原子 1 */}
        <mesh ref={h1Ref}>
          <sphereGeometry args={[0.3, 12, 12]} />
          <meshStandardMaterial color="#ffffff" transparent />
        </mesh>
        
        {/* 水素原子 2 */}
        <mesh ref={h2Ref}>
          <sphereGeometry args={[0.3, 12, 12]} />
          <meshStandardMaterial color="#ffffff" transparent />
        </mesh>

        {/* 結合ライン */}
        <mesh ref={bond1Ref}>
          <cylinderGeometry args={[0.05, 0.05, 1, 8]} />
          <meshStandardMaterial color="#ffffff" transparent opacity={0} emissive="#ffffff" emissiveIntensity={2} />
        </mesh>
        <mesh ref={bond2Ref}>
          <cylinderGeometry args={[0.05, 0.05, 1, 8]} />
          <meshStandardMaterial color="#ffffff" transparent opacity={0} emissive="#ffffff" emissiveIntensity={2} />
        </mesh>
      </group>
    )
  }

  return (
    <group>
      <DistributedObjects count={40} area={45}>
        {(idx) => (
          <group rotation={[Math.random(), Math.random(), 0]}>
            <ReactiveMolecule offset={idx * 1.5} />
          </group>
        )}
      </DistributedObjects>
      
      {/* 背景の小さな分子群（高速に飛び交う） */}
      <DistributedObjects count={60} area={50}>
        {(idx) => {
          const ref = useRef<THREE.Group>(null!)
          useFrame((state) => {
            const t = state.clock.getElapsedTime() + idx
            ref.current.position.x += Math.sin(t * 5) * 0.05
            ref.current.position.z += Math.cos(t * 4) * 0.05
          })
          return (
            <group ref={ref} scale={0.4}>
              <Sphere args={[0.2, 8, 8]}><meshStandardMaterial color="#42a5f5" transparent opacity={0.6} /></Sphere>
            </group>
          )
        }}
      </DistributedObjects>
    </group>
  )
}

// --- 6. 量子的な電子雲パターン (1Å) ---
export const QuantumElectronScale = () => {
  const points = useMemo(() => {
    const p = []
    for (let i = 0; i < 200; i++) {
      const r = Math.pow(Math.random(), 0.5) * 6
      const theta = Math.random() * Math.PI * 2
      const phi = Math.acos(2 * Math.random() - 1)
      p.push(new THREE.Vector3(r * Math.sin(phi) * Math.cos(theta), r * Math.sin(phi) * Math.sin(theta), r * Math.cos(phi)))
    }
    return p
  }, [])
  return (
    <group>
      <DistributedObjects count={15} area={40}>
        {() => (
          <group>
            {points.map((p, idx) => (
              <mesh key={idx} position={p}>
                <sphereGeometry args={[0.09, 4, 4]} />
                <meshBasicMaterial color="#00e5ff" transparent opacity={0.15} />
              </mesh>
            ))}
            <Sphere args={[0.25, 12, 12]}><meshBasicMaterial color="#ffffff" transparent opacity={0.8} /></Sphere>
          </group>
        )}
      </DistributedObjects>
    </group>
  )
}

// --- 7. 原子核の集まり (10fm) ---
export const Nucleus10fmScale = () => {
  const Nucleus = () => {
    const ref = useRef<THREE.Group>(null!)
    useFrame((state) => {
      const t = state.clock.getElapsedTime()
      if (ref.current) {
        ref.current.children.forEach((c, idx) => {
          c.position.x += Math.sin(t * 3 + idx) * 0.06
        })
      }
    })
    return (
      <group ref={ref}>
        {Array.from({ length: 12 }).map((_, idx) => (
          <Sphere key={idx} args={[0.65, 8, 8]} position={[(Math.random() - 0.5) * 1.8, (Math.random() - 0.5) * 1.8, (Math.random() - 0.5) * 1.8]}>
            <meshStandardMaterial color={idx % 2 === 0 ? '#ff1744' : '#29b6f6'} transparent opacity={0.9} />
          </Sphere>
        ))}
      </group>
    )
  }
  return (
    <group>
      <DistributedObjects count={12} area={50}>
        {() => <Nucleus />}
      </DistributedObjects>
    </group>
  )
}

// --- 8. 究極の微細スケール / クォーク (1fm) ---
export const SubAtomicScale = () => {
  const Quarks = () => {
    const ref = useRef<THREE.Group>(null!)
    useFrame((state) => {
      const t = state.clock.getElapsedTime() * 18
      if (ref.current) {
        ref.current.children.forEach((c, idx) => {
          c.position.set(Math.sin(t + idx * 2.2) * 0.7, Math.cos(t * 1.3 + idx) * 0.7, Math.sin(t * 0.9 + idx) * 0.7)
        })
      }
    })
    return (
      <group ref={ref}>
        <Sphere args={[0.35, 8, 8]}><MeshWobbleMaterial color="#ffeb3b" speed={12} factor={1.2} transparent /></Sphere>
        <Sphere args={[0.35, 8, 8]}><MeshWobbleMaterial color="#ffeb3b" speed={12} factor={1.2} transparent /></Sphere>
        <Sphere args={[0.35, 8, 8]}><MeshWobbleMaterial color="#2196f3" speed={12} factor={1.2} transparent /></Sphere>
      </group>
    )
  }
  return (
    <group>
      <DistributedObjects count={8} area={60}>
        {() => <Quarks />}
      </DistributedObjects>
    </group>
  )
}
