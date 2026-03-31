export interface TsunamiParams {
  magnitude: number;
  sourceDepth: number;
  sourceDistance: number; // km
}

export const useTsunamiPhysics = () => {
  const G = 9.81; // m/s^2

  // log10(H0) = 0.5 Mw - 3.3
  const calculateInitialHeight = (mw: number) => {
    return Math.pow(10, 0.5 * mw - 3.3);
  };

  // Green's Law: H(D) = H_source * (D_source / D)^(1/4)
  const calculateHeightAtDepth = (hSource: number, dSource: number, dCurrent: number) => {
    const minDepth = 1.0; 
    const depth = Math.max(dCurrent, minDepth);
    return hSource * Math.pow(dSource / depth, 0.25);
  };

  // Calculate arrival time at distance x from shore (km)
  // Assumes a depth profile function D(x)
  // T(x) = integral from sourceDistance to x of (1/v(x')) dx'
  // Numerical integration for arbitrary profiles
  const calculateArrivalTimeNumerical = (
    x: number, 
    sourceDistanceKm: number, 
    getDepthAtX: (dist: number) => number
  ) => {
    const steps = 100;
    const dx = (sourceDistanceKm - x) / steps;
    let time = 0;
    
    for (let i = 0; i < steps; i++) {
      const currentX = sourceDistanceKm - i * dx;
      const depth = Math.max(getDepthAtX(currentX), 0.1);
      const speed = Math.sqrt(G * depth);
      time += (dx * 1000) / speed;
    }
    return time;
  };

  return {
    calculateInitialHeight,
    calculateHeightAtDepth,
    calculateArrivalTimeNumerical,
  };
};
