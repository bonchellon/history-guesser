import type { ClientMessage, ServerMessage } from '../../../shared/protocol/messages';

export class WsClient {
  private ws?: WebSocket;
  private retries = 0;
  constructor(private url: string, private onMessage: (msg: ServerMessage) => void) {}
  connect() {
    this.ws = new WebSocket(this.url);
    this.ws.onmessage = (e) => this.onMessage(JSON.parse(e.data));
    this.ws.onclose = () => setTimeout(() => { this.retries++; this.connect(); }, Math.min(1000 * 2 ** this.retries, 10_000));
  }
  send(message: ClientMessage) { this.ws?.send(JSON.stringify(message)); }
}
