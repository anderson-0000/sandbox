import { type SatelliteData } from '../services/api';
import { X } from 'lucide-react';

interface InfoOverlayProps {
  selectedSat: SatelliteData | null;
  onClose: () => void;
  altitude: number;
}

export default function InfoOverlay({ selectedSat, onClose, altitude }: InfoOverlayProps) {
  if (!selectedSat) return null;

  return (
    <div className="info-overlay">
      <button className="close-button" onClick={onClose}>
        <X size={18} />
      </button>
      <div className="info-header">
        <h2>{selectedSat.name}</h2>
        <span className={`type-badge ${selectedSat.type}`}>{selectedSat.type.toUpperCase()}</span>
      </div>
      <div className="info-content">
        <div className="info-row">
          <span className="label">NORAD ID</span>
          <span className="value">{selectedSat.catId}</span>
        </div>
        <div className="info-row">
          <span className="label">Altitude</span>
          <span className="value">{Math.round(altitude).toLocaleString()} km</span>
        </div>
        <div className="info-row">
          <span className="label">TLE L1</span>
          <span className="value code">{selectedSat.line1}</span>
        </div>
        <div className="info-row">
          <span className="label">TLE L2</span>
          <span className="value code">{selectedSat.line2}</span>
        </div>
      </div>
    </div>
  );
}
