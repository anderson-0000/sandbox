import { useState, useEffect } from 'react'
import OrbitScene from './components/OrbitScene'
import InfoOverlay from './components/InfoOverlay'
import AltitudeGauge from './components/AltitudeGauge'
import { type SatelliteData } from './services/api'
import './App.css'

function App() {
  const [selectedSat, setSelectedSat] = useState<SatelliteData | null>(null);
  const [satAltitude, setSatAltitude] = useState(0);
  const [cameraAltitude, setCameraAltitude] = useState(0);
  const [hasData, setHasData] = useState(false);

  const handleSelect = (sat: SatelliteData | null, alt: number) => {
    console.log("APP_SELECT:", sat?.name, "Alt:", alt);
    setSelectedSat(sat);
    setSatAltitude(alt);
  };

  // 衛星データがロードされたときに呼び出される
  const onDataLoaded = (count: number) => {
    console.log(`Data loaded successfully: ${count} satellites`);
    if (count > 0) setHasData(true);
  };

  return (
    <div className="app-container">
      <div className="ui-overlay">
        <h1>Orbit Visualizer</h1>
        <p>Earth to Moon Scale (1:1)</p>
      </div>

      <OrbitScene 
        onSelect={handleSelect} 
        selectedSat={selectedSat} 
        onAltitudeChange={setCameraAltitude} 
        onDataLoaded={onDataLoaded}
      />
      
      <div className="side-panel-left">
        {!hasData && (
          <div className="info-placeholder loading">
            <p>TLEデータを読み込み中...</p>
            <p style={{fontSize: '0.7rem', opacity: 0.5}}>APIへの接続を確認しています</p>
          </div>
        )}

        {hasData && (
          <>
            <InfoOverlay 
              selectedSat={selectedSat} 
              altitude={satAltitude} 
              onClose={() => setSelectedSat(null)} 
            />
            
            {!selectedSat && (
              <div className="info-placeholder">
                <p>衛星または宇宙ゴミを選択してください</p>
              </div>
            )}
          </>
        )}

        <AltitudeGauge altitude={cameraAltitude} />
      </div>

      <div className="legend">
        <div className="legend-item"><span className="dot iss"></span> 宇宙ステーション (ISS/天宮)</div>
        <div className="legend-item"><span className="dot telescope"></span> 宇宙望遠鏡 (ハッブル等)</div>
        <div className="legend-item"><span className="dot science_mission"></span> 科学ミッション (JAXA/NASA)</div>
        <div className="legend-item"><span className="dot active"></span> 運用中の衛星</div>
        <div className="legend-item"><span className="dot starlink"></span> スターリンク</div>
        <div className="legend-item"><span className="dot debris"></span> 宇宙ゴミ (デブリ)</div>
      </div>
    </div>
  )
}

export default App
