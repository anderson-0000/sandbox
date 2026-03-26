import { useState, useEffect } from 'react';

export const useKeyboard = () => {
  const [keys, setKeys] = useState({
    forward: false,
    backward: false,
    left: false,
    right: false,
    up: false,
    down: false,
  });

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      switch (e.key.toLowerCase()) {
        case 'w':
        case 'arrowup':
          setKeys((k) => ({ ...k, forward: true }));
          break;
        case 's':
        case 'arrowdown':
          setKeys((k) => ({ ...k, backward: true }));
          break;
        case 'a':
        case 'arrowleft':
          setKeys((k) => ({ ...k, left: true }));
          break;
        case 'd':
        case 'arrowright':
          setKeys((k) => ({ ...k, right: true }));
          break;
        case 'r':
          setKeys((k) => ({ ...k, up: true }));
          break;
        case 'f':
          setKeys((k) => ({ ...k, down: true }));
          break;
      }
    };

    const handleKeyUp = (e: KeyboardEvent) => {
      switch (e.key.toLowerCase()) {
        case 'w':
        case 'arrowup':
          setKeys((k) => ({ ...k, forward: false }));
          break;
        case 's':
        case 'arrowdown':
          setKeys((k) => ({ ...k, backward: false }));
          break;
        case 'a':
        case 'arrowleft':
          setKeys((k) => ({ ...k, left: false }));
          break;
        case 'd':
        case 'arrowright':
          setKeys((k) => ({ ...k, right: false }));
          break;
        case 'r':
          setKeys((k) => ({ ...k, up: false }));
          break;
        case 'f':
          setKeys((k) => ({ ...k, down: false }));
          break;
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    window.addEventListener('keyup', handleKeyUp);
    return () => {
      window.removeEventListener('keydown', handleKeyDown);
      window.removeEventListener('keyup', handleKeyUp);
    };
  }, []);

  return keys;
};
