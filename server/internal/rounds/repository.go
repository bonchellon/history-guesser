package rounds

import (
	"context"
	"math/rand"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Round struct {
	ID int64 `json:"id"`
	Title, PanoramaURL string `json:"title"`
	CorrectLatitude, CorrectLongitude float64 `json:"correctLatitude"`
	CorrectYear, MinYear, MaxYear int `json:"correctYear"`
	LocationName string `json:"locationName"`
}

type Repository interface { Random(context.Context) (Round, error) }

type PGRepository struct { db *pgxpool.Pool }
func NewPGRepository(db *pgxpool.Pool) *PGRepository { return &PGRepository{db:db} }

func (r *PGRepository) Random(ctx context.Context) (Round, error) {
	rows, err := r.db.Query(ctx, `SELECT id,title,panorama_url,correct_latitude,correct_longitude,correct_year,min_year,max_year,location_name FROM rounds`)
	if err != nil { return Round{}, err }
	defer rows.Close()
	all := []Round{}
	for rows.Next() { var x Round; _ = rows.Scan(&x.ID,&x.Title,&x.PanoramaURL,&x.CorrectLatitude,&x.CorrectLongitude,&x.CorrectYear,&x.MinYear,&x.MaxYear,&x.LocationName); all=append(all,x)}
	if len(all)==0 { return Round{ID:1,Title:"Fallback",PanoramaURL:"https://example.com/pano.jpg",MinYear:-3000,MaxYear:2026,CorrectYear:1900}, nil }
	return all[rand.Intn(len(all))], nil
}
