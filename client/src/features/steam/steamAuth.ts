export type SteamIdentity = { steamId: string; username: string; avatarUrl?: string };

export interface SteamProvider {
  signIn(): Promise<SteamIdentity>;
}

export class DevSteamProvider implements SteamProvider {
  async signIn(): Promise<SteamIdentity> {
    return { steamId: `dev-${crypto.randomUUID()}`, username: 'Dev Player' };
  }
}
