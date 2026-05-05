import { useGameStore } from '../../stores/gameStore';

export function TimelinePicker() {
  const { minYear, maxYear, guess, setGuess } = useGameStore();
  return <div className='p-2 bg-zinc-900 rounded'>
    <div>When? {guess.year < 0 ? `${Math.abs(guess.year)} BC` : `${guess.year} AD`}</div>
    <input type='range' min={minYear} max={maxYear} value={guess.year} onChange={(e) => setGuess({ year: Number(e.target.value) })} className='w-full' />
  </div>;
}
