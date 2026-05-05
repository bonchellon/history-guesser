export type ClientMessage =
  | { type: 'room.create' }
  | { type: 'room.join'; code: string }
  | { type: 'room.leave'; code: string }
  | { type: 'player.ready'; code: string; ready: boolean }
  | { type: 'match.start'; code: string }
  | { type: 'guess.update'; code: string; lat: number; lon: number; year: number }
  | { type: 'guess.submit'; code: string; lat: number; lon: number; year: number }
  | { type: 'ping' };

export type ServerMessage =
  | { type: 'room.state'; payload: unknown }
  | { type: 'match.started'; payload: unknown }
  | { type: 'round.started'; payload: unknown }
  | { type: 'round.timer'; payload: { remainingMs: number } }
  | { type: 'guess.accepted'; payload: { playerId: string } }
  | { type: 'round.ended'; payload: unknown }
  | { type: 'match.ended'; payload: unknown }
  | { type: 'error'; payload: { message: string } };
