package steam

type Identity struct { SteamID, Username, AvatarURL string }

type Provider interface { Verify(ticket string) (Identity, error) }

type DevProvider struct{}

func (DevProvider) Verify(ticket string) (Identity, error) {
	return Identity{SteamID: "dev-user", Username: "Dev Player"}, nil
}
