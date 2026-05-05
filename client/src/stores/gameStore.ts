import { create } from 'zustand';

type Guess = { lat: number; lon: number; year: number; submitted: boolean };
type State = {
  roomCode: string;
  panoramaUrl: string;
  minYear: number;
  maxYear: number;
  guess: Guess;
  setGuess: (g: Partial<Guess>) => void;
  setRound: (pano: string, minYear: number, maxYear: number) => void;
};

export const useGameStore = create<State>((set) => ({
  roomCode: '', panoramaUrl: '', minYear: -3000, maxYear: 2026,
  guess: { lat: 0, lon: 0, year: 1900, submitted: false },
  setGuess: (g) => set((s) => ({ guess: { ...s.guess, ...g } })),
  setRound: (panoramaUrl, minYear, maxYear) => set({ panoramaUrl, minYear, maxYear }),
}));
