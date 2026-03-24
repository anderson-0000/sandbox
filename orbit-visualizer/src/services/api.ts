export interface SatelliteData {
  name: string;
  line1: string;
  line2: string;
  catId: string;
  type: 'active' | 'debris' | 'starlink' | 'weather' | 'iss' | 'telescope' | 'science_mission' | 'other';
}

const FALLBACK_ISS: SatelliteData = {
  name: "ISS (ZARYA) [OFFLINE]",
  line1: "1 25544U 98067A   24083.82428241  .00016717  00000-0  30143-3 0  9990",
  line2: "2 25544  51.6416  20.4578 0001004  86.5354  33.9114 15.49503487445039",
  catId: "25544",
  type: 'iss'
};

export async function fetchActiveSatellites(): Promise<SatelliteData[]> {
  try {
    console.log("Fetching active satellites via internal server proxy...");
    
    // Viteのプロキシ (/api-celestrak) を使用
    const response = await fetch('/api-celestrak?GROUP=active&FORMAT=tle');
    
    if (!response.ok) throw new Error(`HTTP Error: ${response.status}`);
    
    const text = await response.text();
    console.log(`Received data from proxy (length: ${text.length})`);

    if (!text || text.length < 100) {
      console.warn("Proxy returned empty or too small content.");
      return [FALLBACK_ISS];
    }

    const sats = parseTLE(text, 'active');
    console.log(`Successfully parsed ${sats.length} satellites.`);
    return sats;
  } catch (error) {
    console.error("Failed to fetch satellites:", error);
    return [FALLBACK_ISS];
  }
}

export async function fetchTLEs(group: string, type: SatelliteData['type']): Promise<SatelliteData[]> {
  return fetchActiveSatellites();
}

function parseTLE(tleData: string, type: SatelliteData['type']): SatelliteData[] {
  const lines = tleData.split('\n').map(l => l.trim()).filter(l => l.length > 0);
  const satellites: SatelliteData[] = [];

  for (let i = 0; i < lines.length; i += 3) {
    if (i + 2 < lines.length) {
      const name = lines[i];
      const line1 = lines[i + 1];
      const line2 = lines[i + 2];
      
      if (!line1.startsWith('1 ') || !line2.startsWith('2 ')) {
        continue;
      }

      let finalType = type;
      const upperName = name.toUpperCase();
      
      if (upperName.includes('ISS') || upperName.includes('ZARYA') || upperName.includes('TIANHE')) {
        finalType = 'iss';
      } else if (upperName.includes('HST') || upperName.includes('HUBBLE') || upperName.includes('JWST')) {
        finalType = 'telescope';
      } else if (upperName.includes('STARLINK')) {
        finalType = 'starlink';
      }

      const catId = line2.substring(2, 7).trim();
      satellites.push({ name, line1, line2, catId, type: finalType });
    }
  }
  return satellites;
}
