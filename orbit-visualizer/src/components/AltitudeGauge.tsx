interface AltitudeGaugeProps {
  altitude: number;
}

export default function AltitudeGauge({ altitude }: AltitudeGaugeProps) {
  const getOrbitRegion = (alt: number) => {
    if (alt < 2000) return '低軌道 (LEO)';
    if (alt < 35786) return '中軌道 (MEO)';
    if (alt < 40000) return '静止軌道 (GEO)';
    if (alt < 360000) return '深宇宙';
    if (alt < 400000) return '月周辺';
    return '月より遠方';
  };

  return (
    <div className="altitude-gauge">
      <div className="gauge-label">カメラ高度</div>
      <div className="gauge-value">
        {Math.round(altitude).toLocaleString()} <span className="unit">km</span>
      </div>
      <div className="gauge-region">{getOrbitRegion(altitude)}</div>
    </div>
  );
}
