import { Canvas } from '@react-three/fiber'
import { OrbitControls, Stars } from '@react-three/drei'
import { useZoomStore } from './store'
import { Tracker } from './Tracker'
import { CameraController } from './CameraController'
import { ScaleContainer } from './ScaleContainer'
import { HumanScale, SkinTextureScale, CellScale, MoleculeScale, AtomicBondScale, AtomScale, NucleusScale } from './components/Scales'
import './App.css'

function App() {
  const zoomLevel = useZoomStore((state) => state.zoomLevel)

  const getScaleLabel = (z: number) => {
    if (z < 0.8) return 'Individual (1.7m)'
    if (z < 2.5) return 'Dermal Layer (1mm)'
    if (z < 4.5) return 'Somatic Cell (20μm)'
    if (z < 6.5) return 'Genetic Structure / DNA (2nm)'
    if (z < 8.5) return 'Molecular Bonds (Å)'
    if (z < 10.5) return 'Subatomic Particles (pm)'
    return 'Quantum Nucleons (fm)'
  }

  return (
    <div className="app-container">
      <div className="ui-overlay">
        <h1>Life Depth Simulator</h1>
        <div className="scale-info">
          <p>Level: <strong>{getScaleLabel(zoomLevel)}</strong></p>
          <p>Zoom Factor: 10^{zoomLevel.toFixed(1)}</p>
          <p className="hint">Pinch or Scroll to Explore</p>
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
        
        {/* レンジをオーバーラップさせて空間の空白をなくす */}
        <ScaleContainer range={[0, 1.8]} currentZoom={zoomLevel}>
          <HumanScale />
        </ScaleContainer>

        <ScaleContainer range={[0.5, 3.5]} currentZoom={zoomLevel}>
          <SkinTextureScale />
        </ScaleContainer>

        <ScaleContainer range={[2.0, 5.5]} currentZoom={zoomLevel}>
          <CellScale />
        </ScaleContainer>

        <ScaleContainer range={[4.0, 7.5]} currentZoom={zoomLevel}>
          <MoleculeScale />
        </ScaleContainer>

        <ScaleContainer range={[6.0, 9.5]} currentZoom={zoomLevel}>
          <AtomicBondScale />
        </ScaleContainer>

        <ScaleContainer range={[8.0, 11.5]} currentZoom={zoomLevel}>
          <AtomScale />
        </ScaleContainer>

        <ScaleContainer range={[10.0, 14.0]} currentZoom={zoomLevel}>
          <NucleusScale />
        </ScaleContainer>

        <OrbitControls enableZoom={false} enablePan={true} rotateSpeed={0.5} />
      </Canvas>
    </div>
  )
}

export default App
