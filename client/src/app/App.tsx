import { PanoramaViewer } from '../features/panorama/PanoramaViewer';
import { GuessMap } from '../features/map/GuessMap';
import { TimelinePicker } from '../features/timeline/TimelinePicker';
import { useGameStore } from '../stores/gameStore';

export function App() {
  const { panoramaUrl } = useGameStore();
  return <div className='h-screen grid grid-rows-[1fr_auto]'>
    <PanoramaViewer url={panoramaUrl || 'https://cdn.example.com/panos/giza_8k.jpg'} />
    <div className='grid grid-cols-3 gap-3 p-3 bg-zinc-950/95 border-t border-zinc-800'>
      <div className='col-span-2'><GuessMap /></div>
      <TimelinePicker />
    </div>
  </div>;
}
