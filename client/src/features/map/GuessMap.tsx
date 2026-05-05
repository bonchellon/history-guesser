import { useEffect, useRef } from 'react';
import maplibregl from 'maplibre-gl';
import { useGameStore } from '../../stores/gameStore';

export function GuessMap() {
  const ref = useRef<HTMLDivElement>(null);
  const setGuess = useGameStore((s) => s.setGuess);
  useEffect(() => {
    const map = new maplibregl.Map({ container: ref.current!, style: 'https://demotiles.maplibre.org/style.json', center: [0,0], zoom: 1 });
    const marker = new maplibregl.Marker();
    map.on('click', (e) => { marker.setLngLat(e.lngLat).addTo(map); setGuess({ lat: e.lngLat.lat, lon: e.lngLat.lng }); });
    return () => map.remove();
  }, [setGuess]);
  return <div ref={ref} className='h-56 w-full' />;
}
