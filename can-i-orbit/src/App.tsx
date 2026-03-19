import { useState, useEffect, useCallback, useRef } from 'react';
import { SpaceScene } from './components/SpaceScene';
import { Dashboard } from './components/Dashboard';
import { 
  EARTH_RADIUS, 
  integrate, 
  checkCollision, 
  checkEscape 
} from './physics/orbital_mechanics';
import type { State } from './physics/orbital_mechanics';

function App() {
  const [launchParams, setLaunchParams] = useState({ velocity: 7.9, angle: 0 });
  const [isLaunched, setIsLaunched] = useState(false);
  const [status, setStatus] = useState<'idle' | 'orbiting' | 'crashed' | 'escaped'>('idle');
  console.log('App Rendering, status:', status);
  const [state, setState] = useState<State>({
    position: { x: EARTH_RADIUS + 500000, y: 0, z: 0 }, // 高度 500km から開始
    velocity: { x: 0, y: 0, z: 0 },
    time: 0,
  });

  const requestRef = useRef<number>();
  const lastTimeRef = useRef<number>();

  const handleLaunch = () => {
    const v_mag = launchParams.velocity * 1000; // km/s -> m/s
    const angle_rad = (launchParams.angle * Math.PI) / 180;
    
    // 打ち上げ地点 (高度 500km の赤道上と仮定)
    const startPos = { x: EARTH_RADIUS + 500000, y: 0, z: 0 };
    // 速度ベクトル (垂直方向 + 水平方向)
    const startVel = {
      x: v_mag * Math.sin(angle_rad),
      y: v_mag * Math.cos(angle_rad),
      z: 0,
    };

    setState({
      position: startPos,
      velocity: startVel,
      time: 0,
    });
    setIsLaunched(true);
    setStatus('orbiting');
  };

  const handleReset = () => {
    setIsLaunched(false);
    setStatus('idle');
    if (requestRef.current) cancelAnimationFrame(requestRef.current);
  };

  const animate = useCallback((time: number) => {
    if (lastTimeRef.current !== undefined && status === 'orbiting') {
      const dt = 10; // シミュレーションのステップ (秒) - リアルタイムより速く回す
      
      setState((prevState) => {
        const newState = integrate(prevState, dt);
        
        // 判定
        if (checkCollision(newState.position)) {
          setStatus('crashed');
          return prevState;
        }
        if (checkEscape(newState.position, newState.velocity)) {
          setStatus('escaped');
          return newState;
        }
        
        return newState;
      });
    }
    lastTimeRef.current = time;
    requestRef.current = requestAnimationFrame(animate);
  }, [status]);

  useEffect(() => {
    requestRef.current = requestAnimationFrame(animate);
    return () => {
      if (requestRef.current) cancelAnimationFrame(requestRef.current);
    };
  }, [animate]);

  const currentAltitude = Math.sqrt(
    state.position.x ** 2 + state.position.y ** 2 + state.position.z ** 2
  ) - EARTH_RADIUS;

  const currentSpeed = Math.sqrt(
    state.velocity.x ** 2 + state.velocity.y ** 2 + state.velocity.z ** 2
  );

  return (
    <div style={{ width: '100vw', height: '100vh', overflow: 'hidden', position: 'relative' }}>
      <Dashboard
        velocity={launchParams.velocity}
        angle={launchParams.angle}
        onVelocityChange={(v) => setLaunchParams({ ...launchParams, velocity: v })}
        onAngleChange={(a) => setLaunchParams({ ...launchParams, angle: a })}
        onLaunch={handleLaunch}
        onReset={handleReset}
        isLaunched={isLaunched}
        status={status}
        stats={{
          speed: currentSpeed,
          altitude: currentAltitude,
        }}
      />
      <SpaceScene state={state} isLaunched={isLaunched} />
    </div>
  );
}

export default App;
