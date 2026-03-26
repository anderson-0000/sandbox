import { Canvas } from '@react-three/fiber'
import { OrbitControls, Stars } from '@react-three/drei'
import { useZoomStore } from './store'
import { Tracker } from './Tracker'
import { CameraController } from './CameraController'
import { ScaleContainer } from './ScaleContainer'
import { 
  HumanScale, 
  SkinTextureScale, 
  BloodCellScale, 
  DNAScale, 
  Molecules1nmScale, 
  QuantumElectronScale, 
  Nucleus10fmScale, 
  SubAtomicScale 
} from './components/Scales'
import './App.css'

function App() {
  const zoomLevel = useZoomStore((state) => state.zoomLevel)

  const getScaleData = (z: number) => {
    if (z < 0.5) return { 
      label: '個体 (1.7m)', 
      desc: 'アニメスタイルの人間モデル。皮膚の表面から体内へズームします。' 
    }
    if (z < 1.8) return { 
      label: '皮膚組織 (1mm)', 
      desc: '皮膚を構成する細胞が規則正しく並んでいます。' 
    }
    if (z < 3.8) return { 
      label: '血液のミクロ世界 (10μm)', 
      desc: '無数の赤血球とリンパ球。血管内をイメージした空間です。' 
    }
    if (z < 5.8) return { 
      label: 'DNAの深淵 (100nm)', 
      desc: '二重螺旋構造を持つDNA鎖。生命の設計図が密集しています。' 
    }
    if (z < 7.8) return { 
      label: '分子の密集地帯 (1nm)', 
      desc: '水分子や有機分子がひしめき合う、化学反応の世界です。' 
    }
    if (z < 10.5) return { 
      label: '原子・電子雲 (1Å)', 
      desc: '原子核の周囲に広がる量子的な電子雲。物質の最小構成要素に近づきます。' 
    }
    if (z < 13.5) return { 
      label: '原子核の塊 (10fm)', 
      desc: '高密度の原子核。陽子と中性子が強大な力で結びついています。' 
    }
    return { 
      label: 'クォークの世界 (1fm)', 
      desc: '究極の最小粒子、クォーク。これ以上分割できない物質の深淵です。' 
    }
  }

  const { label, desc } = getScaleData(zoomLevel)

  return (
    <div className="app-container">
      <div className="ui-overlay">
        <h1>ミクロ・ズーム・シミュレーター</h1>
        <div className="scale-info">
          <p>現在のスケール: <strong>{label}</strong></p>
          <p className="description">{desc}</p>
          <p className="zoom-factor">拡大倍率: 10^{zoomLevel.toFixed(1)} 倍</p>
          <p className="hint">スクロールやピンチで拡大・縮小</p>
        </div>
      </div>

      <Canvas
        camera={{ position: [0, 0, 5], fov: 45 }}
        gl={{ antialias: true, alpha: true }}
      >
        <color attach="background" args={['#000000']} />
        <ambientLight intensity={0.6} />
        <pointLight position={[10, 10, 10]} intensity={2.5} />
        <Stars radius={100} depth={50} count={5000} factor={4} saturation={0} fade speed={1} />

        <Tracker />
        <CameraController />
        
        {/* baseScale を指定することで、各階層が適切なズームで表示される */}
        <ScaleContainer range={[-0.5, 1.5]} currentZoom={zoomLevel} baseScale={0}>
          <HumanScale />
        </ScaleContainer>

        <ScaleContainer range={[0.5, 3.0]} currentZoom={zoomLevel} baseScale={1.5}>
          <SkinTextureScale />
        </ScaleContainer>

        <ScaleContainer range={[2.0, 5.0]} currentZoom={zoomLevel} baseScale={3.5}>
          <BloodCellScale />
        </ScaleContainer>

        <ScaleContainer range={[4.0, 7.0]} currentZoom={zoomLevel} baseScale={5.5}>
          <DNAScale />
        </ScaleContainer>

        <ScaleContainer range={[6.0, 9.0]} currentZoom={zoomLevel} baseScale={7.5}>
          <Molecules1nmScale />
        </ScaleContainer>

        <ScaleContainer range={[8.0, 12.0]} currentZoom={zoomLevel} baseScale={10.0}>
          <QuantumElectronScale />
        </ScaleContainer>

        <ScaleContainer range={[11.0, 15.0]} currentZoom={zoomLevel} baseScale={13.0}>
          <Nucleus10fmScale />
        </ScaleContainer>

        <ScaleContainer range={[14.0, 20.0]} currentZoom={zoomLevel} baseScale={16.5}>
          <SubAtomicScale />
        </ScaleContainer>

        <OrbitControls enableZoom={false} enablePan={true} rotateSpeed={0.5} />
      </Canvas>
    </div>
  )
}

export default App
