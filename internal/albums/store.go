package albums

type Store struct {
    albums []Album
}

func NewStore() *Store {
    return &Store{
        albums: []Album{
            {ID: "1", Title: "Iris", Artist: "Goo Goo Dolls", Price: 56.99},
            {ID: "2", Title: "Piano Man", Artist: "Billy Joel", Price: 17.99},
            {ID: "3", Title: "Purnota", Artist: "Warfaze", Price: 39.99},
        },
    }
}

func (s *Store) GetAll() []Album {
    return s.albums
}

func (s *Store) GetByID(id string) (Album, bool) {
    for _, a := range s.albums {
        if a.ID == id {
            return a, true
        }
    }
    return Album{}, false
}

func (s *Store) Add(a Album) {
    s.albums = append(s.albums, a)
}
